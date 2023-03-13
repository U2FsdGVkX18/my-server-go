package schedule

import (
	"github.com/robfig/cron/v3"
	"my-server-go/service/wx"
)

func Cron() {
	c := cron.New(cron.WithSeconds())
	spec := "0 36 23 * * ?"
	_, err := c.AddFunc(spec, wx.SendMessageEveryMorning)
	if err != nil {
		return
	}
	c.Start()
}
