package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	intChan := make(chan int, 100)
	stringChan := make(chan string, 50)

	wg.Add(1)
	go func() {
		defer wg.Done()
		//使用select来获取channel里面的数据的时候不需要关闭channel
		for {
			select {
			case v := <-intChan:
				fmt.Printf("从 intChan 读取的数据%d\n", v)
			case v := <-stringChan:
				fmt.Printf("从 stringChan 读取的数据%v\n", v)
			default:
				fmt.Printf("数据获取完毕")
				return //注意退出...
			}
		}
	}()
	wg.Add(1)
	go func() { //1.定义一个管道 10个数据int
		for i := 0; i < 100; i++ {
			intChan <- i
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		//2.定义一个管道 5个数据string
		for i := 0; i < 50; i++ {
			stringChan <- "hello" + fmt.Sprintf("%d", i)
		}
		wg.Done()
	}()
	wg.Wait()
}
