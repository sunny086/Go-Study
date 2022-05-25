package main

import "fmt"

func main() {
	//声明一个int数组
	ids := []int{1, 2, 3, 4, 5}
	fmt.Println(ids)
	fmt.Printf("%T:%v\n", &ids)
	testPoint(&ids)
}

func testPoint(ids *[]int) {
	fmt.Printf("%T:%v\n", ids, ids)
	fmt.Printf("%T:%v\n", *ids, *ids)
}
