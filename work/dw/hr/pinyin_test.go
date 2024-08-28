package hr

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"strings"
	"testing"
)

func TestPY(t *testing.T) {
	p := pinyin.NewArgs()
	str := strings.Join(pinyin.LazyPinyin("大猩猩", p), "")
	fmt.Println(str)
}
