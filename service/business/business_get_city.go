package business

import (
	"encoding/json"
	"my-server-go/config/mysql"
	"my-server-go/config/redis"
	logger "my-server-go/tool/log"
	"time"
)

// GetRainCityForMysql 从DB中获取正在下雨的城市并插入redis
func GetRainCityForMysql() {
	db := mysql.Connect()
	var businessCityList []mysql.BusinessCityList
	db.Select("city_id").Find(&businessCityList)
	var citys []string
	for _, v := range businessCityList {
		var cityName string
		err := db.Model(&mysql.BusinessCityWeather{}).Select("city_name").
			Where("city_id = ? AND weather_now LIKE ?", v.CityId, "%雨%").
			Limit(1).Scan(&cityName).Error
		if err != nil {
			logger.Write(err)
			continue
		}
		if cityName == "" {
			continue
		}
		citys = append(citys, cityName)
	}
	marshal, _ := json.Marshal(citys)
	//插入redis
	redis.SetValue("businessRainCity", marshal, 2400*1000*time.Millisecond)
	logger.Write("businessRainCity数据写入redis完成")
}

// GetRainCityForRedis 从redis中获取正在下雨的城市并返回给接口
func GetRainCityForRedis() []string {
	value := redis.GetValue("businessRainCity")
	var data []string
	err := json.Unmarshal([]byte(value), &data)
	if err != nil {
		return data
	}
	return data
}
