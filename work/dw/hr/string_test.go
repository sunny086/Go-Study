package hr

import (
	"fmt"
	"strings"
	"testing"
)

// TestSplit 如果字符串为空，Split 会返回一个只有一个元素的切片，该元素的值是空字符串
func TestSplit(t *testing.T) {
	str := ""
	strs := strings.Split(str, ",")
	fmt.Println(strs)
}
