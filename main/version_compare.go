package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	updateVersionStrList := strings.Split("1.0.14", ".")
	currentVersionStrList := strings.Split("1.0.130", ".")
	updateVersion := updateVersionStrList[len(updateVersionStrList)-1]
	currentVersion := currentVersionStrList[len(currentVersionStrList)-1]
	//转数字
	atoi, _ := strconv.Atoi(updateVersion)
	atoi1, _ := strconv.Atoi(currentVersion)
	fmt.Println(atoi >= atoi1)

}
