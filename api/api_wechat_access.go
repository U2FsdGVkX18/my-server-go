package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"my-server-go/config/mysql"
	logger "my-server-go/tool/log"
	"my-server-go/tool/wechataes"
	"net/http"
	"time"
)

const sCorpID = "ww8d5186f5aa839ee7"

const sToken = "RlWIx3oywb6O"

const sEncodingAESKey = "YE7McxpnpLYFAO8qIxrkjbxqjTjt6rfSjBeEiV3YDNv"

// WeChatReqMsgInfo 接收到的消息加密格式
type WeChatReqMsgInfo struct {
	ToUserName string
	AgentID    string
	Encrypt    string
}

// MsgContent 接收到的消息经过解密后的格式
type MsgContent struct {
	ToUsername   string  `xml:"ToUserName"`
	FromUsername string  `xml:"FromUserName"`
	CreateTime   uint32  `xml:"CreateTime"`
	MsgType      string  `xml:"MsgType"`
	Content      string  `xml:"Content"`
	Msgid        uint64  `xml:"MsgId"`
	Agentid      uint32  `xml:"AgentID"`
	Event        string  `xml:"Event"`
	Latitude     float32 `xml:"Latitude"`
	Longitude    float32 `xml:"Longitude"`
	Precision    float32 `xml:"Precision"`
	AppType      string  `xml:"AppType"`
}

func WeChatAccess(ginServer *gin.Engine) {
	//接入Prometheus
	reqCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of HTTP requests",
	}, []string{"method", "endpoint", "status"})

	//请求持续时间
	reqDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests",
		Buckets: prometheus.DefBuckets,
	}, []string{"method", "endpoint", "status"})

	reqConcurrent := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "http_concurrent_requests",
		Help: "Number of concurrent HTTP requests",
	})

	prometheus.MustRegister(reqCounter, reqDuration, reqConcurrent)

	//使用中间件
	ginServer.Use(func(context *gin.Context) {
		method := context.Request.Method
		endpoint := context.FullPath()

		start := time.Now()
		reqConcurrent.Inc()

		context.Next()

		duration := time.Since(start)
		status := context.Writer.Status()

		reqConcurrent.Dec()

		reqCounter.WithLabelValues(method, endpoint, fmt.Sprintf("%d", status)).Inc()
		reqDuration.WithLabelValues(method, endpoint, fmt.Sprintf("%d", status)).Observe(duration.Seconds())
	})

	//Add Prometheus metrics endpoint
	ginServer.GET("/metrics", gin.WrapH(promhttp.Handler()))

	var wechatGroup = ginServer.Group("/wechat")
	{
		wechatGroup.GET("/recall", func(context *gin.Context) {
			//获取参数
			var msg_signature = context.Query("msg_signature")
			var timestamp = context.Query("timestamp")
			var nonce = context.Query("nonce")
			var echostr = context.Query("echostr")
			logger.Write("接口参数为:", msg_signature, timestamp, nonce, echostr)
			//验证URL
			echoStr := VerifyUrl(msg_signature, timestamp, nonce, echostr)
			//响应
			context.String(http.StatusOK, echoStr)
			return
		})
		wechatGroup.POST("/recall", func(context *gin.Context) {
			//获取参数
			var msg_signature = context.Query("msg_signature")
			var timestamp = context.Query("timestamp")
			var nonce = context.Query("nonce")
			post_data, _ := context.GetRawData()
			logger.Write("接口参数为:", msg_signature, timestamp, nonce, string(post_data))
			//处理消息
			message := ProcessMessage(msg_signature, timestamp, nonce, post_data)
			context.String(http.StatusOK, message)
			return
		})
	}
}

// VerifyUrl 验证url
func VerifyUrl(msg_signature string, timestamp string, nonce string, echostr string) string {
	wxcpt := wxmsgcrypt.NewWXBizMsgCrypt(sToken, sEncodingAESKey, sCorpID, wxmsgcrypt.JsonType)
	echoStr, cryptError := wxcpt.VerifyURL(msg_signature, timestamp, nonce, echostr)
	if cryptError != nil {
		logger.Write("VerifyUrl 验证失败!", cryptError)
	}
	logger.Write("VerifyUrl 验证成功!", "resp:echoStr:"+string(echoStr))
	return string(echoStr)
}

