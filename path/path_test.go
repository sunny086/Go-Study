package path

import (
	"fmt"
	"path"
	"path/filepath"
	"testing"
)

func TestPathJoin(t *testing.T) {
	// 路径拼接
	fmt.Println(path.Join("//a", "//b", "/c//"))
	fmt.Println(path.Join("a", "b", "c", "d"))
	fmt.Println(path.Join("a", "b", "c", "d", "e"))
}

func TestFilePath(t *testing.T) {
	// 路径操作
	fmt.Println(path.IsAbs("./abs.txt")) // false：判断是否是绝对路径
	// 获取绝对路径
	absPath, err := filepath.Abs("./abs.txt")
	if err != nil {
		fmt.Println("get abs path err: ", err)
	}
	fmt.Println("absPath: ", absPath)
}
