package _chan

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestChanForRange(t *testing.T) {
	//使用for range遍历通道，当通道被关闭的时候就会退出for range,如果没有关闭管道就会报个错误fatal error: all goroutines are asleep - deadlock!
	var ch1 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
	close(ch1) //关闭管道
	//for range循环遍历管道的值  ,注意：管道没有key
	for v := range ch1 {
		fmt.Println(v)
	}
}

func TestChanFor(t *testing.T) {
	//通过for循环遍历管道的时候管道可以不关闭
	var ch2 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch2 <- i
	}
	for j := 0; j < 10; j++ {
		fmt.Println(<-ch2)
	}
	//管道里面没有数据了 也会报错 fatal error: all goroutines are asleep - deadlock!
	for j := 0; j < 10; j++ {
		fmt.Println(<-ch2)
	}
}

var wg sync.WaitGroup

// 写数据
func write(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("【写入】数据%v成功\n", i)
		//写慢 读快 写阻塞读 不影响
		time.Sleep(time.Millisecond * 100)
	}
	//for range遍历 所以需要关闭管道
	close(ch)
	wg.Done()
}

func read(ch chan int) {
	for v := range ch {
		fmt.Printf("【读取】数据%v成功\n", v)
		time.Sleep(time.Millisecond * 10)
	}
	wg.Done()
}

func TestChanSyncReadWrite1(t *testing.T) {
	var ch = make(chan int, 10)
	wg.Add(1)
	go write(ch)
	wg.Add(1)
	go read(ch)

	wg.Wait()
	fmt.Println("退出...")
}

// for循环未关闭的chan也会阻塞panic，但是for循环关闭的chan没问题输出默认值
func TestChanSyncReadWrite2(t *testing.T) {
	// 1、创建channel
	var ch1 = make(chan int, 3)
	wg.Add(1)
	go func() {
		for i := 1; i <= 30; i++ {
			num := <-ch1
			fmt.Println(num)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(time.Second)
			ch1 <- i
		}
		close(ch1)
		wg.Done()
	}()

	wg.Wait()
}

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
