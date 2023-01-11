package string

import (
	"strings"
	"testing"
)

// 替换指定字符 可设置替换次数
func TestReplace(t *testing.T) {
	str := "AABBBBBCCDD"
	trim1 := strings.Replace(str, "B", "!", -1)
	t.Log(trim1)
	trim2 := strings.Replace(str, "B", "!", 3)
	t.Log(trim2)
}

func TestReplace2(t *testing.T) {
	str := "AABBBBBCCDD"
	all := strings.ReplaceAll(str, "B", "!")
	t.Log(all)
}
