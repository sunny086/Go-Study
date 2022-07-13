package main

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
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
