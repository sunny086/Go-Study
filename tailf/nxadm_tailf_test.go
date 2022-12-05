package tailf

import (
	"fmt"
	"github.com/nxadm/tail"

	"os"
	"testing"
)

// 监听文本新增内容
func TestNxadmTailf(t *testing.T) {
	//	"github.com/hpcloud/tail"
	//	"github.com/nxadm/tail"
	config := tail.Config{Follow: true, ReOpen: true, Location: &tail.SeekInfo{Whence: os.SEEK_END}}
	file, err := tail.TailFile("../tailf/tailf.txt", config)
	if err != nil {
		fmt.Println(err)
	}
	for line := range file.Lines {
		fmt.Println(line.Text)
	}
}