// ProcessMessage 处理接收到的消息
func ProcessMessage(msg_signature string, timestamp string, nonce string, post_data []byte) string {
	wxcpt := wxmsgcrypt.NewWXBizMsgCrypt(sToken, sEncodingAESKey, sCorpID, wxmsgcrypt.JsonType)
	//定义结构体
	var weChatReqMsgInfo WeChatReqMsgInfo
	//将返回的xml转为结构体
	err := xml.Unmarshal(post_data, &weChatReqMsgInfo)
	if err != nil {
		logger.Write("ProcessMessage xml解析失败!", err)
	}
	//将结构体转为json
	data, err := json.Marshal(weChatReqMsgInfo)
	//把json给解密函数,返回一个xml
	msg, cryptError := wxcpt.DecryptMsg(msg_signature, timestamp, nonce, data)
	if cryptError != nil {
		logger.Write("ProcessMessage 消息处理失败!", *cryptError)
	} else {
		logger.Write("ProcessMessage 消息处理成功!", string(msg))
	}

	//定义结构体
	var msgContent MsgContent
	//将xml转为结构体
	err = xml.Unmarshal(msg, &msgContent)
	if err != nil {
		logger.Write("ProcessMessage 消息解析失败!", err)
	} else {
		logger.Write("ProcessMessage 消息解析成功!", msgContent)
	}

	switch msgContent.MsgType {
	case "text":
		logger.Write("ProcessMessage 检测到机器人接收到消息的内容:", msgContent.Content)
		username := weChatReqMsgInfo.ToUserName
		switch msgContent.Content {
		case "能力":
			str := "【赋予能力✅】" + "\n" +
				"\n" +
				"暂无" +
				"\n" +
				""
			text := CreatePassiveRespText(username, timestamp, nonce, str)
			return text
		default:
			text := CreatePassiveRespText(username, timestamp, nonce, "未知")
			return text
		}
	case "event":
		if msgContent.Event == "LOCATION" {
			username := msgContent.FromUsername
			var Latitude = fmt.Sprintf("%f", msgContent.Latitude)
			var Longitude = fmt.Sprintf("%f", msgContent.Longitude)
			location := Latitude + ":" + Longitude
			db := mysql.Connect()
			result := db.Where("user_name", username).First(&mysql.QywxUserLocation{})
			if result.RowsAffected == 1 {
				//更新
				db.Model(&mysql.QywxUserLocation{}).Where("user_name", username).Update("user_location", location)
				logger.Write("ProcessMessage 查询到用户和位置数据已存在,更新位置数据:", username)
			} else {
				//插入
				db.Create(&mysql.QywxUserLocation{
					UserName:     username,
					UserLocation: location,
				})
				logger.Write("ProcessMessage 查询到用户和位置数据不存在,插入位置数据:", username)
			}
			logger.Write("ProcessMessage 检测到机器人接收到的位置信息:", Latitude, Longitude)
		}
	}
	return ""
}

type WeChatPassiveRespMsgInfo struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      string   `xml:"Encrypt" json:"Encrypt"`
	MsgSignature string   `xml:"MsgSignature" json:"MsgSignature"`
	TimeStamp    string   `xml:"TimeStamp" json:"TimeStamp"`
	Nonce        string   `xml:"Nonce" json:"Nonce"`
}

// CreatePassiveRespText 构造被动消息
func CreatePassiveRespText(username string, timestamp string, nonce string, content string) string {
	//拿到封装好的被动消息结构体
	weChatPassiveRespMsg := GetPassiveRespMessageBodyText(username, timestamp, "text", content)
	//将被动消息结构体转为xml
	weChatPassiveRespMsgXml, _ := xml.Marshal(weChatPassiveRespMsg)
	logger.Write("weChatPassiveRespMsg :", string(weChatPassiveRespMsgXml))
	//将该xml进行加密,得到一个json
	wxcpt := wxmsgcrypt.NewWXBizMsgCrypt(sToken, sEncodingAESKey, sCorpID, wxmsgcrypt.JsonType)
	msg, _ := wxcpt.EncryptMsg(string(weChatPassiveRespMsgXml), timestamp, nonce)
	logger.Write("weChatPassiveRespMsg EncryptMsg:", string(msg))
	var weChatPassiveRespMsgInfo WeChatPassiveRespMsgInfo
	//将json转为结构体
	_ = json.Unmarshal(msg, &weChatPassiveRespMsgInfo)
	//将结构体转为xml返回
	marshal, _ := xml.Marshal(weChatPassiveRespMsgInfo)
	logger.Write("weChatPassiveRespMsg weChatPassiveRespMsgInfo:", string(marshal))
	return string(marshal)
}

// WeChatPassiveRespMsg 被动响应消息结构体
type WeChatPassiveRespMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUsername   string   `xml:"ToUserName"`
	FromUsername string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
}

// GetPassiveRespMessageBodyText 封装被动响应消息格式
func GetPassiveRespMessageBodyText(userName string, createTime string, messageType string, content string) WeChatPassiveRespMsg {
	weChatPassiveRespMsg := WeChatPassiveRespMsg{
		ToUsername:   userName,
		FromUsername: userName,
		CreateTime:   createTime,
		MsgType:      messageType,
		Content:      content,
	}
	return weChatPassiveRespMsg
}
