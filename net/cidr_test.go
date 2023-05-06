package net

import (
	"fmt"
	"net"
	"net/netip"
	"testing"
)

// 直接使用net.ParseCIDR来解析cidr字符串，返回的是IP和IPNet;
// 但是发现parse解析出来的Ip是字符串里面的，但是IPNet.IP如果不是32位的bitSize那就是是0结尾的；
func TestParseCIDR1(t *testing.T) {
	cidr, ipNet, _ := net.ParseCIDR("10.25.4.128/23")
	t.Log(cidr.Equal(ipNet.IP)) //false
	t.Log(ipNet.IP.String())    //10.25.10.0--------->只要不是32，解析出来的都是0结尾的
	t.Log(cidr)                 //10.25.10.1--------->我们发现parse解析出来的Ip是字符串里面的，但是IPNet里面的IP是0结尾的
}

// netip.ParsePrefix 解析子网掩码的另外一种方式 可以看一下
// 可以看一下 net.netip_parse_prefix_test.go的测试
func TestParseCIDR2(t *testing.T) {
	prefix, _ := netip.ParsePrefix("10.25.10.1/24")
	t.Log(prefix)          //10.25.10.1/24
	t.Log(prefix.Addr())   //10.25.10.1
	t.Log(prefix.Bits())   //24
	t.Log(prefix.String()) //10.25.10.1/24
}

func TestParseCIDRMask(t *testing.T) {
	mask := net.CIDRMask(24, 32)
	fmt.Println(mask)
}
