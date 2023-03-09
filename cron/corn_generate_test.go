package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeParse(t *testing.T) {
	//解析时分秒08:08:08
	t1, _ := time.Parse("15:04:05", "08:08:08")
	hour := t1.Hour()
	minute := t1.Minute()
	second := t1.Second()
	fmt.Println(hour, minute, second)
}

// TestCronGenerate1 时分秒参数生成cron表达式 然后生成执行时间
func TestCronGenerate1(t *testing.T) {
	dailyCron := generateDailyCron(02, 01, 01)
	GenerateCornExecuteTime(dailyCron)
}

// 时分秒参数生成cron表达式
func generateDailyCron(hour, minute, second int) string {
	spec := fmt.Sprintf("%d %d %d * * ?", second, minute, hour)
	return spec
}

// TestCronGenerate2 指定每周几的时分秒执行
func TestCronGenerate2(t *testing.T) {
	weeklyCron := generateWeeklyCron(1, 1, 1, 1)
	GenerateCornExecuteTime(weeklyCron)
}

func generateWeeklyCron(dayOfWeek, hour, minute, second int) string {
	spec := fmt.Sprintf("%d %d %d * * %d", second, minute, hour, dayOfWeek)
	return spec
}

func TestCronGenerate3(t *testing.T) {
	monthlyCron := generateMonthCron(1, 1, 1, 1)
	GenerateCornExecuteTime(monthlyCron)
}
func generateMonthCron(dayOfMonth, hour, minute, second int) string {
	spec := fmt.Sprintf("%d %d %d %d * ?", second, minute, hour, dayOfMonth)
	return spec
}

//// generateCronExpression 生成给定的时分秒年月日的 cron 表达式
//func generateCronExpression(second int, minute int, hour int, day int, month int, year int) (string, error) {
//	// 根据输入的参数生成一个 time.Time 对象
//	t := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
//
//	// 创建一个 Cron 对象
//	c := cron.New(cron.WithSeconds())
//
//	// 根据 time.Time 对象生成 cron 表达式
//	spec := c.SpecParser().Parse(t.Format("05 04 15 02 01 2006"))
//
//	// 返回 cron 表达式
//	return spec, nil
//}

//func Test2(t *testing.T) {
//	c := cron.New()
//
//	// 设置定时任务，每5秒执行一次
//	entryID, err := c.AddFunc("*/5 * * * *", func() {
//		t.Log("定时任务执行时间：", time.Now().Format("2006-01-02 15:04:05"))
//	})
//	c.Start()
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("定时任务已添加，EntryID：", entryID)
//
//	go func() {
//		time.Sleep(10 * time.Second)
//		// 修改定时任务，改为每10秒执行一次
//		_, err = c.AddFunc("*/10 * * * *", func() {
//			t.Log("定时任务修改后执行时间：", time.Now().Format("2006-01-02 15:04:05"))
//		})
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println("定时任务已修改，EntryID：", entryID)
//		c.Start()
//	}()
//
//	//c.Start()
//
//	//等待10秒钟后停止定时任务
//	time.Sleep(30 * time.Second)
//
//	c.Stop()
//}
