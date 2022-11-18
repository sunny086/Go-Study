package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
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

// TestGetAllFileTime 获取指定目录下所有文件的修改时间
func TestGetAllFileTime(t *testing.T) {
	//递归获取目录下的所有文件
	var files []string
	files, _ = GetAllFilePath("C:\\Users\\BigStar\\Desktop\\网藤\\firewall")

	fmt.Println("目录下的所有文件如下")
	for i := 0; i < len(files); i++ {
		fmt.Println("文件名：", files[i])

		// 获取文件原来的访问时间，修改时间
		finfo, _ := os.Stat(files[i])

		// linux环境下代码如下
		//linuxFileAttr := finfo.Sys().(*syscall.Stat_t)
		//fmt.Println("文件创建时间", SecondToTime(linuxFileAttr.Ctim.Sec))
		//fmt.Println("最后访问时间", SecondToTime(linuxFileAttr.Atim.Sec))
		//fmt.Println("最后修改时间", SecondToTime(linuxFileAttr.Mtim.Sec))

		// windows下代码如下
		winFileAttr := finfo.Sys().(*syscall.Win32FileAttributeData)
		fmt.Println("文件创建时间：", SecondToTime(winFileAttr.CreationTime.Nanoseconds()/1e9))
		fmt.Println("最后访问时间：", SecondToTime(winFileAttr.LastAccessTime.Nanoseconds()/1e9))
		fmt.Println("最后修改时间：", SecondToTime(winFileAttr.LastWriteTime.Nanoseconds()/1e9))
	}
	now := time.Now()
	d, _ := time.ParseDuration("-24h")
	fmt.Printf("当前时间：%v\n", now)
	fmt.Printf("当前时间戳：%v\n", now.Unix())
	fmt.Println(now.Add(d).Format("2006-01-02 15:04:05"))
}

// 把秒级的时间戳转为time格式
func SecondToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// GetAllFilePath 递归获取指定目录下的所有文件的全路径
func GetAllFilePath(pathname string) ([]string, error) {
	var result []string

	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Printf("读取文件目录失败，pathname=%v, err=%v \n", pathname, err)
		return result, err
	}

	// 所有文件/文件夹
	for _, fi := range fis {
		fullname := pathname + "/" + fi.Name()
		// 是文件夹则递归进入获取;是文件，则压入数组
		if fi.IsDir() {
			temp, err := GetAllFilePath(fullname)
			if err != nil {
				fmt.Printf("读取文件目录失败,fullname=%v, err=%v", fullname, err)
				return result, err
			}
			result = append(result, temp...)
		} else {
			result = append(result, fullname)
		}
	}

	return result, nil
}
