package file

import (
	"fmt"
	"os"
	"testing"
)

func TestFileExist(t *testing.T) {
	// 判断文件是否存在
	_, err := os.Stat("abs.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file.go does not exist")
		}
	} else {
		fmt.Println("file.go exists")
	}
}
