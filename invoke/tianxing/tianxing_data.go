package tianxing

import (
	"github.com/tidwall/gjson"
	"io"
	"my-server-go/invoke"
	"my-server-go/tool"
)

// 私钥
const apiSecretKey = "66c7c0fdd54fa9e23b14937b77a29a0d"

// API
const basicUrl = "https://apis.tianapi.com"

// GoodMorningWords 早安心语
func GoodMorningWords() string {
	url := basicUrl + "/zaoan/index?key=" + apiSecretKey
	resp, err := invoke.SendGet(url, nil, nil)
	if err != nil {
		return ""
	}
	body, _ := io.ReadAll(resp.Body)
	return gjson.Get(string(body), "result.content").String()
}

func GoodNightWords() string {
	url := basicUrl + "/wanan/index?key=" + apiSecretKey
	resp, err := invoke.SendGet(url, nil, nil)
	if err != nil {
		return ""
	}
	body, _ := io.ReadAll(resp.Body)
	return gjson.Get(string(body), "result.content").String()
}

// InspirationalQuotes 励志名言
func InspirationalQuotes() string {
	url := basicUrl + "/lzmy/index?key=" + apiSecretKey
	resp, err := invoke.SendGet(url, nil, nil)
	if err != nil {
		return ""
	}
	body, _ := io.ReadAll(resp.Body)
	return gjson.Get(string(body), "result.saying").String()
}

// Holidays 节假日
func Holidays() string {
	url := basicUrl + "/jiejiari/index"
	params := make(map[string]string)
	params["key"] = apiSecretKey
	params["date"] = tool.GetSystemCurrentDate()
	params["mode"] = "1"
	resp, err := invoke.SendGet(url, params, nil)
	if err != nil {
		return ""
	}
	body, _ := io.ReadAll(resp.Body)
	result := gjson.Get(string(body), "result.list").Array()[0]

	return "今天是：" + result.Get("date").String() + " " + result.Get("cnweekday").String() + "\n" +
		"农历日期：" + result.Get("lunaryear").String() + " " + result.Get("lunarmonth").String() + " " + result.Get("lunarday").String() + "\n" +
		"温馨提示：" + result.Get("tip").String() + "\n" +
		"调休日期：" + result.Get("remark").String() + "\n" +
		"节假日期：" + result.Get("vacation").String() + "\n" +
		"日期类型：" + result.Get("info").String() + "\n" +
		"节日名称：" + result.Get("name").String() + "\n" +
		"拼假建议：" + result.Get("rest").String()

}

// SentenceOfTheDay 每日一句
func SentenceOfTheDay() string {
	url := basicUrl + "/one/index?key=" + apiSecretKey
	resp, err := invoke.SendGet(url, nil, nil)
	if err != nil {
		return ""
	}
	body, _ := io.ReadAll(resp.Body)
	return gjson.Get(string(body), "result.word").String()
}
