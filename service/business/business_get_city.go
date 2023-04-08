package business

import (
	"encoding/json"
	"my-server-go/config/mysql"
	"my-server-go/config/redis"
	logger "my-server-go/tool/log"
	"time"
)

type Result struct {
	Area     string
	Province string
	CityName string
}

// GetRainCityForMysql 从DB中获取正在下雨的城市并插入redis
func GetRainCityForMysql() {
	db := mysql.Connect()
	var businessCityList []mysql.BusinessCityList
	db.Select("city_id").Find(&businessCityList)
	var citys []Result
	for _, v := range businessCityList {
		var result Result
		err := db.Model(&mysql.BusinessCityWeather{}).Select("area,province,city_name").
			Where("city_id = ? AND weather_now LIKE ?", v.CityId, "%雨%").
			Limit(1).Scan(&result).Error
		if err != nil {
			logger.Write(err)
			continue
		}
		citys = append(citys, result)
	}
	marshal, _ := json.Marshal(citys)
	//插入redis
	redis.SetValue("businessRainCity", marshal, 2400*1000*time.Millisecond)
	logger.Write("businessRainCity数据写入redis完成")
}

// GetRainCityForRedis 从redis中获取正在下雨的城市并返回给接口
func GetRainCityForRedis() []string {
	value := redis.GetValue("businessRainCity")
	data := make([]string, 0)
	err := json.Unmarshal([]byte(value), &data)
	if err != nil {
		return data
	}
	return data
}
