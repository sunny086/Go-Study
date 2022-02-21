package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中
	str1 := make([]byte, 0, 100)
	str1 = strconv.AppendInt(str1, 4567, 10)
	str1 = strconv.AppendBool(str1, false)
	str1 = strconv.AppendQuote(str1, "abcdefg")
	str1 = strconv.AppendQuoteRune(str1, '单')
	fmt.Println(string(str1)) // 4567false"abcdefg"'单'

	// Format 系列函数把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e) // false 123.23 1234 12345 1023

	// Parse 系列函数把字符串转换为其他类型
	f, _ := strconv.ParseBool("false")
	g, _ := strconv.ParseFloat("123.23", 64)
	h, _ := strconv.ParseInt("1234", 10, 64)
	i, _ := strconv.ParseUint("12345", 10, 64)
	j, _ := strconv.Atoi("1023")
	fmt.Println(f, g, h, j, i, j) // false 123.23 1234 1023 12345 1023
}
