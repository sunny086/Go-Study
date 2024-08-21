package slice

import (
	"testing"
)

func TestReverse(t *testing.T) {
	source := []string{
		"10.25.10.1/24",   //10.25.10.1-10.25.10.254
		"10.25.10.10/24",  //10.25.10.1-10.25.10.254
		"10.25.11.128/25", //10.25.11.129-10.25.11.255
		"10.25.10.128/23", //10.25.10.1-10.25.11.254
	}
	//翻转切片
	Reverse(source)
	t.Log(source)
}

func Reverse(list []string) {
	//翻转数组
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}
