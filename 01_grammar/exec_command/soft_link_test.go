package exec_command

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestCopySoftLinkCommand(t *testing.T) {
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
