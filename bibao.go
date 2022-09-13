package main

import (
	"fmt"
	"runtime"
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
	gomaxprocs := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(gomaxprocs)
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
