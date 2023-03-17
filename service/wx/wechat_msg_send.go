package wx

import (
	"my-server-go/config/mysql"
	"my-server-go/invoke/tianxing"
	"my-server-go/invoke/wx"
	"my-server-go/invoke/xinzhi"
	logger "my-server-go/tool/log"
)

func SendMessageEveryMorning() {
	//早安心语
	goodMorningWords := tianxing.GoodMorningWords()
	//励志名言
	inspirationalQuotes := tianxing.InspirationalQuotes()
	//节假日
	holidays := tianxing.Holidays()
	//每日一句
	sentenceOfTheDay := tianxing.SentenceOfTheDay()

	//天气实况数据
	//获得位置信息
	db := mysql.Connect()
	var userLocation mysql.QywxUserLocation
	db.Where("user_name", "LiHongWei").Find(&userLocation)

	//查询天气,获取现在天气数据
	weatherNow := xinzhi.GetWeatherNow(userLocation.UserLocation)

	//获取逐日天气预报-今日
	weatherDaily := xinzhi.GetWeatherDaily(userLocation.UserLocation)
	weatherDailyToDay := weatherDaily["toDay"]

	//数据更新时间
	weatherDailyLastUpdate := weatherDaily["last_update"]

	//获取生活指数
	lifeSuggestion := xinzhi.GetLifeSuggestion(userLocation.UserLocation)

	//组装消息体并发送
	var message = "【早安】" + "\n" +
		"\n" +
		goodMorningWords + "\n" +
		"\n" +
		holidays + "\n" +
		"地点：" + weatherNow["path"] + "\n" +
		"白天天气☀️：" + weatherDailyToDay["text_day"] + "\n" +
		"晚间天气：" + weatherDailyToDay["text_night"] + "\n" +
		"最高温度℃：" + weatherDailyToDay["high"] + "℃" + "\n" +
		"最低温度℃：" + weatherDailyToDay["low"] + "℃" + "\n" +
		"降水概率☔️：" + weatherDailyToDay["precip"] + "%\n" +
		"降水量☔️：" + weatherDailyToDay["rainfall"] + "mm\n" +
		"风速：" + weatherDailyToDay["wind_speed"] + "km/h\n" +
		"风力等级：" + weatherDailyToDay["wind_scale"] + "级\n" +
		"相对湿度：" + weatherDailyToDay["humidity"] + "%\n" +
		"数据更新时间：" + weatherDailyLastUpdate["last_update"] + "\n" +
		"紫外线强弱：" + lifeSuggestion["uv"] + "\n" +
		"穿衣指标：" + lifeSuggestion["dressing"] + "\n" +
		"是否容易感冒：" + lifeSuggestion["flu"] + "\n" +
		"是否适合旅游：" + lifeSuggestion["travel"] + "\n" +
		"是否适合运动：" + lifeSuggestion["sport"] + "\n" +
		"是否适合洗车：" + lifeSuggestion["car_washing"] + "\n" +
		"数据更新时间：" + lifeSuggestion["last_update"] + "\n" +
		"\n" +
		"当前天气：" + weatherNow["text"] + "\n" +
		"当前温度℃：" + weatherNow["temperature"] + "℃" + "\n" +
		"数据更新时间：" + weatherNow["last_update"] + "\n" +
		"\n" +
		"今日格言：" + inspirationalQuotes + "\n" +
		"\n" +
		"每日一句：" + sentenceOfTheDay + "\n" +
		""
	code := wx.SendWxMessage(message)
	if code == 0 {
		logger.Write("SendMessageEveryMorning 消息发送成功!")
	} else {
		logger.Write("SendMessageEveryMorning 消息发送失败!")
	}
}