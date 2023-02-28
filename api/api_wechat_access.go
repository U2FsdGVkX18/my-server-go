package api

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"log"
	logger "my-server-go/commons/tools/log"
	"my-server-go/commons/tools/wechataes/wxmsgcrypt"
	"net/http"
)

const sCorpID = "ww8d5186f5aa839ee7"

const sToken = "RlWIx3oywb6O"

const sEncodingAESKey = "YE7McxpnpLYFAO8qIxrkjbxqjTjt6rfSjBeEiV3YDNv"

type WeChatReqMsgInfo struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string   `xml:"ToUserName"`
	AgentID    string   `xml:"AgentID"`
	Encrypt    string   `xml:"Encrypt"`
}

type MsgContent struct {
	ToUsername   string `json:"ToUserName"`
	FromUsername string `json:"FromUserName"`
	CreateTime   uint32 `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Content      string `json:"Content"`
	Msgid        uint64 `json:"MsgId"`
	Agentid      uint32 `json:"AgentId"`
}

func WeChatAccess(ginServer *gin.Engine) {
	var wechatGroup = ginServer.Group("/wechat")
	{
		wechatGroup.GET("/recall", func(context *gin.Context) {
			var msg_signature = context.Query("msg_signature")
			var timestamp = context.Query("timestamp")
			var nonce = context.Query("nonce")
			var echostr = context.Query("echostr")
			logger.Write("接口参数为:", msg_signature, timestamp, nonce, echostr)
			//验证URL
			echoStr := VerifyUrl(msg_signature, timestamp, nonce, echostr)
			context.String(http.StatusOK, echoStr)
		})
		wechatGroup.POST("/recall", func(context *gin.Context) {
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
	msg, cryptError := wxcpt.DecryptMsg(msg_signature, timestamp, nonce, post_data)
	if cryptError != nil {
		log.Println("ProcessMessage 消息处理失败!", cryptError)
		logger.Write("ProcessMessage 消息处理失败!", cryptError)
	}

	//创建对象
	var msgContent MsgContent
	err := json.Unmarshal(msg, &msgContent)
	if err != nil {
		log.Println("ProcessMessage 消息解析失败!", err)
		logger.Write("ProcessMessage 消息解析失败!", err)
	} else {
		log.Println("ProcessMessage 消息解析成功!", msgContent)
		logger.Write("ProcessMessage 消息解析成功!", msgContent)
	}

}
