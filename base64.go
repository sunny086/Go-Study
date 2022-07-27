package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	encode := Base64Encode("admin:123456")
	decode := Base64Decode(encode) //YWRtaW46MTIzNDU2
	fmt.Println(decode)
}

func Base64Encode(encodeStr string) string {
	res := base64.StdEncoding.EncodeToString([]byte(encodeStr))
	fmt.Println(res)
	return res
}

func Base64Decode(decodeStr string) string {
	s, _ := base64.StdEncoding.DecodeString(decodeStr)
	fmt.Printf("base64解码结果为：%s", string(s))
	return string(s)
}
