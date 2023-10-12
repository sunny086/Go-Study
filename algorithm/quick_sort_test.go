package algorithm

import (
	"fmt"
	"testing"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0] // 选择第一个元素作为基准
	var less, greater []int

	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	// 递归地对小于基准值和大于基准值的子数组进行排序
	less = quickSort(less)
	greater = quickSort(greater)

	// 将排序后的子数组与基准值组合在一起
	sortedArr := append(less, pivot)
	sortedArr = append(sortedArr, greater...)

	return sortedArr
}

func TestQuickSort(t *testing.T) {
	arr := []int{12, 11, 13, 5, 6, 7, 8, 4, 14}

	fmt.Println("排序前:", arr)
	sortedArr := quickSort(arr)
	fmt.Println("排序后:", sortedArr)
}
