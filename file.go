package main

import (
	"bufio"
	"os"
)

func main() {
	//err := ioutil.WriteFile("industry.rules", []byte(""), 0666)
	//if err != nil {
	//	panic(err)
	//}
	dstFile, err := os.OpenFile("industry.rules", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	bufWriter := bufio.NewWriter(dstFile)
	bufWriter.WriteString("alert tcp any any -> any any (msg:\"test\";sid:1000001;)")
	bufWriter.Flush()
	dstFile.Close()
}
