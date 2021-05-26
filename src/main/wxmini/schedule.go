package wxmini

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func NewSchedule(spec string,f func()){
	c := cron.New()
	_, err := c.AddFunc(spec, func() {
		fmt.Printf("定时任务启动 %s \n", time.Now())
		f()
	})
	if err!=nil {
		fmt.Printf("err %s\n",err.Error())
	}
	fmt.Printf("定时任务	%v\n",c.Entries())
	c.Run()
}
