package net

import (
	"fmt"
	"net"
	"testing"
)

func TestIp(t *testing.T) {
	fmt.Println(net.ParseIP("192.0.2.1"))
	fmt.Println(net.ParseIP("2001:db8::68"))
	fmt.Println(net.ParseIP("192.0.2"))
}
