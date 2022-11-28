package file

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"time"
)

// TestWriteFileByBuff 缓冲流写文件 最后要flush 如果不flush 会导致文件内容不全
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
		//flush操作 defer是最后 接下来的业务如果涉及文件读取 那读到的就是空文件
		bufWriter.Flush()
		dstFile.Close()
		fmt.Println("文件写入耗时：", time.Now().Sub(st).Seconds(), "s")
	}()
	for i := 0; i < 100; i++ {
		bufWriter.WriteString(strconv.Itoa(i) + "\n")
	}
}

// TestWriteFileByString
func TestWriteFileByString(t *testing.T) {
	filePath := "./write_string.txt"
	_, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer f.Close()

	// 写入文件内容 按字符串写 WriteString()：
	n, err := f.WriteString("xujinshan")
	if err != nil {
		fmt.Println("write err: ", err)
		return
	}
	fmt.Println("write number = ", n)
}

// TestWriteFileByBytes
func TestWriteFileByBytes(t *testing.T) {
	filePath := "./write_byte.txt"
	_, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer f.Close()
	// 写入文件内容 写入字节 Write()：
	n, err := f.Write([]byte("123hello"))
	if err != nil {
		fmt.Println("write err: ", err)
		return
	}
	fmt.Println("write number = ", n)
}
