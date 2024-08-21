package _chan

import (
	"fmt"
	"testing"
)

// 通过for循环遍历管道的时候，确定边界没有超出管道实际数据长度时，chan可以不关闭
func TestChanFor(t *testing.T) {
	var ch2 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch2 <- i
	}
	for j := 0; j < 10; j++ {
		fmt.Println(<-ch2)
	}
	////管道里面没有数据了 也会报错 fatal error: all goroutines are asleep - deadlock!
	//for j := 0; j < 10; j++ {
	//	fmt.Println(<-ch2)
	//}
}
