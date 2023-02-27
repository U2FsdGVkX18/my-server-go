package api

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"my-server-go/commons/tools/log"
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

func WeChatAccess(ginServer *gin.Engine) {
	var wechatGroup = ginServer.Group("/wechat")
	{
		wechatGroup.GET("/recall", func(context *gin.Context) {
			var msg_signature = context.Query("msg_signature")
			var timestamp = context.Query("timestamp")
			var nonce = context.Query("nonce")
			var echostr = context.Query("echostr")
			//log
			log.Write("接口参数为:", msg_signature, timestamp, nonce, echostr)
			//验证URL
			echoStr := VerifyUrl(msg_signature, timestamp, nonce, echostr)
			context.String(http.StatusOK, echoStr)
		})
		wechatGroup.POST("/recall", func(context *gin.Context) {
			var msg_signature = context.Query("msg_signature")
			var timestamp_string = context.Query("timestamp")
			var nonce = context.Query("nonce")
			data, _ := context.GetRawData()
			//log
			log.Write("接口参数为:", msg_signature, timestamp_string, nonce, data)
			//将string -> int
			//timestamp, _ := strconv.Atoi(timestamp_string)
			//处理消息
			//GetMessage(msg_signature, timestamp, nonce, data)
		})
	}
}

func VerifyUrl(msg_signature string, timestamp string, nonce string, echostr string) string {
	wxcpt := wxmsgcrypt.NewWXBizMsgCrypt(sToken, sEncodingAESKey, sCorpID, wxmsgcrypt.JsonType)
	echoStr, cryptError := wxcpt.VerifyURL(msg_signature, timestamp, nonce, echostr)
	if cryptError != nil {
		log.Write("验证失败!")
	}
	log.Write("验证成功!", echoStr)
	return string(echoStr)
}

func GetMessage(msg_signature string, timestamp int, nonce string, info WeChatReqMsgInfo) {

}
