package err

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	var s []int
	s = append(s, 1)
	fmt.Println(s)

}

func Test02(t *testing.T) {
	var m map[string]string
	//m = make(map[string]string)
	m["a"] = "a"
	fmt.Println(m)
}
