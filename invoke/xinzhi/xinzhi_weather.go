package xinzhi

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"my-server-go/invoke"
)

// 私钥
const apiSecretKey = "SVb6HuTfbwzu0pNrK"

// API
const basicUrl = "https://api.seniverse.com/v3"

// GetWeatherNow 获取天气实况(位置,天气,温度)
func GetWeatherNow(location string) map[string]string {
	url := basicUrl + "/weather/now.json?key=" + apiSecretKey + "&location=" + location
	resp := invoke.SendGet(url, nil, nil)

	//defer关闭io流
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, _ := io.ReadAll(resp.Body)
	//返回的是一个json进行解析
	result := gjson.Get(string(body), "results").Array()[0]
	fmt.Println(gjson.Get(string(body), "results").Array())
	//定义结果map
	var weatherNowMap = make(map[string]string)
	weatherNowMap["name"] = result.Get("location.name").String()
	weatherNowMap["path"] = result.Get("location.path").String()
	weatherNowMap["text"] = result.Get("now.text").String()
	weatherNowMap["temperature"] = result.Get("now.temperature").String()
	weatherNowMap["last_update"] = result.Get("last_update").String()
	//返回map
	return weatherNowMap
}

// GetWeatherDaily 获取逐日天气预报(免费版只能获取未来三天)
func GetWeatherDaily(location string) map[string]map[string]string {
	url := basicUrl + "/weather/daily.json?key=" + apiSecretKey + "&location=" + location
	resp := invoke.SendGet(url, nil, nil)

	//defer关闭io流
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, _ := io.ReadAll(resp.Body)
	//返回的是一个json进行解析
	result := gjson.Get(string(body), "results").Array()[0]
	//今日天气
	toDayWeatherData := result.Get("daily").Array()[0]
	var toDayWeatherMap = make(map[string]string)
	JsonToMapProcess(toDayWeatherData, toDayWeatherMap)
	//明日天气
	tomorrowWeatherData := result.Get("daily").Array()[1]
	var tomorrowWeatherMap = make(map[string]string)
	JsonToMapProcess(tomorrowWeatherData, tomorrowWeatherMap)
	//后天天气
	theDayAfterTomorrowWeatherData := result.Get("daily").Array()[2]
	var theDayAfterTomorrowWeatherMap = make(map[string]string)
	JsonToMapProcess(theDayAfterTomorrowWeatherData, theDayAfterTomorrowWeatherMap)
	//数据更新时间
	var weatherDailyLastUpdateMap = make(map[string]string)
	weatherDailyLastUpdateMap["last_update"] = result.Get("last_update").String()
	//将三天天气一起返回
	var collectMap = make(map[string]map[string]string)
	collectMap["toDay"] = toDayWeatherMap
	collectMap["tomorrow"] = tomorrowWeatherMap
	collectMap["theDayAfterTomorrow"] = theDayAfterTomorrowWeatherMap
	collectMap["last_update"] = weatherDailyLastUpdateMap

	return collectMap
}

// GetLifeSuggestion 获取生活指数
func GetLifeSuggestion(location string) map[string]string {
	url := basicUrl + "/life/suggestion.json?key=" + apiSecretKey + "&location=" + location
	resp := invoke.SendGet(url, nil, nil)

	//defer关闭io流
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, _ := io.ReadAll(resp.Body)
	//返回的是一个json进行解析
	result := gjson.Get(string(body), "results").Array()[0]
	var lifeSuggestionMap = make(map[string]string)
	lifeSuggestionMap["uv"] = result.Get("suggestion.uv.brief").String()
	lifeSuggestionMap["dressing"] = result.Get("suggestion.dressing.brief").String()
	lifeSuggestionMap["flu"] = result.Get("suggestion.flu.brief").String()
	lifeSuggestionMap["travel"] = result.Get("suggestion.travel.brief").String()
	lifeSuggestionMap["sport"] = result.Get("suggestion.sport.brief").String()
	lifeSuggestionMap["car_washing"] = result.Get("suggestion.car_washing.brief").String()
	lifeSuggestionMap["last_update"] = result.Get("last_update").String()

	return lifeSuggestionMap
}

func JsonToMapProcess(result gjson.Result, dataMap map[string]string) {
	dataMap["date"] = result.Get("date").String()
	dataMap["text_day"] = result.Get("text_day").String()
	dataMap["text_night"] = result.Get("text_night").String()
	dataMap["high"] = result.Get("high").String()
	dataMap["low"] = result.Get("low").String()
	dataMap["precip"] = fmt.Sprintf("%.1f", result.Get("precip").Float()*100)
	dataMap["rainfall"] = fmt.Sprintf("%.1f", result.Get("rainfall").Float()*100)
	dataMap["wind_speed"] = result.Get("wind_speed").String()
	dataMap["wind_scale"] = result.Get("wind_scale").String()
	dataMap["humidity"] = result.Get("humidity").String()
}
