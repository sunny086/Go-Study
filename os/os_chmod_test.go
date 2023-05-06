package os

import (
	"fmt"
	"os"
	"testing"
)

func TestChmod(t *testing.T) {
	err := os.Chmod("/opt/usb/soft_link/go-admin", 0777)
	fmt.Println(err)
}
