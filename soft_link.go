package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	//err := exec.Command("ln", "-s", "/opt/usb/version/41cefa78-b5a2-4d34-bb96-f6405a5eb315/", "/opt/usb/soft_link/").Run()
	//err := CommandTest("ln -s /opt/usb/version/41cefa78-b5a2-4d34-bb96-f6405a5eb315/ /opt/usb/soft_link")
	//fmt.Println(err)
	err := os.Symlink("/opt/usb/version/origin/", "/opt/usb/soft_link")
	fmt.Println(err)
}

func CommandTest(cmd string) error {
	c := exec.Command("bash", "-c", cmd)
	// 此处是windows版本
	// c := exec.Command("cmd", "/C", cmd)
	output, err := c.CombinedOutput()
	log.Println(string(output))
	return err
}
