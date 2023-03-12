package wx

import (
	"encoding/json"
	"fmt"
	"io"
	"my-server-go/config/redis"
	"my-server-go/invoke"
	logger "my-server-go/tool/log"
	"time"
)

// SendWxMessage 主动发送微信消息
func SendWxMessage(content string) {
	//构造主动发送消息体
	map1 := make(map[string]any)
	map1["touser"] = "@all"
	map1["agentid"] = "1000002"
	map1["msgtype"] = "text"
	map2 := make(map[string]string)
	map2["content"] = content
	map1["text"] = map2
	marshal, _ := json.Marshal(map1)

	//获取token
	token := GetAccessToken()
	sendMessageUrl := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + token
	//发送请求
	resp := invoke.SendPost(sendMessageUrl, marshal, nil)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)
	m := make(map[string]any)
	_ = json.Unmarshal(body, &m)
	fmt.Println(m["errcode"])
}

type TokenBody struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetAccessToken 获取accessToken
func GetAccessToken() string {
	//首先从redis中获取token
	token := redis.GetValue("wxAccessToken")
	if token == "" {
		//如果token值为空,发送请求请求token
		const corpId = "ww8d5186f5aa839ee7"
		const corpSecret = "FQcVXdjCMF2N6GVFGQcbe6izTKtS8tlHi8fPpCgl2PU"
		const getAccessTokenUrl = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
		//封装url参数
		params := make(map[string]string)
		params["corpId"] = corpId
		params["corpSecret"] = corpSecret
		//发送请求
		resp := invoke.SendGet(getAccessTokenUrl, params, nil)
		//defer关闭io流
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
		//读取body
		body, _ := io.ReadAll(resp.Body)
		//定义结构体
		var tokenBody TokenBody
		//将json格式转为对应结构体
		_ = json.Unmarshal(body, &tokenBody)
		//先把请求到的token放入redis
		redis.SetValue("redisWxAccessToken", tokenBody.AccessToken, 7200*1000*time.Millisecond)
		//再将数据进行返回
		logger.Write("newWxAccessToken :", tokenBody.AccessToken)
		return tokenBody.AccessToken
	} else {
		//否则直接返回redis中的token,减少请求提高效率
		logger.Write("redisWxAccessToken :", token)
		return token
	}
}
