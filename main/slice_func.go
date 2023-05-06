package main

import "fmt"

func test(s []int) {
	fmt.Printf("test---%p\n", s) // 打印与main函数相同的地址
	s = append(s, 1, 2, 3, 4, 5)
	//%p	十六进制表示的一个地址值
	fmt.Printf("test---%p\n", s) // 一旦append的数据超过切片长度，则会打印新地址
	fmt.Println("test---", s)    // [0 0 0 1 2 3 4 5]
}

func main() {

	s1 := make([]int, 3)
	test(s1)
	fmt.Printf("main---%p\n", s1) // 不会因为test函数内的append而改变
	fmt.Println("main---", s1)    // [ 0 0 0]
}
