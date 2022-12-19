package net

import (
	"fmt"
	"net"
	"net/netip"
	"testing"
)

func TestIp(t *testing.T) {
	fmt.Println(net.ParseIP("192.0.2.1"))
	fmt.Println(net.ParseIP("2001:db8::68"))
	fmt.Println(net.ParseIP("192.0.2"))
}

func TestNetIp(t *testing.T) {

}

func TestNetIPV4(t *testing.T) {
	pv4 := net.IPv4(8, 8, 8, 8)
	t.Log(pv4)
}

func TestNetIpParsePrefix(t *testing.T) {
	prefix, err := netip.ParsePrefix("10.25.110.1/3")
	if err != nil {
		fmt.Println(err)
	}
	t.Log(prefix)                 //10.25.110.1/3
	t.Log(prefix.Addr())          //10.25.110.1
	t.Log(prefix.Bits())          //3
	t.Log(prefix.Addr().BitLen()) //32
	t.Log(prefix.IsSingleIP())    //false
}
