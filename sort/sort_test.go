package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortInts(t *testing.T) {
	numbers := []int{4, 2, 8, 5, 1, 3, 9, 6, 7}
	// 对切片进行排序
	sort.Ints(numbers)
	fmt.Println(numbers)
}

func TestSortStrings1(t *testing.T) {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println(strs)
}

func TestSortStrings2(t *testing.T) {
	strs := []string{"ETH1", "ETH3", "ETH0", "ETH2"}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})
	fmt.Println(strs)
}

func TestSortSlice(t *testing.T) {
	// 1. 定义一个结构体切片
	type person struct {
		Name string
		Age  int
	}
	// 2. 初始化一个结构体切片
	people := []person{
		{"张三", 18},
		{"李四", 20},
		{"王五", 19},
	}
	// 3. 对切片进行排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}
