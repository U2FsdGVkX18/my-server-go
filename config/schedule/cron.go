package schedule

import (
	"github.com/robfig/cron/v3"
	"my-server-go/config/mysql"
	"my-server-go/invoke/douban"
	"my-server-go/invoke/notion"
	wx2 "my-server-go/invoke/wx"
	"my-server-go/invoke/xinzhi/business"
	business2 "my-server-go/service/business"
	"my-server-go/service/wx"
	logger "my-server-go/tool/log"
)

func Job() {
	//从数据库中获取cron表达式
	//初始化(秒级别,并增加错误回调函数)
	c := cron.New(cron.WithSeconds(), cron.WithChain(cron.Recover(cron.DefaultLogger)))
	//配置定时任务1
	var cron1 string
	mysql.DB.Model(mysql.Scheduled{}).Select("cron").Where("id = ?", 1).Limit(1).Scan(&cron1)
	_, err := c.AddJob(cron1, &everyMorning{})
	if err != nil {
		logger.Write("EveryMorning定时任务执行err", err)
	}
	//配置定时任务2
	var cron2 string
	mysql.DB.Model(mysql.Scheduled{}).Select("cron").Where("id = ?", 2).Limit(1).Scan(&cron2)
	_, err = c.AddJob(cron2, &everyHour{})
	if err != nil {
		logger.Write("EveryHour定时任务执行err", err)
	}
	//配置定时任务3
	var cron3 string
	mysql.DB.Model(mysql.Scheduled{}).Select("cron").Where("id = ?", 3).Limit(1).Scan(&cron3)
	_, err = c.AddJob(cron3, &everyNight{})
	if err != nil {
		logger.Write("EveryNight定时任务执行err", err)
	}
	//配置定时任务4
	var cron4 string
	mysql.DB.Model(mysql.Scheduled{}).Select("cron").Where("id = ?", 4).Limit(1).Scan(&cron4)
	_, err = c.AddJob(cron4, &everyDayZero{})
	if err != nil {
		logger.Write("EveryDayZero定时任务执行err", err)
	}
	//配置定时任务5
	var cron5 string
	mysql.DB.Model(mysql.Scheduled{}).Select("cron").Where("id = ?", 5).Limit(1).Scan(&cron5)
	_, err = c.AddJob(cron5, &everyWeekOne{})
	if err != nil {
		logger.Write("EveryWeekZero定时任务执行err", err)
	}
	//配置定时任务6
	var cron6 string
	mysql.DB.Model(mysql.Scheduled{}).Select("cron").Where("id = ?", 6).Limit(1).Scan(&cron6)
	_, err = c.AddJob(cron6, &everyMonthTwo{})
	if err != nil {
		logger.Write("EveryMonthZero定时任务执行err", err)
	}
	//配置定时任务7
	_, err = c.AddJob("0 0/30 * * * ?", &businessEveryHour{})
	if err != nil {
		logger.Write("BusinessEveryHour定时任务执行err", err)
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

type businessEveryHour struct{}

func (businessEveryHour *businessEveryHour) Run() {
	business.GetAllCityWeatherInsertDB()
	business2.GetRainCityForMysql()
}
