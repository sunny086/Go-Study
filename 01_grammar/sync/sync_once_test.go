package sync

import (
	"fmt"
	"sync"
	"testing"
)

// sync.Once用来保证函数只执行一次
func TestSyncOnce(t *testing.T) {
	var once sync.Once
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			once.Do(func() {
				fmt.Println("once")
			})
		}
		wg.Done()
	}()
	wg.Wait()
}
