package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"testing"
	"time"
)

// TestGetAllFileTime 获取指定目录下所有文件的修改时间 os.TestOsStat
func TestGetAllFileTime(t *testing.T) {
	//递归获取目录下的所有文件
	var files []string
	files, _ = GetAllFilePath("C:\\Users\\BigStar\\Desktop\\网藤\\firewall")

	fmt.Println("目录下的所有文件如下")
	for i := 0; i < len(files); i++ {
		fmt.Println("文件名：", files[i])

		// 获取文件原来的访问时间，修改时间
		finfo, _ := os.Stat(files[i])

		// linux环境下代码如下
		//linuxFileAttr := finfo.Sys().(*syscall.Stat_t)
		//fmt.Println("文件创建时间", SecondToTime(linuxFileAttr.Ctim.Sec))
		//fmt.Println("最后访问时间", SecondToTime(linuxFileAttr.Atim.Sec))
		//fmt.Println("最后修改时间", SecondToTime(linuxFileAttr.Mtim.Sec))

		// windows下代码如下
		winFileAttr := finfo.Sys().(*syscall.Win32FileAttributeData)
		fmt.Println("文件创建时间：", SecondToTime(winFileAttr.CreationTime.Nanoseconds()/1e9))
		fmt.Println("最后访问时间：", SecondToTime(winFileAttr.LastAccessTime.Nanoseconds()/1e9))
		fmt.Println("最后修改时间：", SecondToTime(winFileAttr.LastWriteTime.Nanoseconds()/1e9))
	}
	now := time.Now()
	d, _ := time.ParseDuration("-24h")
	fmt.Printf("当前时间：%v\n", now)
	fmt.Printf("当前时间戳：%v\n", now.Unix())
	fmt.Println(now.Add(d).Format("2006-01-02 15:04:05"))
}

// 把秒级的时间戳转为time格式
func SecondToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// GetAllFilePath 递归获取指定目录下的所有文件的全路径
func GetAllFilePath(pathname string) ([]string, error) {
	var result []string

	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Printf("读取文件目录失败，pathname=%v, err=%v \n", pathname, err)
		return result, err
	}

	// 所有文件/文件夹
	for _, fi := range fis {
		fullname := pathname + "/" + fi.Name()
		// 是文件夹则递归进入获取;是文件，则压入数组
		if fi.IsDir() {
			temp, err := GetAllFilePath(fullname)
			if err != nil {
				fmt.Printf("读取文件目录失败,fullname=%v, err=%v", fullname, err)
				return result, err
			}
			result = append(result, temp...)
		} else {
			result = append(result, fullname)
		}
	}

	return result, nil
}
