package cronTask

import (
	"github.com/robfig/cron/v3"
	"log"
)

func CronTaskInit() {

	c := cron.New()
	c.AddFunc("*/1 * * * *", func() {
		log.Println("每分钟执行一次")
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("执行")
	})

	c.Start()

}
