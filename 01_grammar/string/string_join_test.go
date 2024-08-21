package string

import (
	"fmt"
	"strings"
	"testing"
)

// func Join(elems []string, sep string) string
// string sep is placed between elements in the resulting string.
func TestStringJoin(t *testing.T) {
	s := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	join := strings.Join(s, ",")
	fmt.Println(join) //1,2,3,4,5,6,7,8,9
}
