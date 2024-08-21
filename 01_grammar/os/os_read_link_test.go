package os

import (
	"fmt"
	"os"
	"testing"
)

// 获取真实得软链接的目标地址
func TestReadRealDestinationOfSoftLink(t *testing.T) {
	readlink, err := os.Readlink("/opt/usb/soft_link")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(readlink) //    会得到这个真实目标地址 /opt/usb/version/origin
}
