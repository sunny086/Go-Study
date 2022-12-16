package net

import (
	"net"
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

func TestParseCIDR(t *testing.T) {
	ip, ipNet, err := net.ParseCIDR("10.25.10.1/24")
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(ip.Equal(ipNet.IP))
	t.Log(ipNet.IP.String())
}
