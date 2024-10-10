package dcp

import (
	"fmt"
	"testing"
	"time"
)

func TestLatestSixMonth01(t *testing.T) {
	var (
		xMonth = make([]string, 0)
		now, _ = time.Parse("2006-01-02", "2024-06-01")
		//now = time.Now()
	)

	for i := 6; i > 0; i-- {
		xMonth = append(xMonth, now.AddDate(0, -i, 0).Format("2006-01"))
	}

	t.Log(xMonth)
}
func TestLatestSixMonth02(t *testing.T) {
	var (
		AllMonth = []string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"}
		xMonth   = make([]string, 6)
		now, _   = time.Parse("2006-01-02", "2024-06-01")
		//now = time.Now()
		month = now.Month()
		year  = now.Year()
	)

	// 动态生成从上个月开始的最近六个月的横坐标
	if month > 6 {
		xMonth = AllMonth[month-6-1 : month-1]
		// 拼接年份
		for i, v := range xMonth {
			xMonth[i] = fmt.Sprintf("%d年%s", year, v)
		}
	} else {
		currentYearMonth := AllMonth[0 : month-1]

		left := 6 - (month - 1)

		lastYearMonth := AllMonth[12-left : 12]
		// 拼接年份
		for i, v := range lastYearMonth {
			xMonth[i] = fmt.Sprintf("%d年%s", year-1, v)
		}
		for i, v := range currentYearMonth {
			xMonth[len(lastYearMonth)+i] = fmt.Sprintf("%d年%s", year, v)
		}
	}
	t.Log(xMonth)
}
