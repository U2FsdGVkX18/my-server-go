package schedule

import (
	"github.com/robfig/cron/v3"
	"my-server-go/invoke/douban"
	wx2 "my-server-go/invoke/wx"
	"my-server-go/service/wx"
	logger "my-server-go/tool/log"
)

func Job() {
	//初始化(秒级别,并增加错误回调函数)
	c := cron.New(cron.WithSeconds(), cron.WithChain(cron.Recover(cron.DefaultLogger)))
	//配置定时任务1
	EveryMorning := "0 0 8 * * ?"
	_, err := c.AddJob(EveryMorning, &everyMorning{})
	if err != nil {
		logger.Write("EveryMorning定时任务执行err", err)
	}
	//配置定时任务2
	EveryHour := "0 0/30 8-22 * * ?"
	_, err = c.AddJob(EveryHour, &everyHour{})
	if err != nil {
		logger.Write("EveryHour定时任务执行err", err)
	}
	//配置定时任务3
	EveryNight := "0 0 23 * * ?"
	_, err = c.AddJob(EveryNight, &everyNight{})
	if err != nil {
		logger.Write("EveryNight定时任务执行err", err)
	}
	//配置定时任务4
	EveryDayZero := "0 0 0 * * ?"
	_, err = c.AddJob(EveryDayZero, &everyDayZero{})
	if err != nil {
		logger.Write("EveryDayZero定时任务执行err", err)
	}
	//配置定时任务5
	EveryWeekZero := "0 27 15 * * ?"
	_, err = c.AddJob(EveryWeekZero, &everyWeekZero{})
	if err != nil {
		logger.Write("EveryWeekZero定时任务执行err", err)
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
	var message = "【豆瓣爬虫】" + "\n" +
		"\n" +
		"新片电影排行数据spider结束" + "|" +
		"正在上映电影数据spider结束" + "|" +
		"即将上映电影数据spider结束" + "\n" +
		""
	code := wx2.SendWxMessage(message)
	if code == 0 {
		logger.Write("everyDayZero 消息发送成功!")
	} else {
		logger.Write("everyDayZero 消息发送失败!", code)
	}
}

type everyWeekZero struct{}

func (everyWeekZero *everyWeekZero) Run() {
	douban.GetHighScoreTVShowRanking()
	douban.GetHotTestPublishBookRanking()
	douban.GetHighSalesPublishBookRanking()
	douban.GetHotTestOriginalBookRanking()
	var message = "【豆瓣爬虫】" + "\n" +
		"\n" +
		"高分电视剧排行数据spider结束" + "|" +
		"出版书籍中热度最高排行数据spider结束" + "|" +
		"出版出版书籍中销量最高排行数据spider结束" + "|" +
		"出版原创书籍中热度最高排行数据spider结束" + "\n" +
		""
	code := wx2.SendWxMessage(message)
	if code == 0 {
		logger.Write("everyWeekZero 消息发送成功!")
	} else {
		logger.Write("everyWeekZero 消息发送失败!", code)
	}
}
