package api

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"log"
	logger "my-server-go/tool/log"
	"my-server-go/tool/wechataes/wxmsgcrypt"
	"net/http"
)

const sCorpID = "ww8d5186f5aa839ee7"

const sToken = "RlWIx3oywb6O"

const sEncodingAESKey = "YE7McxpnpLYFAO8qIxrkjbxqjTjt6rfSjBeEiV3YDNv"

type WeChatReqMsgInfo struct {
	ToUserName string
	AgentID    string
	Encrypt    string
}

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
		})
		wechatGroup.POST("/recall", func(context *gin.Context) {
			//获取参数
			var msg_signature = context.Query("msg_signature")
			var timestamp = context.Query("timestamp")
			var nonce = context.Query("nonce")
			post_data, _ := context.GetRawData()
			logger.Write("接口参数为:", msg_signature, timestamp, nonce, string(post_data))
			//处理消息
			ProcessMessage(msg_signature, timestamp, nonce, post_data)

		})
	}
}

func VerifyUrl(msg_signature string, timestamp string, nonce string, echostr string) string {
	wxcpt := wxmsgcrypt.NewWXBizMsgCrypt(sToken, sEncodingAESKey, sCorpID, wxmsgcrypt.JsonType)
	echoStr, cryptError := wxcpt.VerifyURL(msg_signature, timestamp, nonce, echostr)
	if cryptError != nil {
		logger.Write("VerifyUrl 验证失败!", cryptError)
	}
	logger.Write("VerifyUrl 验证成功!", "resp:echoStr:"+string(echoStr))
	return string(echoStr)
}

func ProcessMessage(msg_signature string, timestamp string, nonce string, post_data []byte) {
	wxcpt := wxmsgcrypt.NewWXBizMsgCrypt(sToken, sEncodingAESKey, sCorpID, wxmsgcrypt.JsonType)
	//定义结构体
	var weChatReqMsgInfo WeChatReqMsgInfo
	//将返回的xml转为结构体
	err := xml.Unmarshal(post_data, &weChatReqMsgInfo)
	if err != nil {
		return
	}
	//将结构体转为json
	data, err := json.Marshal(weChatReqMsgInfo)
	//把json给解密函数,返回一个xml
	msg, cryptError := wxcpt.DecryptMsg(msg_signature, timestamp, nonce, data)
	if cryptError != nil {
		logger.Write("ProcessMessage 消息处理失败!", *cryptError)
	} else {
		log.Println("ProcessMessage 消息处理成功!", string(msg))
		logger.Write("ProcessMessage 消息处理成功!", string(msg))
	}

	//定义结构体
	var msgContent MsgContent
	//将xml转为结构体
	err = xml.Unmarshal(msg, &msgContent)
	if err != nil {
		logger.Write("ProcessMessage 消息解析失败!", err)
	} else {
		log.Println("ProcessMessage 消息解析成功!", msgContent)
		logger.Write("ProcessMessage 消息解析成功!", msgContent)
	}

	switch msgContent.MsgType {
	case "text":
		log.Println("ProcessMessage 检测到机器人接收到消息的内容:", msgContent.Content)
		logger.Write("ProcessMessage 检测到机器人接收到消息的内容:", msgContent.Content)
	case "event":
		if msgContent.Event == "LOCATION" {
			var Latitude = msgContent.Latitude
			var Longitude = msgContent.Longitude
			logger.Write("ProcessMessage 检测到机器人接收到的位置信息:", Latitude, Longitude)
		}
	}
}
