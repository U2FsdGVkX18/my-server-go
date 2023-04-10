package business

import (
	"encoding/json"
	"my-server-go/config/mysql"
	"my-server-go/config/redis"
	logger "my-server-go/tool/log"
	"time"
)

// Result 定义查询数据库部分字段的结构体
type Result struct {
	Area     string
	Province string
	CityName string
}

// GetRainCityForMysql 从DB中获取正在下雨的城市并插入redis
func GetRainCityForMysql() {
	var cityIds []string
	err := mysql.DB.Model(mysql.BusinessCityList{}).Select("city_id").Scan(&cityIds).Error
	if err != nil {
		logger.Write("GetRainCityForMysql:", err)
	}
	var citys []Result
	for _, v := range cityIds {
		var result Result
		err := mysql.DB.Model(&mysql.BusinessCityWeather{}).Select("area,province,city_name").
			Where("city_id = ? AND weather_now LIKE ?", v, "%雨%").
			Limit(1).Scan(&result).Error
		if err != nil {
			logger.Write(err)
			continue
		}
		//查询记录会返回空的结果,这里进行过滤
		if result.CityName == "" {
			continue
		}
		citys = append(citys, result)
	}
	marshal, _ := json.Marshal(citys)
	//插入redis
	redis.SetValue("businessRainCity", marshal, 2400*1000*time.Millisecond)
	logger.Write("businessRainCity数据写入redis完成")
}

// GetRainCityForRedis 接口调用,从redis中获取正在下雨的城市并返回给接口
func GetRainCityForRedis() []Result {
	value := redis.GetValue("businessRainCity")
	data := make([]Result, 0)
	err := json.Unmarshal([]byte(value), &data)
	if err != nil {
		return data
	}
	return data
}
