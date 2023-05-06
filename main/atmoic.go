package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var atomicWg sync.WaitGroup

func main() {
	m := make(map[string]*int64)
	var num int64 = 0
	m["1"] = &num
	atomicWg.Add(1)
	go f1(m)
	atomicWg.Add(1)
	go f2(m)
	atomicWg.Add(1)
	go f1(m)
	atomicWg.Add(1)
	go f2(m)

	delete(m, "1")

	atomicWg.Wait()
	fmt.Println(*m["1"])
}

func f1(m map[string]*int64) {
	for i := 0; i < 1000000; i++ {
		atomic.AddInt64(m["1"], 1)
	}
	atomicWg.Done()
}

func f2(m map[string]*int64) {
	for i := 0; i < 100000; i++ {
		atomic.AddInt64(m["1"], -1)
	}
	atomicWg.Done()
}
