package main

import (
	"fmt"
	"sync"
)

var x int

func Increase() func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	//bibao_test1()

	//bibao_test2()
	bibao_test3()
}

//闭包测试1 闭包( Closure)在某些编程语言中也被称为 Lambda表达式（如Java） 函数+引用环境=闭包
func bibao_test1() {
	increase := Increase()
	println(increase())
	println(increase())
	println(increase())
}

//闭包测试2 并发中的闭包
func bibao_test2() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

//共享的环境变量作为函数参数传递
func bibao_test3() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
