package net

import (
	"net"
	"net/netip"
	"strconv"
	"strings"
	"testing"
)

// TestCIDRMaskSuffix2Bit 测试子网掩码255.255.255.0形式的数据转成字节位数
func TestCIDRMaskSuffix2Bit(t *testing.T) {
	cidr := SubnetMask2CIDR("255.255.254.0")
	t.Log(cidr)
}

// SubnetMask2CIDR 255.255.255.0 ---->24
func SubnetMask2CIDR(subnetMask string) int {
	//subnetMask  a.b.c.d
	split := strings.Split(subnetMask, ".")
	a, _ := strconv.Atoi(split[0])
	b, _ := strconv.Atoi(split[1])
	c, _ := strconv.Atoi(split[2])
	d, _ := strconv.Atoi(split[3])
	size, _ := net.IPv4Mask(byte(a), byte(b), byte(c), byte(d)).Size()
	return size
}

// 直接使用net.ParseCIDR来解析cidr字符串，返回的是IP和IPNet;
// 但是发现parse解析出来的Ip是字符串里面的，但是IPNet.IP如果不是32位的bitSize那就是是0结尾的；
func TestParseCIDR1(t *testing.T) {
	cidr, ipNet, _ := net.ParseCIDR("10.25.10.1/24")
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
