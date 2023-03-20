package schedule

import (
	"github.com/robfig/cron/v3"
	"my-server-go/service/wx"
	logger "my-server-go/tool/log"
)

func Job() {
	//初始化
	c := cron.New(cron.WithSeconds(), cron.WithChain(cron.Recover(cron.DefaultLogger)))
	//配置定时任务1
	EveryMorning := "0 0 8 * * ?"
	_, err := c.AddJob(EveryMorning, &everyMorning{})
	if err != nil {
		logger.Write("EveryMorning定时任务执行err", err)
	}
	//配置定时任务2
	EveryHour := "0 0/30 8-23 * * ?"
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
