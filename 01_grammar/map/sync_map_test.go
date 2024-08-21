package _map

import (
	"fmt"
	"sync"
	"testing"
)

var num = 0
var addTest *AddTest

func init() {
	addTest = &AddTest{}
}

type AddTest struct {
	m sync.Mutex
}

func (at *AddTest) increment(wg *sync.WaitGroup) {
	//互斥锁
	at.m.Lock() //当有线程进去进行加锁
	num++
	at.m.Unlock() //出来后解锁，其他线程才可以进去
	wg.Done()
}

func (at *AddTest) decrement(wg *sync.WaitGroup) {
	//互斥锁
	at.m.Lock() //当有线程进去进行加锁
	num--
	at.m.Unlock() //出来后解锁，其他线程才可以进去
	wg.Done()
}

var w sync.WaitGroup
var aa map[int]int

func TestSyncMap(t *testing.T) {
	var bb sync.Map
	var wg sync.WaitGroup
	//aa = make(map[int]int)
	wg.Add(2)
	go func() {
		//wg.Add(1)
		for i := 0; i < 100; i++ {
			//aa[i] = i+1
			//fmt.Println("a")
			bb.Store(i, i+1)
		}
		wg.Done()
	}()

	go func() {

		for i := 0; i < 100; i++ {
			//aa[i] = i+1
			//fmt.Println("a")
			bb.Store(i, i+1)
		}
		wg.Done()
	}()
	wg.Wait()
	bb.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
