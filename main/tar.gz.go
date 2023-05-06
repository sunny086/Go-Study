package main

import (
	"archive/tar"
	"fmt"
	"golift.io/xtractr"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//TarTest()
	x := &xtractr.XFile{
		FilePath:  "D:\\ZZZ\\go-usb.tar.gz",
		OutputDir: "D:\\ZZZ\\", // do not forget this.
	}

	// size is how many bytes were written.
	// files may be nil, but will contain any files written (even with an error).
	size, files, _, err := xtractr.ExtractFile(x)
	if err != nil || files == nil {
		log.Fatal(size, files, err)
	}

	log.Println("Bytes written:", size, "Files Extracted:\n -", strings.Join(files, "\n -"))

}

func TarTest() {
	// 解压需要使用tar.NewReader方法, 这个方法接收一个io.Reader对象
	// 那边怎么从源文件得到io.Reader对象呢？
	// 这边通过os.Open打开文件,会得到一个os.File对象，
	// 因为他实现了io.Reader的Read方法，所有可以直接传递给tar.NewReader
	file, err := os.Open("D:/ZZZ/go-usb.tar.gz")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// tar对象读取文件内容, 遍历输出文件内容
	tr := tar.NewReader(file)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s文件内容:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}
