package dcp

import (
	"fmt"
	"testing"
	"time"
)

func TestDays(t *testing.T) {
	now := time.Now()
	currentYear := now.Year()
	daysInMonth := make([]int, 12)
	daysInMonthMap := make(map[int]int)

	// 遍历每个月，获取每个月的天数
	for month := time.January; month <= time.December; month++ {
		lastDayOfNextMonth := time.Date(currentYear, month+1, 0, 0, 0, 0, 0, time.UTC)
		daysInMonth[month-1] = lastDayOfNextMonth.Day()
		daysInMonthMap[int(month)] = lastDayOfNextMonth.Day()
	}
	// 打印每个月的天数
	for month, days := range daysInMonth {
		fmt.Printf("Month %d: %d days\n", month+1, days)
	}
}
