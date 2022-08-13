package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {

	//获取操作系统类型
	//GetOsType()
	//获取真实得软链接的目标地址
	//ReadRealDestinationOfSoftLink()

}

//ReadRealDestinationOfSoftLink 获取真实得软链接的目标地址
func ReadRealDestinationOfSoftLink() {
	readlink, err := os.Readlink("/opt/usb/soft_link")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(readlink) //    会得到这个真实目标地址 /opt/usb/version/origin
}

// GetOsType 获取操作系统类型
func GetOsType() {
	//判断是否是windows系统
	fmt.Println("OS:", runtime.GOOS)
	sysType := runtime.GOOS
	if sysType == "windows" {
		fmt.Println("windows")
	}
}
