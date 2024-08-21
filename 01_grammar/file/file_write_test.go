package file

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

// TestWriteFileByBuff 缓冲流写文件 最后要flush 如果不flush 会导致文件内容不全
func TestWriteFileByBuff(t *testing.T) {
	filepath := "./write_buff.txt"
	dstFile, _ := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	bufWriter := bufio.NewWriter(dstFile)
	st := time.Now()
	defer func() {
		//flush操作 defer是最后 接下来的业务如果涉及文件读取 那读到的就是空文件
		bufWriter.Flush()
		dstFile.Close()
		fmt.Println("文件写入耗时：", time.Now().Sub(st).Seconds(), "s")
	}()
	for i := 0; i < 10; i++ {
		bufWriter.WriteString(strconv.Itoa(i) + "\n")
	}
}

// TestWriteFileByString os.O_APPEND 会在文件末尾追加内容
func TestWriteFileByString(t *testing.T) {
	filePath := "./write_string.txt"
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer f.Close()

	// 写入文件内容 按字符串写 WriteString()：
	n, err := f.WriteString("xujinshan\n")
	if err != nil {
		fmt.Println("write err: ", err)
		return
	}
	fmt.Println("write number = ", n)
}

// TestWriteFileByBytes os.O_TRUNC 会清空文件内容 os.O_CREATE 会创建文件
func TestWriteFileByBytes(t *testing.T) {
	filePath := "./write_byte.txt"
	f, err := os.OpenFile(filePath, os.O_TRUNC|os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer f.Close()
	// 写入文件内容 写入字节 Write()：
	n, err := f.Write([]byte("123hello\n"))
	if err != nil {
		fmt.Println("write err: ", err)
		return
	}
	fmt.Println("write number = ", n)
}
