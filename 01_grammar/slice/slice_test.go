package slice

import (
	"fmt"
	"testing"
)

func TestDiff(t *testing.T) {
	src := []string{"10.25.10.1", "10.25.10.2", "10.25.10.3", "10.25.10.4", "10.25.10.5", "10.25.10.6"}
	dest := []string{"10.25.10.1", "10.25.10.2"}
	differenceStrings := DifferenceStrings(src, dest)
	fmt.Println(differenceStrings)
}

func ContainsString(src []string, dest string) bool {
	for _, item := range src {
		if item == dest {
			return true
		}
	}
	return false
}

// DifferenceStrings 取前者src与后者dest两个字符串列表的差集
func DifferenceStrings(src []string, dest []string) []string {
	res := make([]string, 0)
	for _, item := range src {
		if !ContainsString(dest, item) {
			res = append(res, item)
		}
	}
	return res
}

// IntersectionStrings 取两个字符串列表的交集
func IntersectionStrings(src []string, dest []string) []string {
	res := make([]string, 0)
	for _, item := range src {
		if ContainsString(dest, item) {
			res = append(res, item)
		}
	}
	return res
}

// UnionString 取两个字符串列表的并集
func UnionStrings(src []string, dest []string) []string {
	res := make([]string, 0)
	res = append(res, src...)
	for _, item := range dest {
		if !ContainsString(res, item) {
			res = append(res, item)
		}
	}
	return res
}
