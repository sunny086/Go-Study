package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"time"
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

func TestWriteFileByBuff(t *testing.T) {
	err := ioutil.WriteFile("D:\\test.txt", []byte(""), 0666)
	dstFile, err := os.OpenFile("D:\\test.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	bufWriter := bufio.NewWriter(dstFile)
	st := time.Now()
	defer func() {
		//flush操作
		bufWriter.Flush()
		dstFile.Close()
		fmt.Println("文件写入耗时：", time.Now().Sub(st).Seconds(), "s")
	}()
	for i := 0; i < 100; i++ {
		bufWriter.WriteString(strconv.Itoa(i) + "\n")
	}
}
