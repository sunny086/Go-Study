package _chan

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

// chan同步读写
func TestChanSyncReadWrite1(t *testing.T) {
	var ch = make(chan int, 10)
	wg.Add(1)
	//write
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Printf("【写入】数据%v成功\n", i)
			//写慢 读快 写阻塞读 不影响
			time.Sleep(time.Millisecond * 100)
		}
		//for range遍历 所以需要关闭管道
		close(ch)
		wg.Done()
	}()

	wg.Add(1)
	//read
	go func() {
		for v := range ch {
			fmt.Printf("【读取】数据%v成功\n", v)
			time.Sleep(time.Millisecond * 10)
		}
		wg.Done()
	}()

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
