package dw

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	unix := time.Now().Unix()
	t.Log(unix)
	t.Log(time.Unix(unix, 0).Format("2006-01-02 15:04:05"))
	milli := time.Now().UnixMilli()
	t.Log(milli)
	t.Log(time.UnixMilli(milli).Format("2006-01-02 15:04:05"))
}
func TestTime2(t *testing.T) {
	format := time.Now().Format("2006-01-02T15:04:05")
	t.Log(format)
	expectedTime, _ := time.Parse("2006-01-02", "2023-03-13")
	t.Log(expectedTime.Format("2006-01-02T15:04:05.000Z"))

}

func TestTime3(t *testing.T) {
	expectedTime, _ := time.Parse("2006-01-02 15:04:05", "2024-01-01 00:00:00")

	t.Log(expectedTime.Year())

	date := expectedTime.AddDate(0, 0, -1)

	t.Log(date.Year())

	t.Log(date.Year())
	t.Log(date.Month())

}

func TestSlice(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}

func TestArrStr(t *testing.T) {
	// 定义一个数组
	arr := []string{"item1", "item2", "item3"}
	if arr == nil {
		fmt.Println("arr is nil")
		return
	}

	// 将数组转换为 JSON 字符串
	arrJSON, err := json.Marshal(arr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arrJSON))

	//将JSON字符串转换为数组
	var arr2 []string
	err = json.Unmarshal(arrJSON, &arr2)
	if err != nil {
		panic(err)
	}
	fmt.Println(arr2)
}
