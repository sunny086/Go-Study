package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- 1
	}
	close(ch)
	for i := 0; i < 6; i++ {
		fmt.Println(<-ch)
	}

}
