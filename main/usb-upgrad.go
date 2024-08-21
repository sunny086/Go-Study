package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	err := Command("sh /root/work/usb/depoly/config/usb-server-dev.sh")
	if err != nil {
		log.Println("执行系统升级的shell命令异常：", err.Error())
	}
}

// Command 这里为了简化，我省去了stderr和其他信息
func Command(cmd string) error {
	c := exec.Command("bash", "-c", cmd)
	// 此处是windows版本
	// c := exec.Command("cmd", "/C", cmd)
	output, err := c.CombinedOutput()
	fmt.Println(string(output))
	return err
}
