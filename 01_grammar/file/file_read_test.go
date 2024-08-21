package file

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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

// TestReadJpgFile 读取jpg文件
func TestReadJpgFile(t *testing.T) {
	file, err := os.Open("D:\\1.jpg")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	result := judgeType(file)
	fmt.Println("判断结果: ", result)
}

func judgeType(file *os.File) bool {
	buf := make([]byte, 20)
	//读取文件的前20个字节
	n, _ := file.Read(buf)
	//把byte转换为十六进制
	fileCode := bytesToHexString(buf[:n])
	picMap := make(map[string]string)
	picMap["ffd8ffe0"] = "jpg"
	picMap["ffd8ffe1"] = "jpg"
	picMap["ffd8ffe8"] = "jpg"
	picMap["89504e47"] = "png"
	for k, _ := range picMap {
		if strings.HasPrefix(fileCode, k) {
			return true
		}
	}
	return false
}

// byte转换16进制
func bytesToHexString(src []byte) string {
	res := bytes.Buffer{}
	if src == nil || len(src) <= 0 {
		return ""
	}
	temp := make([]byte, 0)
	i, length := 100, len(src)
	if length < i {
		i = length
	}
	for j := 0; j < i; j++ {
		sub := src[j] & 0xFF
		hv := hex.EncodeToString(append(temp, sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0), 10))
		}
		res.WriteString(hv)
	}
	return res.String()
}

func TestCheckFileExist(t *testing.T) {
	src := "D:\\test.txt"
	suffix := "txt"
	//判断指定目录下的文件是否包含指定的后缀名文件
	fileInfo, err := os.Stat(src)
	if err != nil {
		fmt.Println("文件不存在")
		return
	}
	if fileInfo.IsDir() {
		//判断是否是目录
		fmt.Println("指定的文件不是文件")
	}
	if !strings.HasSuffix(src, suffix) {
		fmt.Println("文件后缀名不是指定的后缀名")
	} else {
		fmt.Println("文件后缀名是指定的后缀名")
	}
}
