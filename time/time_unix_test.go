package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeUnix(t *testing.T) {
	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（毫秒）：%v;\n", time.Now().UnixNano()/1e6)
	fmt.Printf("时间戳（纳秒）：%v;\n", time.Now().UnixNano())
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n", time.Now().UnixNano()/1e9)
}

func TestTimeAddDuration(t *testing.T) {
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
	fmt.Println("一分钟前：", m1.Format("2006-01-02 15:04:05"))

	// 1个小时前
	h, _ := time.ParseDuration("-1h")
	h1 := now.Add(h)
	fmt.Println("一小时前：", h1.Format("2006-01-02 15:04:05"))

	// 一天前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println("一天前：", d1.Format("2006-01-02 15:04:05"))

	fmt.Println("=========================================")

	// 1分钟后
	mm, _ := time.ParseDuration("1m")
	mm1 := now.Add(mm)
	fmt.Println("一分钟后", mm1.Format("2006-01-02 15:04:05"))

	// 1小时后
	hh, _ := time.ParseDuration("1h")
	hh1 := now.Add(hh)
	fmt.Println("一小时后", hh1.Format("2006-01-02 15:04:05"))

	// 一天后
	dd, _ := time.ParseDuration("24h")
	dd1 := now.Add(dd)
	fmt.Println("一天后", dd1.Format("2006-01-02 15:04:05"))

	fmt.Println("=========================================")

	// Sub 计算两个时间差
	subM := now.Sub(m1)
	fmt.Println(subM.Minutes(), "分钟")

	sumH := now.Sub(h1)
	fmt.Println(sumH.Hours(), "小时")

	sumD := now.Sub(d1)
	fmt.Printf("%v 天\n", sumD.Hours()/24)
}

// 当前时间转换为时间戳
func TestNow2Timestamp(t *testing.T) {
	fmt.Printf("当前的时间戳是：%v", time.Now().Unix())
}

// 指定时间转为时间戳
func TestAssignTim2Timestamp(t *testing.T) {
	str := "2023-02-24 00:00:00"
	tm, _ := time.Parse("2006-01-02 15:04:05", str)
	fmt.Printf(str+"的时间戳是：%v", tm.Unix())
}

// 时间戳转为时间
func TestTimestamp2Time(t *testing.T) {
	//redis中配置的令牌过期时间 15768000000 单位秒
	timestamp := int64(17443323226)
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj.Format("2006-01-02 15:04:05"))
}
