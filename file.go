package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {

}

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
