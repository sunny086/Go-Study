package recover

import (
	"fmt"
	"testing"
)

// recover 只能在 defer 函数中调用才能捕获 panic
// recover 只有在 defer 函数中直接调用且有panic才会生效，否则会返回 nil
// recover 只能捕获当前 goroutine 的 panic，不能跨 goroutine，每个 goroutine 都有自己的 panic 状态
// recover 必须在 defer 函数中调用才能捕获 panic，否则 panic 会一直往上传递至goroutine的顶层，导致goroutine崩溃

// TestRecover1 先产生panic 然后再执行的defer recover 不能捕获panic
func TestRecover1(t *testing.T) {
	fmt.Println("start")
	panic1()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from panic: ", err)
		}
	}()
	fmt.Println("end")
}

// TestRecover2 先执行defer recover 然后再产生panic defer recover 可以捕获panic
func TestRecover2(t *testing.T) {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from panic: ", err)
		}
	}()
	panic1()
	fmt.Println("end")
}

func TestRecover3(t *testing.T) {
	err := recoverReturn()
	fmt.Println("recoverReturn: ", err)
}

func recoverReturn() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from panic: ", err)

		}
	}()
	panic2()
	return nil
}

// 显式触发 panic
func panic1() {
	panic("panic1")
}

// 数组越界
func panic2() {
	arr := []int{1, 2, 3}
	_ = arr[5] // 访问数组超出范围的索引，会导致 panic
}

// 尝试对空指针进行解引用，会导致 panic
func panic3() {
	var ptr *int
	_ = *ptr
}

// 除数为零，会导致 panic
//func panic4() {
//	_ = 10 / 0
//}

// 将 int 类型断言为 string 类型，会导致 panic
func panic5() {
	var val interface{} = 42
	_ = val.(string)
}

// 尝试关闭已经关闭的通道，会导致 panic
func panic6() {
	ch := make(chan int)
	close(ch)
	close(ch)
}

// 对未初始化的 map 进行读操作，会导致 panic
func panic9() {
	var m map[string]int
	_ = m["key"]
}
