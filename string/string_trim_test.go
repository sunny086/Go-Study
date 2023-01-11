package string

import (
	"strings"
	"testing"
)

// 去掉字符串s中首部以及尾部与字符串cutset中每个相匹配的字符
func TestTrim1(t *testing.T) {
	str := "    111a2a3 a4a5a6  a7a8a  91111       "
	trim := strings.Trim(str, " ")
	t.Log(trim) //111a2a3 a4a5a6  a7a8a  91111
}

func TestTrim2(t *testing.T) {
	str := "(0,'10.4.14.102','10.130.130.130',1,'UDP','fins',56,0)"
	trim := strings.Trim(str, "(")
	t.Log(trim)
}

func TestTrimSpace(t *testing.T) {
	str := "    111a2a3 a4a5a6  a7a8a  91111       "
	space := strings.TrimSpace(str)
	// 和trim相同
	t.Log(space) //111a2a3 a4a5a6  a7a8a  91111
}

func TestTrimLeft(t *testing.T) {
	str := "AAABBBCCCAAABBBCCCDDDzzz"
	left := strings.TrimLeft(str, "A")
	t.Log(left)
	//AAABBBCCCAAABBBCCCDDDAAA
	//   BBBCCCAAABBBCCCDDDAAA
	right := strings.TrimRight(str, "z")
	t.Log(right)
	//AAABBBCCCAAABBBCCCDDDAAA
	//AAABBBCCCAAABBBCCCDDD
}

// If s doesn't start with prefix, s is returned unchanged.
func TestTrimPrefix(t *testing.T) {
	str := "ABCDEFG"
	prefix := strings.TrimPrefix(str, "A")
	t.Log(prefix) //BCDEFG
}

// If s doesn't end with suffix, s is returned unchanged.
func TestTrimSuffix(t *testing.T) {
	str := "ABCDEFGABA"
	suffix := strings.TrimSuffix(str, "A")
	t.Log(suffix) //ABCDEFGAB
}
