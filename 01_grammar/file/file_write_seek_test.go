package file

import (
	"fmt"
	"io"
	"os"
	"testing"
)

// TestSeekCurrent TestSeekStart 效果一样
// TestSeekCurrent
func TestSeekCurrent(t *testing.T) {
	filepath := "./seek.txt"
	//修改文件的读写指针位置 Seek()，包含两个参数：
	//参数1：偏移量，为正数时向文件尾偏移，为负数时向文件头偏移
	//参数2：偏移的开始位置，包括：
	//io.SeekStart：从文件起始位置开始
	//io.SeekCurrent：从文件当前位置开始
	//io.SeekEnd：从文件末尾位置开始
	f, _ := os.OpenFile(filepath, os.O_RDWR, 6)
	off, _ := f.Seek(3, io.SeekCurrent)
	fmt.Println(off) // 5
	n, _ := f.WriteAt([]byte("88"), off)
	fmt.Println(n)
	f.Close()
}

// TestSeekStart
func TestSeekStart(t *testing.T) {
	filepath := "./seek.txt"
	//修改文件的读写指针位置 Seek()，包含两个参数：
	//参数1：偏移量，为正数时向文件尾偏移，为负数时向文件头偏移
	//参数2：偏移的开始位置，包括：
	//io.SeekStart：从文件起始位置开始
	//io.SeekCurrent：从文件当前位置开始
	//io.SeekEnd：从文件末尾位置开始
	f, _ := os.OpenFile(filepath, os.O_RDWR, 6)
	off, _ := f.Seek(3, io.SeekStart)
	fmt.Println(off) // 5
	n, _ := f.WriteAt([]byte("99"), off)
	fmt.Println(n)
	f.Close()
}

// TestSeekEnd 可以从文件末尾实现追加功能
// TestSeekEnd 从指定偏移位置写文件 一个汉字占用三个字节
func TestSeekEnd(t *testing.T) {
	filepath := "./seek.txt"
	//修改文件的读写指针位置 Seek()，包含两个参数：
	//参数1：偏移量，为正数时向文件尾偏移，为负数时向文件头偏移
	//参数2：偏移的开始位置，包括：
	//io.SeekStart：从文件起始位置开始
	//io.SeekCurrent：从文件当前位置开始
	//io.SeekEnd：从文件末尾位置开始
	f, _ := os.OpenFile(filepath, os.O_RDWR, 6)
	off, _ := f.Seek(0, io.SeekEnd)
	fmt.Println(off) // 5
	n, _ := f.WriteAt([]byte("\n99"), off)
	fmt.Println(n)
	f.Close()
}
