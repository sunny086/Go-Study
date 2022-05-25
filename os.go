package main

import (
	"fmt"
	"runtime"
)

func main() {

	//判断是否是windows系统
	fmt.Println("OS:", runtime.GOOS)

	sysType := runtime.GOOS

	if sysType == "windows" {
		fmt.Println("windows")
	}

}
