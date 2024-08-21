package file

import (
	"fmt"
	"path/filepath"
	"testing"
)

// TestAbsPath 测试绝对路径
func TestAbsPath(t *testing.T) {
	// 路径操作
	fmt.Println(filepath.IsAbs("./abs.txt")) // false：判断是否是绝对路径
	// 获取绝对路径
	absPath, err := filepath.Abs("./abs.txt")
	if err != nil {
		fmt.Println("get abs path err: ", err)
	}
	fmt.Println("absPath: ", absPath)

}
