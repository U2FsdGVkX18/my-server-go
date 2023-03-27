package test

import (
	"fmt"
	"my-server-go/api"
	"my-server-go/config/mysql"
	"my-server-go/config/redis"
	"my-server-go/invoke/douban"
	"my-server-go/invoke/notion"
	"my-server-go/invoke/tianxing"
	"my-server-go/invoke/wx"
	"my-server-go/invoke/xinzhi"
	logger "my-server-go/tool/log"
	"testing"
)

func TestOne(t *testing.T) {
	var msg_signature = "bb65bc2127f86d862df8e6917daa7ef4a7b1733d"
	var timestamp = "1677596202"
	var nonce = "1676975221"
	var data = "<xml><ToUserName><![CDATA[ww8d5186f5aa839ee7]]></ToUserName><Encrypt><![CDATA[MMs94DR1S7fGleh0mKyjA3RNgwuNNGu0YyGimD+99GB16gCpuhBkvWxaTt20L1PC6Ni0VBWnpSdlpUsWseUbFpmsRtt8aFkTdoyRBe8C0gx9hM8bLrOWGcdOJrtXaGIUnOF8H8UuinQXLjO/uAulBKLKE7TiMFXtvaQ62/Iuzc5UKdh8bAbGUk+iOY1nUkh3L5BSPpyWHWVKEFyLkumjUCWZV4L11lSuG9nqbDVVFhdHLT/Du3TCX/To4DW7DIUyjgpzARVjAPzBzGvYYe1Nq1Y3RkjbwdWRWz824xhgmYEpiUr4XOYlfqTWljydXOV+NdNmJBXc/WDnG4u2jo1HsUjYRsWzaYqux4CX3dm1WI6L9iDJB87F5Ldp90yoxuf5rBdb3xLtssxBu8S4zievwlZVzRnQWN33Xvg0fUKHjo0=]]></Encrypt><AgentID><![CDATA[1000002]]></AgentID></xml>"
	api.ProcessMessage(msg_signature, timestamp, nonce, []byte(data))
}

func TestTwo(t *testing.T) {
	api.CreatePassiveRespText("lihongwei", "123123123", "nonce", "1")
}

func Test3(t *testing.T) {
	token := wx.GetAccessToken()
	fmt.Println(token)
}

func Test4(t *testing.T) {
	wx.SendWxMessage("2312312312")
}

func Test5(t *testing.T) {
	db := mysql.Connect()
	//var qywx mysql.QywxUserLocation
	//result := db.Where("user_name", "lihongwei").First(&mysql.QywxUserLocation{})
	//result := db.Where("user_name", "LiHongWei").Find(&mysql.QywxUserLocation{})
	//fmt.Println(result.RowsAffected)
	//db.Create(&mysql.QywxUserLocation{
	//	UserName:     "lihongwei",
	//	UserLocation: "123",
	//})
	//db.Save(&mysql.QywxUserLocation{
	//	UserName:     "lihongwei",
	//	UserLocation: "location",
	//})
	result := db.Model(&mysql.QywxUserLocation{}).Where("user_name", "LiHongWei").Update("user_location", "asdasd")

	//result := db.Where("user_name", "lihongwei").Update("user_location", "1211111:22")
	fmt.Println(result.RowsAffected)
}

func Test6(t *testing.T) {

	daily := xinzhi.GetWeatherDaily("30.292601:120.039001")
	fmt.Println(daily)
}

func Test7(t *testing.T) {
	str := tianxing.Holidays()
	fmt.Println(str)
}

func Test8(t *testing.T) {
	douban.GetHighScoreTVShowRanking()
}

func Test9(t *testing.T) {
	//redis.SetValue("key", "token", 7200*1000*time.Millisecond)
	//value := redis.GetValue("wxAccessToken")
	//fmt.Println(value)
	redis.SetValue("wxAccessToken", "123", 7200*1000*1000)
	value := redis.GetValue("wxAccessToken")

	fmt.Println("redis:", value)
}

func Test10(t *testing.T) {
	//wx2.SendMessageEveryMorning()
	logger.Write("123123123")
}

func Test11(t *testing.T) {
	mysql.CreateTables()
}

func Test12(t *testing.T) {
	db := mysql.Connect()
	var sch mysql.Scheduled
	var cron string
	db.Select("Cron").Where("id = ?", 1).First(&sch).Scan(&cron)
	db.Select("Cron").Where("id = ?", 1).First(&mysql.Scheduled{}).Scan(&cron)
	fmt.Println(cron)
}

func Test13(t *testing.T) {
	notion.DeleteDataBaseData("7ec1b96bd47f4c5c86fab914c24b3c73")
}
