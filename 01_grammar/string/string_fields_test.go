package string

import (
	"fmt"
	"strings"
	"testing"
)

// splits the string s around each instance of one or more consecutive white space characters,
// as defined by unicode.IsSpace,
// returning a slice of substrings of s or an empty slice if s contains only white space.
func TestStringFields(t *testing.T) {
	s := "    111a2a3 a4a5a6  a7a8a  91111       "
	trim := strings.Fields(s)
	fmt.Println(trim)
}
