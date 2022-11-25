package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

// 根据文件头获取文件类型
func GetFileType(file []byte) string {
	if len(file) < 4 {
		return ""
	}
	if file[0] == 0xFF && file[1] == 0xD8 && file[2] == 0xFF {
		return "jpg"
	}
	if file[0] == 0x89 && file[1] == 0x50 && file[2] == 0x4E && file[3] == 0x47 {
		return "png"
	}
	if file[0] == 0x47 && file[1] == 0x49 && file[2] == 0x46 {
		return "gif"
	}
	if file[0] == 0x42 && file[1] == 0x4D {
		return "bmp"
	}
	if file[0] == 0x49 && file[1] == 0x49 && file[2] == 0x2A && file[3] == 0x00 {
		return "tiff"
	}
	if file[0] == 0x4D && file[1] == 0x4D && file[2] == 0x00 && file[3] == 0x2A {
		return "tiff"
	}
	return ""
}

func main() {
	//readJpgFile()
	//createOpenWriteFile()
	//fileSeek()
	//osStat()
	//pathMkdir()
	//fileRemove()
	//readFile()
	//removeExcludePath()
	//CopySoftLinkCommand()
	//GetAllFileTime()

}

func CopySoftLinkCommand() {
	err2 := exec.Command("bash", "-c", "cp -d /opt/usb/soft_link_bak11 /opt/usb/soft_link_bak111").Run()
	fmt.Println(err2)

	//下面这种 直接copy会报错被拷贝的是个目录
	/*	file, err := os.Open("/opt/usb/soft_link")
		if err != nil {
			fmt.Println(err)
		}
		create, err := os.Create("/opt/usb/soft_link_copy")
		if err != nil {
			fmt.Println(err)
		}
		written, err := io.Copy(create, file)
		if err != nil {
			//2022/08/13 16:14:39 read /opt/usb/soft_link: is a directory
			log.Fatal(err)
		}
		log.Printf("Copied %d bytes.", written)*/

}

func removeExcludePath() {
	//dir, err := ioutil.ReadDir(models.TempDir)
	//for _, d := range dir {
	//	os.RemoveAll(path.Join([]string{models.TempDir, d.Name()}...))
	//}
	dir, _ := ioutil.ReadDir("./dd")
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{"./dd", d.Name()}...))
	}
}

func fileRemove() {
	err := os.RemoveAll("./dd")
	if err != nil {
		fmt.Println("remove err:", err)
		return
	}
}

func pathMkdir() {
	// 路径操作
	fmt.Println(filepath.IsAbs("./test.txt")) // false：判断是否是绝对路径
	fmt.Println(filepath.Abs("./test.txt"))   // 转换为绝对路径

	// 创建目录
	err := os.Mkdir("./test", os.ModePerm)
	if err != nil {
		fmt.Println("mkdir err: ", err)
		return
	}

	// 创建多级目录
	err = os.MkdirAll("./dd/rr", os.ModePerm)
	if err != nil {
		fmt.Println("mkdirAll err: ", err)
		return
	}
}

func osStat() {
	fileInfo, err := os.Stat("D:\\test.txt")
	if err != nil {
		fmt.Println("stat err: ", err)
		return
	}
	fmt.Printf("%T\n", fileInfo) // *os.fileStat
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Sys())
	fmt.Println(fileInfo.Name())
}

// createOpenWriteFile 创建、打开、写入文件
func createOpenWriteFile() {
	f, err := os.Create("D:\\test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f) // 打印文件指针
	f.Close()

	f, err = os.OpenFile("D:\\test.txt", os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer f.Close()

	// 写入文件内容 写入字节 Write()：
	n, err := f.Write([]byte("123hello"))
	if err != nil {
		fmt.Println("write err: ", err)
		return
	}
	fmt.Println("write number = ", n)

	// 写入文件内容 按字符串写 WriteString()：
	n, err = f.WriteString("xujinshan") // 会将前5个字符替换为 hello
	if err != nil {
		fmt.Println("write err: ", err)
		return
	}
	fmt.Println("write number = ", n)
}
