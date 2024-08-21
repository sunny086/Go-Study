package net

import (
	"net"
	"net/netip"
	"testing"
)

// 使用net解析crdr返回的IpNet测试CIDR是否包含单个IP，入参是IP
func TestCidrContainsSingleIP(t *testing.T) {
	ip := net.ParseIP("10.25.10.1")
	t.Log(ip)
	_, ipNet, err := net.ParseCIDR("10.25.10.1/24")
	if err != nil {
		t.Error(err)
	}
	t.Log(ipNet)              //10.25.10.0/24-------------->ipNet.IP和ipNet.Mask就是这个ip的地址和掩码
	t.Log(ipNet.IP)           //10.25.10.0
	t.Log(ipNet.Mask)         //ffffff00------->255.255.255.0
	t.Log(ipNet.Contains(ip)) //true 子网掩码里面包含这个ip地址
}

// 使用netip解析返回的prefix测试CIDR是否包含单个IP，入参是Addr
func TestCidrContainsSingleIP2(t *testing.T) {
	// 解析cidr
	prefix, err := netip.ParsePrefix("10.25.10.1/24")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(prefix)
	// 解析ip
	addr, err := netip.ParseAddr("10.25.10.1")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(addr)
	contains := prefix.Contains(addr)
	t.Log(contains) //true
}

func TestCidrContainsCidr(t *testing.T) {

	//prefix的addr相同，但是bitSize不同，比较bitSize大小，舍弃小的cidr

	//source := []string{
	//	"10.25.10.1/24",
	//	"10.25.10.1/28",
	//}
	bigIp, _ := netip.ParsePrefix("10.25.10.1/24")
	smallIp, _ := netip.ParsePrefix("10.25.10.128/25")
	t.Log(smallIp.Addr())
	contains := bigIp.Contains(smallIp.Addr())
	t.Log(contains)

	//cidr, ipNet, err := net.ParseCIDR("10.25.2.128/23")
	//cidr, ipNet, err := net.ParseCIDR("10.25.2.128/24")

}

// ContainsCIDR 子网a 是否包含 子网b
// b 是 a 的子集
// return true - b是a的子网; false b 不是 a 的子网
func ContainsCIDR(a, b *net.IPNet) bool {
	ones1, _ := a.Mask.Size()
	ones2, _ := b.Mask.Size()
	return ones1 <= ones2 && a.Contains(b.IP)
}
