package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"testing"
	"time"
)

func TestCronParse(t *testing.T) {
	GenerateCornExecuteTime("0 0 31 * *")
	//// 生成cron表达式
	//spec := "0 0 28-31 * *"
	//expr := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow).MustParse(spec)
	//fmt.Println(expr)
}

// GenerateCornExecuteTime 测试cron表达式的执行时间
func GenerateCornExecuteTime(cronStr string) {
	// 解析cron表达式
	specParser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	schedule, err := specParser.Parse(cronStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印下一次执行时间
	now := time.Now()
	for i := 0; i < 20; i++ {
		nextTime := schedule.Next(now)
		fmt.Println(nextTime)
		now = nextTime
	}
}
