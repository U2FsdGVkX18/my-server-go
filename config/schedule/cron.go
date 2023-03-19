package schedule

import (
	"github.com/robfig/cron/v3"
	"my-server-go/service/wx"
	logger "my-server-go/tool/log"
)

func Job() {
	c := cron.New(cron.WithSeconds(), cron.WithChain(cron.Recover(cron.DefaultLogger)))
	//spec := "0 0 8 * * ?"
	spec := "0 0/1 * * * ? "
	_, err := c.AddJob(spec, &everyMorning{})
	if err != nil {
		logger.Write("everyMorning定时任务执行err", err)
	}
	c.Start()
}

type everyMorning struct{}

func (everyMorning *everyMorning) Run() {
	wx.SendMessageEveryMorning()
}

//type myjob2 struct{}
//
//func (j *myjob2) Run() {
//	return
//}
