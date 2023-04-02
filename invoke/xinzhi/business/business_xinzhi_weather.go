package business

import (
	"errors"
	"github.com/tidwall/gjson"
	"io"
	"my-server-go/config/mysql"
	"my-server-go/invoke"
	logger "my-server-go/tool/log"
	"time"
)

// 私钥
const apiSecretKey = "SVb6HuTfbwzu0pNrK"

// API
const basicUrl = "https://api.seniverse.com/v3"

// GetAllCityWeatherInsertDB GetAllCityWeather 获取每个城市的天气数据插入并插入到数据库中
func GetAllCityWeatherInsertDB() {
	db := mysql.Connect()
	var businessCityList []mysql.BusinessCityList
	db.Select("city_id,city_name").Find(&businessCityList)
	for _, v := range businessCityList {
		time.Sleep(3 * time.Second)
		weatherNow, err := GetWeatherNow(v.CityId)
		if err != nil {
			db.Create(&mysql.BusinessCityWeather{
				CityName: v.CityName,
			})
			logger.Write("GetAllCityWeather:", v.CityName, err)
			continue
		} else {
			db.Create(&mysql.BusinessCityWeather{
				CityId:         v.CityId,
				CityName:       v.CityName,
				WeatherNow:     weatherNow["text"],
				TemperatureNow: weatherNow["temperature"],
				DataUpdate:     weatherNow["last_update"],
			})
		}
	}
}

// GetWeatherNow 获取天气实况(位置,天气,温度)
func GetWeatherNow(location string) (map[string]string, error) {
	url := basicUrl + "/weather/now.json?key=" + apiSecretKey + "&location=" + location
	resp := invoke.SendGet(url, nil, nil)
	//defer关闭io流
	defer resp.Body.Close()
	//获取body
	body, _ := io.ReadAll(resp.Body)
	statusCode := gjson.Get(string(body), "status_code").String()
	status := gjson.Get(string(body), "status").String()
	if statusCode == "AP010014" || statusCode == "AP010006" {
		var weatherNowMap = make(map[string]string)
		return weatherNowMap, errors.New(status)
	}
	//返回的是一个json进行解析
	result := gjson.Get(string(body), "results").Array()[0]
	//定义结果map
	var weatherNowMap = make(map[string]string)
	weatherNowMap["text"] = result.Get("now.text").String()
	weatherNowMap["temperature"] = result.Get("now.temperature").String()
	weatherNowMap["last_update"] = result.Get("last_update").String()
	//返回map
	return weatherNowMap, nil
}
