package main

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GetAllFiles(dirPth string) (files []string, err error) {
	fis, err := ioutil.ReadDir(filepath.Clean(filepath.ToSlash(dirPth)))
	if err != nil {
		return nil, err
	}

	for _, f := range fis {
		_path := filepath.Join(dirPth, f.Name())

		if f.IsDir() {
			fs, _ := GetAllFiles(_path)
			files = append(files, fs...)
			continue
		}

		// 指定格式
		switch filepath.Ext(f.Name()) {
		case ".png", ".jpg":
			files = append(files, _path)
		}
	}
	return files, nil
}

func CheckFile(src string, suffix string) error {
	//判断指定目录下的文件是否包含指定的后缀名文件
	fileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if fileInfo.IsDir() {
		//判断是否是目录
		return errors.New("指定的文件不是文件")
	}
	if !strings.HasSuffix(src, suffix) {
		return errors.New("指定的文件不是" + suffix)
	}
	return nil
}

//根据文件头获取文件类型
func GetFileType(file []byte) string {
	if len(file) < 4 {
		return ""
	}
	if file[0] == 0xFF && file[1] == 0xD8 && file[2] == 0xFF {
		return "jpg"
	}
	if file[0] == 0x89 && file[1] == 0x50 && file[2] == 0x4E && file[3] == 0x47 {
		return "png"
	}
	if file[0] == 0x47 && file[1] == 0x49 && file[2] == 0x46 {
		return "gif"
	}
	if file[0] == 0x42 && file[1] == 0x4D {
		return "bmp"
	}
	if file[0] == 0x49 && file[1] == 0x49 && file[2] == 0x2A && file[3] == 0x00 {
		return "tiff"
	}
	if file[0] == 0x4D && file[1] == 0x4D && file[2] == 0x00 && file[3] == 0x2A {
		return "tiff"
	}
	return ""
}

var picMap map[string]string

func init() {
	picMap = make(map[string]string)
	picMap["ffd8ffe0"] = "jpg"
	picMap["ffd8ffe1"] = "jpg"
	picMap["ffd8ffe8"] = "jpg"
	picMap["89504e47"] = "png"
}

func main() {
	//readJpgFile()
	//createOpenWriteFile()
	//fileSeek()
	//osStat()
	//pathMkdir()
	//fileRemove()

}

func fileRemove() {
	err := os.RemoveAll("./dd")
	if err != nil {
		fmt.Println("remove err:", err)
		return
	}
}

func pathMkdir() {
	// 路径操作
	fmt.Println(filepath.IsAbs("./test.txt")) // false：判断是否是绝对路径
	fmt.Println(filepath.Abs("./test.txt"))   // 转换为绝对路径

	// 创建目录
	err := os.Mkdir("./test", os.ModePerm)
	if err != nil {
		fmt.Println("mkdir err: ", err)
		return
	}

	// 创建多级目录
	err = os.MkdirAll("./dd/rr", os.ModePerm)
	if err != nil {
		fmt.Println("mkdirAll err: ", err)
		return
	}
}

func osStat() {
	fileInfo, err := os.Stat("D:\\test.txt")
	if err != nil {
		fmt.Println("stat err: ", err)
		return
	}
	fmt.Printf("%T\n", fileInfo) // *os.fileStat
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Sys())
	fmt.Println(fileInfo.Name())
}

func fileSeek() {
	//修改文件的读写指针位置 Seek()，包含两个参数：
	//参数1：偏移量，为正数时向文件尾偏移，为负数时向文件头偏移
	//参数2：偏移的开始位置，包括：
	//io.SeekStart：从文件起始位置开始
	//io.SeekCurrent：从文件当前位置开始
	//io.SeekEnd：从文件末尾位置开始
	f, _ := os.OpenFile("D:\\test.txt", os.O_RDWR, 6)
	off, _ := f.Seek(5, io.SeekStart)
	fmt.Println(off) // 5
	n, _ := f.WriteAt([]byte("111"), off)
	fmt.Println(n)
	f.Close()
}

//createOpenWriteFile 创建、打开、写入文件
func createOpenWriteFile() {
	f, err := os.Create("D:\\test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f) // 打印文件指针
	f.Close()

	f, err = os.OpenFile("D:\\test.txt", os.O_APPEND|os.O_RDWR, os.ModeAppend)
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

	// 写入文件内容 按字符串写 WriteString()：
	n, err = f.WriteString("xujinshan") // 会将前5个字符替换为 hello
	if err != nil {
		fmt.Println("write err: ", err)
		return
	}
	fmt.Println("write number = ", n)
}

//readJpgFile 读取jpg文件
func readJpgFile() {
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
	n, _ := file.Read(buf)

	fileCode := bytesToHexString(buf[:n])
	for k, _ := range picMap {
		if strings.HasPrefix(fileCode, k) {
			return true
		}
	}
	return false
}

//获取16进制
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
