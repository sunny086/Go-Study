package main

import "fmt"

func main() {

	slice1 := make([]int, 5, 10)
	fmt.Println(len(slice1)) // 5
	fmt.Println(cap(slice1)) // 10
	fmt.Println(slice1)      // [0 0 0 0 0]

	slice1 = append(slice1, 1, 2)
	fmt.Println(slice1)

	sliceTemp := make([]int, 3)
	slice1 = append(slice1, sliceTemp...)
	fmt.Println(slice1)

	s1 := []int{1, 3, 6, 9}
	s2 := make([]int, 10) //必须给与充足的空间
	num := copy(s2, s1)

	fmt.Println(s1)  //[1 3 6 9]
	fmt.Println(s2)  //[1 3 6 9 0 0 0 0 0 0]
	fmt.Println(num) //4

}
