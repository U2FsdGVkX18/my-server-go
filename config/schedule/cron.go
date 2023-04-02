package schedule

import (
	"github.com/robfig/cron/v3"
	"my-server-go/config/mysql"
	"my-server-go/invoke/douban"
	"my-server-go/invoke/notion"
	wx2 "my-server-go/invoke/wx"
	"my-server-go/invoke/xinzhi/business"
	"my-server-go/service/wx"
	logger "my-server-go/tool/log"
)

func Job() {
	//从数据库中获取cron表达式
	db := mysql.Connect()
	//初始化(秒级别,并增加错误回调函数)
	c := cron.New(cron.WithSeconds(), cron.WithChain(cron.Recover(cron.DefaultLogger)))
	//配置定时任务1
	var scheduled1 mysql.Scheduled
	db.Select("cron").Where("id = ?", 1).Find(&scheduled1)
	_, err := c.AddJob(scheduled1.Cron, &everyMorning{})
	if err != nil {
		logger.Write("EveryMorning定时任务执行err", err)
	}
	//配置定时任务2
	var scheduled2 mysql.Scheduled
	db.Select("cron").Where("id = ?", 2).Find(&scheduled2)
	_, err = c.AddJob(scheduled2.Cron, &everyHour{})
	if err != nil {
		logger.Write("EveryHour定时任务执行err", err)
	}
	//配置定时任务3
	var scheduled3 mysql.Scheduled
	db.Select("cron").Where("id = ?", 3).Find(&scheduled3)
	_, err = c.AddJob(scheduled3.Cron, &everyNight{})
	if err != nil {
		logger.Write("EveryNight定时任务执行err", err)
	}
	//配置定时任务4
	var scheduled4 mysql.Scheduled
	db.Select("cron").Where("id = ?", 4).Find(&scheduled4)
	_, err = c.AddJob(scheduled4.Cron, &everyDayZero{})
	if err != nil {
		logger.Write("EveryDayZero定时任务执行err", err)
	}
	//配置定时任务5
	var scheduled5 mysql.Scheduled
	db.Select("cron").Where("id = ?", 5).Find(&scheduled5)
	_, err = c.AddJob(scheduled5.Cron, &everyWeekOne{})
	if err != nil {
		logger.Write("EveryWeekZero定时任务执行err", err)
	}
	//配置定时任务6
	var scheduled6 mysql.Scheduled
	db.Select("cron").Where("id = ?", 6).Find(&scheduled6)
	_, err = c.AddJob(scheduled6.Cron, &everyMonthTwo{})
	if err != nil {
		logger.Write("EveryMonthZero定时任务执行err", err)
	}

	c.Start()
}

type everyMorning struct{}

func (everyMorning *everyMorning) Run() {
	wx.SendMessageEveryMorning()
}

type everyHour struct{}

func (everyHour *everyHour) Run() {
	wx.SendMessageEveryHour()
	business.GetAllCityWeatherInsertDB()
}

type everyNight struct{}

func (everyNight *everyNight) Run() {
	wx.SendMessageEveryNight()
}

type everyDayZero struct{}

func (everyDayZero *everyDayZero) Run() {
	douban.GetNewMovieRanking()
	douban.GetMovieNowShowing()
	douban.GetMovieComingSoon()
	var message1 = "【豆瓣爬虫】" + "\n" +
		"\n" +
		"新片电影排行数据spider结束" + "|" +
		"正在上映电影数据spider结束" + "|" +
		"即将上映电影数据spider结束" + "\n" +
		""
	wx2.SendWxMessage(message1)
	notion.SyncNewMovieRanking()
	notion.SyncMovieNowShowing()
	notion.SyncMovieComingSoon()
	var message2 = "【数据同步】" + "\n" +
		"\n" +
		"新片电影排行数据同步notion结束" + "|" +
		"正在上映电影数据同步notion结束" + "|" +
		"即将上映电影数据同步notion结束" + "\n" +
		""
	wx2.SendWxMessage(message2)
}

type everyWeekOne struct{}

func (everyWeekOne *everyWeekOne) Run() {
	douban.GetHighScoreTVShowRanking()
	douban.GetHotTestPublishBookRanking()
	douban.GetHighSalesPublishBookRanking()
	douban.GetHotTestOriginalBookRanking()
	var message1 = "【豆瓣爬虫】" + "\n" +
		"\n" +
		"高分电视剧排行数据spider结束" + "|" +
		"出版书籍中热度最高排行数据spider结束" + "|" +
		"出版出版书籍中销量最高排行数据spider结束" + "|" +
		"出版原创书籍中热度最高排行数据spider结束" + "\n" +
		""
	wx2.SendWxMessage(message1)
	notion.SyncHighScoreTVShowRanking()
	notion.SyncHotTestPublishBookRanking()
	notion.SyncHighSalesPublishBookRanking()
	notion.SyncHotTestOriginalBookRanking()
	var message2 = "【数据同步】" + "\n" +
		"\n" +
		"高分电视剧排行数据同步notion结束" + "|" +
		"出版书籍中热度最高排行数据同步notion结束" + "|" +
		"出版出版书籍中销量最高排行数据同步notion结束" + "|" +
		"出版原创书籍中热度最高排行数据同步notion结束" + "\n" +
		""
	wx2.SendWxMessage(message2)
}

type everyMonthTwo struct{}

func (everyMonthTwo *everyMonthTwo) Run() {
	douban.GetTop250MovieRanking()
	douban.GetTop250BookRanking()
	var message1 = "【豆瓣爬虫】" + "\n" +
		"\n" +
		"TOP250电影数据spider结束" + "|" +
		"TOP250读书数据spider结束" + "\n" +
		""
	wx2.SendWxMessage(message1)
	notion.SyncTop250MovieRanking()
	notion.SyncTop250BookRanking()
	var message2 = "【数据同步】" + "\n" +
		"\n" +
		"TOP250电影数据同步notion结束" + "|" +
		"TOP250读书数据同步notion结束" + "\n" +
		""
	wx2.SendWxMessage(message2)
}
