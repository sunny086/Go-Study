package main

import "fmt"

func main() {

	//sliceAppend()
	//sliceCopy()
	//sliceStartEnd()
	//sliceChangeAfterAppend()
}

func sliceChangeAfterAppend() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{-1, -2, -3}
	s3 := append(s1[:1], s2...)
	fmt.Println(s3)
	fmt.Println(s1)
}

// sliceStartEnd 含头不含尾
func sliceStartEnd() {
	s1 := []int{1, 3, 6, 9, 11, 33, 44, 55, 66, 77, 88, 99}
	slice := s1[3:6]
	fmt.Println(slice) //[9 11 33]
}

// sliceCopy 切片拷贝
func sliceCopy() {
	s1 := []int{1, 3, 6, 9}
	s2 := make([]int, 10) //必须给与充足的空间
	num := copy(s2, s1)

	fmt.Println(s1)  //[1 3 6 9]
	fmt.Println(s2)  //[1 3 6 9 0 0 0 0 0 0]
	fmt.Println(num) //4
}

// sliceAppend 切片拼接
func sliceAppend() {
	slice1 := make([]int, 5, 10)
	fmt.Println(len(slice1)) // 5
	fmt.Println(cap(slice1)) // 10
	fmt.Println(slice1)      // [0 0 0 0 0]

	slice1 = append(slice1, 1, 2)
	fmt.Println(slice1) //[0 0 0 0 0 1 2]

	sliceTemp := make([]int, 3)
	slice1 = append(slice1, sliceTemp...)
	fmt.Println(slice1) //[0 0 0 0 0 1 2 0 0 0]
}
