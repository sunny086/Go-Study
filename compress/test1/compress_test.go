package test1

import (
	"compress/gzip"
	"io"
	"os"
	"testing"
)

// 两个文件打开后，将文件内容写入到压缩文件中，copy两次相当于合并了文件内容 并没有实现多文件压缩
func TestCompress(t *testing.T) {
	inputFile1, err := os.Open("1.txt")
	inputFile2, err := os.Open("2.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile1.Close()
	defer inputFile2.Close()

	outputFile, err := os.Create("compress.zip")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	_, err = io.Copy(gzipWriter, inputFile1)
	_, err = io.Copy(gzipWriter, inputFile2)
	if err != nil {
		panic(err)
	}
}
