package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

func main() {
	//time.now()
	nowTime := time.Now()
	//年月日
	year, month, day := nowTime.Date()
	fmt.Println(year, month, day) // 2019 November 01
	//时分秒
	hour, min, sec := nowTime.Clock()
	fmt.Println(hour, min, sec)
	//年月日 时分秒
	fmt.Println(nowTime.Year())
	fmt.Println(nowTime.Month())
	fmt.Println(nowTime.Hour())
	// 指今年一共过了多少天
	fmt.Println(nowTime.YearDay())

	fmt.Println("--------------------------")

	const timeLayout = "2006-01-02 15:04:05"

	rfc3339, err := time.Parse(time.RFC3339, "2022-03-02T14:24:59Z")
	fmt.Println(err)
	fmt.Println(rfc3339)
	format := rfc3339.Format(timeLayout)
	fmt.Println(format)
	fmt.Println("--------------------------")
	now := time.Now()
	s := now.Format(timeLayout)
	fmt.Println(s)
	fmt.Println("--------------------------")
	//var b = []byte("2022-03-02T14:24:59Z")
	b, _ := time.Now().MarshalJSON()
	s2 := string(b[:])
	fmt.Println(s2)

	////当前时间
	//nowTime := time.Now(); //	nowTime :=	time.Now().UTC() 协调时间;
	//fmt.Printf("%v\n", nowTime)
	//fmt.Printf("%02d.%02d.%4d\n", nowTime.Year(), nowTime.Month(), nowTime.Year())
	////转成普通的 yyyy-MM-dd 普通时间 没有像java那样的格式 要死记 总计 1234567
	//nowTimeStr := nowTime.Format("2006-01-02 15:04:05")
	//fmt.Printf("%v\n", nowTimeStr)
	////转成时间戳 然后新增一天
	//addTime := nowTime.Unix() + 3600*24;
	////将时间戳转化为 Time
	//tm := time.Unix(addTime, 0)
	//fmt.Printf("addTime: %v\n", time.Unix(addTime, 0).Format("2006-01-02 15:04:05"))
	//// 判断两个时间的大小
	//isTrue := tm.After(nowTime)
	//fmt.Printf("%v\n", isTrue)
	////计算两个时间的时间差
	//subTime := tm.Sub(nowTime)
	//fmt.Printf("时间差毫秒 %v ,秒 %v\n", subTime.Milliseconds(), subTime.Seconds())
	//// 10分钟前
	//m, _ := time.ParseDuration("-10m") // -1h
	//nowTimeMin := nowTime.Add(m)
	//fmt.Printf("nowTimeMin %v\n ", nowTimeMin)
	////10分钟后
	//tenAfter, _ := time.ParseDuration("10m");
	//nowTimetenAfter := nowTime.Add(tenAfter)
	//fmt.Printf("nowTimetenAfter %v\n ", nowTimetenAfter)

	var str = "2022-04-23T11:20:47+08:00"
	s3 := string(str)
	fmt.Println(s3)
}

const timeLayout = "2006-01-02 15:04:05"

type FormatTime struct {
	time.Time
}

func (t *FormatTime) String() string {
	return t.Format(timeLayout)
}

func (t *FormatTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(timeLayout, timeStr)
	*t = FormatTime{t1}
	return err
}

func (t FormatTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", t.Time.Format(timeLayout))
	return []byte(formatted), nil
}

func (t FormatTime) Value() (driver.Value, error) {
	// FormatTime 转换成 time.Time 类型
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}

	return t.Time.Format(timeLayout), nil
}

func (t *FormatTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = FormatTime{vt}
	default:
		return errors.New("类型处理错误")
	}
	return nil
}
