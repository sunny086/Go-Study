package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Chmod("/opt/usb/soft_link/go-admin", 0777)
	fmt.Println(err)
}
