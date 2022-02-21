package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "    111a2a3 a4a5a6  a7a8a  91111       "
	trim := strings.Fields(s)
	fmt.Println(trim)
}
