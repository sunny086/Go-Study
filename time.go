package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//TimeAdd()

}

// TimeAdd 时间加减
func TimeAdd() {
	// Add 时间相加
	now := time.Now()
	// ParseDuration parses a duration string.
	// A duration string is a possibly signed sequence of decimal numbers,
	// each with optional fraction and a unit suffix,
	// such as "300ms", "-1.5h" or "2h45m".
	//  Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// 1分钟前
	m, _ := time.ParseDuration("-1m")
	m1 := now.Add(m)
	fmt.Println(m1.Format("2006-01-02 15:04:05"))

	// 1个小时前
	h, _ := time.ParseDuration("-1h")
	h1 := now.Add(h)
	fmt.Println(h1.Format("2006-01-02 15:04:05"))

	// 一天前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println(d1.Format("2006-01-02 15:04:05"))

	printSplit(50)

	// 1分钟后
	mm, _ := time.ParseDuration("1m")
	mm1 := now.Add(mm)
	fmt.Println(mm1.Format("2006-01-02 15:04:05"))

	// 1小时后
	hh, _ := time.ParseDuration("1h")
	hh1 := now.Add(hh)
	fmt.Println(hh1.Format("2006-01-02 15:04:05"))

	// 一天后
	dd, _ := time.ParseDuration("24h")
	dd1 := now.Add(dd)
	fmt.Println(dd1.Format("2006-01-02 15:04:05"))

	printSplit(50)

	// Sub 计算两个时间差
	subM := now.Sub(m1)
	fmt.Println(subM.Minutes(), "分钟")

	sumH := now.Sub(h1)
	fmt.Println(sumH.Hours(), "小时")

	sumD := now.Sub(d1)
	fmt.Printf("%v 天\n", sumD.Hours()/24)
}

func printSplit(count int) {
	fmt.Println(strings.Repeat("#", count))
}
