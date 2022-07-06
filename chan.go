package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int, 10)
	//ch = write(ch)
	//close(ch)
	//for i := 0; i < 6; i++ {
	//	fmt.Println(<-ch)
	//}
	wg.Add(1)
	go write(ch)
	wg.Add(1)
	go print(ch)
	wg.Add(1)
	go write(ch)
	wg.Wait()
}

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- 1
	}

	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		ch <- i
	}
	wg.Done()

}

func print(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
	wg.Done()
}
