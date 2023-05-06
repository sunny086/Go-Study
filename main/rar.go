package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"golift.io/xtractr"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	//直接md5 go-usb目录d41d8cd98f00b204e9800998ecf8427e
	//压缩123321命名go1.rar	28a81cec2f059db505562beb047a51cb
	//压缩123321命名go1.rar 手动改名字go2.rar	28a81cec2f059db505562beb047a51cb
	//压缩123321命名go2.rar	2f84b059b34670c11d7376a6813a74a8
	//压缩123321命名go2.rar 手动改名字go1.rar	2f84b059b34670c11d7376a6813a74a8
	//md5Str := md5_test("D:/ZZZ/go-usb")
	md5Str := md5_test("D:/usb-upgrade/upgrade2/go-usb2.rar")
	//md5Str := md5_test("D:/ZZZ/go-usb/go2.rar")

	fmt.Println(md5Str)
}

func SHA1File(path string) string {
	buf := make([]byte, 1024)
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer f.Close()
	r := bufio.NewReader(f)
	r.Read(buf)
	h := sha1.New()
	h.Write(buf)
	toString := hex.EncodeToString(h.Sum(nil))
	return toString
}

func Md5File(path string) string {
	buf := make([]byte, 1024)
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer f.Close()
	r := bufio.NewReader(f)
	r.Read(buf)
	h := md5.New()
	h.Write(buf)
	toString := hex.EncodeToString(h.Sum(nil))
	return toString
}

func md5_test(path string) (md5Str string) {
	f, _ := os.Open(path)
	content, _ := ioutil.ReadAll(f)
	h := md5.New()
	h.Write(content)
	md5Str = hex.EncodeToString(h.Sum(nil))
	return
}

func RarExtractByPassword() {
	x := &xtractr.XFile{
		FilePath:  "D:/usb.rar",
		OutputDir: "D:/ZZZ", // do not forget this.
		Password:  "123",
	}

	// size is how many bytes were written.
	// files may be nil, but will contain any files written (even with an error).
	size, files, _, err := xtractr.ExtractFile(x)
	if err != nil || files == nil {
		log.Fatal(size, files, err)
	}

	log.Println("Bytes written:", size, "Files Extracted:\n -", strings.Join(files, "\n -"))
}
