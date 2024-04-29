package money

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestDad(t *testing.T) {
	// 读取文件内容 test文件可以直接默认当前根目录 如果是main的话 需要指定路径 ./money/money.txt
	content, err := os.ReadFile("money.txt")
	if err != nil {
		fmt.Println("读取文件时发生错误:", err)
		return
	}
	contentStr := []byte(strings.ReplaceAll(string(content), "一", "-"))
	fmt.Println(string(contentStr))

}
