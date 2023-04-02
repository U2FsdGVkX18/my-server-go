package business

import (
	"my-server-go/config/mysql"
	logger "my-server-go/tool/log"
)

// GetRainCity 从DB中获取正在下雨的城市并返回数组
func GetRainCity() []string {
	db := mysql.Connect()
	var businessCityList []mysql.BusinessCityList
	db.Select("city_id").Find(&businessCityList)
	var citys []string
	for _, v := range businessCityList {
		var cityName string
		err := db.Model(&mysql.BusinessCityWeather{}).Select("city_name").
			Where("city_id = ? AND weather_now LIKE ?", v.CityId, "%雨%").
			Order("ABS(TIMESTAMPDIFF(SECOND, NOW(), created_at))").
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
	return citys
}
