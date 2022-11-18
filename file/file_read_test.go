package file

import (
	"fmt"
	"io"
	"os"
	"testing"
)

// TestReadFileBySpecifiedByte 按字节读取文件
func TestReadFileBySpecifiedByte(t *testing.T) {
	f, err := os.OpenFile("D:\\test.txt", os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer f.Close()
	readByte := make([]byte, 128) // 指定要读取的长度
	for {
		n, err := f.Read(readByte)
		// 将数据读取如切片，返回值 n 是实际读取到的字节数
		if err != nil && err != io.EOF {
			// 如果读到了文件末尾：EOF 即 end of file
			fmt.Println("read file : ", err)
			break
		}
		fmt.Println("read:\n", string(readByte[:n]))
		if n < 128 {
			fmt.Println("read end")
			break
		}
	}
}

// TestReadFileAtOneTime 一次性读取文件
func TestReadFileAtOneTime(t *testing.T) {
	bytes, err := os.ReadFile("D:\\test.txt")
	if err != nil {
		fmt.Println("read file err: ", err)
		return
	}
	fmt.Println(string(bytes))
}
