package file

import (
	"fmt"
	"os"
	"testing"
)

// ---------------------------------------------------------------------------------------------------------------------
var txtRelPath = "./test.txt"
var dirRelPath = "./test"
var multiDirRelPath = "./dd/rr"

// TestInMkdirRelativePath 在文件的相对路径创建文件、文件夹、多级文件夹 执行结束后可执行 TestRemoveFromRelativePath 测试删除
func TestInMkdirRelativePath(t *testing.T) {

	//当前go文件的相对路径-创建txt
	_, err := os.Create(txtRelPath)
	if err != nil {
		fmt.Println("create err: ", err)
	}
	// 当前go文件的相对路径-创建目录
	err = os.Mkdir(dirRelPath, os.ModePerm)
	if err != nil {
		fmt.Println("mkdir err: ", err)
	}
	// 当前go文件的相对路径-创建多级目录
	err = os.MkdirAll(multiDirRelPath, os.ModePerm)
	if err != nil {
		fmt.Println("mkdirAll err: ", err)
	}
}

// TestRemoveFromRelativePath 删除文件、文件夹、多级文件夹
func TestRemoveFromRelativePath(t *testing.T) {
	// 删除文件
	err := os.Remove(txtRelPath)
	if err != nil {
		fmt.Println("remove err: ", err)
	}
	// 当前文件为根目录-删除目录
	err = os.Remove(dirRelPath)
	if err != nil {
		fmt.Println("remove err: ", err)
	}
	// 当前文件为根目录-删除多级目录
	err = os.RemoveAll("./dd")
	if err != nil {
		fmt.Println("removeAll err: ", err)
	}
}

// ---------------------------------------------------------------------------------------------------------------------
var txtRootPath = "../test.txt"
var dirRootPath = "../file_test_mkdir/t1"

// TestMkdirInRootPath 当前项目的根路径位置
func TestMkdirInRootPath(t *testing.T) {
	//创建txt
	_, err := os.Create(txtRootPath)
	if err != nil {
		fmt.Println("create err: ", err)
	}

	// 项目根路径创建目录
	err = os.MkdirAll(dirRootPath, os.ModePerm)
	if err != nil {
		fmt.Println("mkdir err: ", err)
	}
}

// TestRemoveFromRootPath 删除根路径开始的文件文件夹
func TestRemoveFromRootPath(t *testing.T) {
	// 删除文件
	err := os.Remove(txtRootPath)
	if err != nil {
		fmt.Println("remove err: ", err)
	}
	// 项目根路径-删除多级目录
	err = os.RemoveAll("../file_test_mkdir")
	if err != nil {
		fmt.Println("removeAll err: ", err)
	}
}

// ---------------------------------------------------------------------------------------------------------------------
