package os

import (
	"fmt"
	"os"
	"testing"
)

func TestOsStat(T *testing.T) {
	filepath := "./stat.txt"
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("stat err: ", err)
		return
	}
	//%T 输出值的类型，注意int32和int是两种不同的类型，编译器不会自动转换，需要类型转换。
	fmt.Printf("%T\n", fileInfo) // *os.fileStat
	fmt.Println("size:", fileInfo.Size())
	fmt.Println("mode:", fileInfo.Mode())
	fmt.Println("modTime:", fileInfo.ModTime())
	fmt.Println("isDir:", fileInfo.IsDir())
	fmt.Println("Sys:", fileInfo.Sys())
	fmt.Println("info:", fileInfo.Name())
}
