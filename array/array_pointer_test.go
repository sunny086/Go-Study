package array

import (
	"fmt"
	"testing"
)

func TestArrayPointer(t *testing.T) {
	var arr = [5]int{1, 2, 3, 4, 5}
	modifyArr1(arr)
	fmt.Println(arr)
	modifyArr2(&arr)
	fmt.Println(arr)
}
func modifyArr1(a [5]int) {
	a[1] = 20
}

func modifyArr2(a *[5]int) {
	a[1] = 20
}
