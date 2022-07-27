package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Grown() {
	p.Age += 1
}

func main() {

	//syncOnce()
	syncMutex()

}

func syncMutex() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	num := 0
	// 开启10个协程，每个协程都让共享数据 num + 1
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock() // 加锁，阻塞其他协程获取锁
			num += 1
			mutex.Unlock() // 解锁
			wg.Done()
		}()
	}
	wg.Wait()
	// 输出1000，如果没有加锁，则输出的数据很大可能不是1000
	fmt.Println("num = ", num)
}

//syncOnce 负责只执行一次，也即全局唯一操作
func syncOnce() {
	var once sync.Once
	var wg sync.WaitGroup

	p := &Person{
		"比尔",
		0,
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(func() {
				p.Grown()
			})
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("年龄是：", p.Age) // 只会输出 1
}
