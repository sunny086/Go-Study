package main

import (
	"fmt"
	"strings"
)

func main() {
	var argsValue = "fuc.systemSetting=1;fuc.networkTopology=0;fuc.assertManagement=1;fuc.safeAreaManagement=1;prod.usbSafeDevice=20;prod.safeAudit=2;"
	var args = strings.Split(argsValue, ";")
	//[systemSetting networkTopology assertManagement safeAreaManagement]
	var menuList = make([]string, 0)
	var menuAuthList = make([]string, 0)
	var menuNoAuthList = make([]string, 0)
	//map[assertManagement:1 networkTopology:1 safeAreaManagement:1 systemSetting:1]
	var menuMap = make(map[string]string)
	var menuAuthMap = make(map[string]string)
	var menuNoMap = make(map[string]string)
	for _, arg := range args {
		if strings.HasPrefix(arg, "fuc.") {
			//将0和1塞进不同的集合中
			if strings.HasSuffix(arg, "1") {
				menuAuthList = append(menuAuthList, arg[4:len(arg)-2])
				menuAuthMap[arg[4:len(arg)-2]] = "1"
			} else {
				menuNoAuthList = append(menuNoAuthList, arg[4:len(arg)-2])
				menuNoMap[arg[4:len(arg)-2]] = "0"
			}
			//截取fuc后面的参数 获取菜单的key
			menuList = append(menuList, arg[4:len(arg)-2])
			//转换成map
			menuMap[arg[4:len(arg)-2]] = arg[len(arg)-1:]
		}
	}
	fmt.Println(menuList)
	fmt.Println(menuMap)
	fmt.Println(menuAuthList)
	fmt.Println(menuAuthMap)
	fmt.Println(menuNoAuthList)
	fmt.Println(menuNoMap)
}
