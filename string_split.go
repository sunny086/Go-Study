package main

import (
	"fmt"
	"strings"
)

//golang字符串操作
func main() {
	s := "1a2a3a4a5a6a7a8a9"
	//func SplitN(s, sep string, n int) []string { return genSplit(s, sep, 0, n)
	//按照sep分割n份，slices s into substrings separated by sep
	splitN := strings.SplitN(s, "a", 4)
	fmt.Println(splitN)      //[1 2 3 4a5a6a7a8a9]
	fmt.Println(len(splitN)) //4

	//func SplitAfterN(s, sep string, n int) []string
	//子串和sep结合后的切片，slices s into substrings after each instance of sep
	splitAfterN := strings.SplitAfterN(s, "a", 4)
	fmt.Println(splitAfterN)      //[1a 2a 3a 4a5a6a7a8a9]
	fmt.Println(len(splitAfterN)) //4

}
