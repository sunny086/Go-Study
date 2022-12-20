package net

import (
	"net"
	"testing"
)

// 测试CIDR是否包含单个IP
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
