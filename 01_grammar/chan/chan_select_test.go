package _chan

import (
	"fmt"
	"testing"
	"time"
)

// 在某些场景下我们需要同时从多个通道接收数据,这个时候就可以用到golang中给我们提供的select多路复用
func TestChanSelect1(t *testing.T) {
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

// 在某些场景下我们需要同时从多个通道接收数据,这个时候就可以用到golang中给我们提供的select多路复用
func TestChanSelect2(t *testing.T) {
	// 在某些场景下我们需要同时从多个通道接收数据,这个时候就可以用到golang中给我们提供的select多路复用
	//1.定义一个管道 10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//2.定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
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
}

func TestChanSelect3(t *testing.T) {
	// 创建管道
	ch := make(chan string, 10)
	// 子协程写数据
	go func() {
		for {
			select {
			// 写数据
			case ch <- "hello":
				fmt.Println("write hello")
			default:
				fmt.Println("channel full")
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	// 取数据
	for s := range ch {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}
