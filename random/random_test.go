package random

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	mathRand "math/rand"
	"testing"
	"time"
)

// math/rand伪随机生成的数字是确定的，不论在什么机器、什么时间，只要执行的随机代码一样，那么生成的随机数就一样。
// 为了尽量随机性，那么我们可以每次使用不同的seed来启动程序，就可以保证每次启动都产生新的随机数
// golang支持两种随机数生成方式：
// math/rand          // 伪随机
// crypto/rand        // 真随机
func TestRandomStr1(t *testing.T) {
	fmt.Println(mathRand.Intn(100))
	//设置随机数种子，由于种子数值，每次启动都不一样
	//所以每次随机数都是随机的
	mathRand.Seed(time.Now().UnixNano())
	//随机生成100以内的正整数
	fmt.Println(mathRand.Intn(100))
}

// crypto/rand 就是从这个地方读“真随机”数字返回，但性能比较慢。比上面慢10倍以上
func TestRandomStr2(t *testing.T) {
	var n int32
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	fmt.Println(n)
}

// 通用随机数算法
func TestRandomStr3(t *testing.T) {
	t.Log(randomString(10, 0))
	t.Log(randomString(10, 1))
	t.Log(randomString(10, 2))
	t.Log(randomString(10, 3))
}

/*
  - size 随机码的位数
  - kind 0    // 纯数字
    1    // 小写字母
    2    // 大写字母
    3    // 数字、大小写字母
*/
func randomString(size int, kind int) string {
	ikind, kinds, rsbytes := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	mathRand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = mathRand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		rsbytes[i] = uint8(base + mathRand.Intn(scope))
	}
	return string(rsbytes)
}
