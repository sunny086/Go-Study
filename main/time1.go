package main

import (
	"fmt"
	"strings"
	"time"
)

var formatArr = []string{"2006-01-02 15:04:05", "2006-01-02T15:04:05", "2006-01-02T15:04:05+08:00", "2006-01-02",
	"2006-01-02 15:04:05.999999999 -0700 MST", "2006-01-02 15:04:05.999999999 -0700 -0700",
	"2006-01-02 15:04:05.999999999 -07:00", "2006-01-02 15:04:05.999999999 -0700",
	"2006-01-02T15:04:05Z", "Jan 02 15:04:05 2006 GMT", "1月 02 15:04:05 2006 GMT"}

// StrToDate 时间转换
func StrToDate(dateStr string) time.Time {
	dateStr = strings.TrimSpace(dateStr)
	if len(dateStr) == 0 {
		return time.Time{}
	}
	for _, format := range formatArr {
		dateTime, err := time.ParseInLocation(format, strings.TrimSpace(dateStr), time.Local)
		if err != nil {
			continue
		}
		return dateTime
	}
	return time.Time{}
}

func main() {
	date := StrToDate("2022-04-23T11:20:47+08:00")
	fmt.Println(date)

	var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhc2NvcGUiOiIxIiwiZXhwIjo0ODA0MjM2MTI0LCJpZGVudGl0eSI6MiwibmljZSI6Inh1amluc2hhbiIsIm9yaWdfaWF0IjoxNjUwNjAwMTI0LCJyb2xlaWQiOjEsInJvbGVrZXkiOiJhZG1pbiIsInJvbGVuYW1lIjoi6LaF57qn566h55CG5ZGYIn0.FLcxm28MfXdd2K4kcfYLKsiOpCcfKAbWbyaBr9FhWJQ"
	fmt.Println(len(token))

	unix := time.Now().Unix()
	fmt.Println(unix)

	fmt.Println(time.Now().UnixNano())

	//毫秒转时间
	fmt.Println(time.Unix(4804244484, 0).Format("2006-01-02 15:04:05"))
	fmt.Println("签发时间：", time.Unix(1650623212, 0).Format("2006-01-02 15:04:05"))
	fmt.Println("当前时间-最大刷新时间：", time.Unix(1650623878, 0).Format("2006-01-02 15:04:05"))

	fmt.Println("===============================")
	format := time.Now().Format("2006-01-02")
	fmt.Println(format)
}
