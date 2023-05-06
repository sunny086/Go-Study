package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := os.Open("D:\\readAll.rar")
	defer file.Close()
	content, _ := ioutil.ReadAll(file)
	h := sha1.New()
	h.Write(content)
	toString := hex.EncodeToString(h.Sum(nil))
	fmt.Println(toString)
}
