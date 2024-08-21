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

func TestNetIpParseAddr(t *testing.T) {
	ip := "10.25.10.1"
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		t.Error(err)
	}
	t.Log(addr)
}

func TestPrefixForm(t *testing.T) {
	prefix, err := netip.ParsePrefix("10.25.10.1/24")
	if err != nil {
		t.Error(err)
	}
	t.Log(prefix)
}

func TestIpContains(t *testing.T) {
	ip := net.ParseIP("10.25.10.1")
	t.Log(ip)
	cidr, ipNet, err := net.ParseCIDR("10.25.10.1/24")
	if err != nil {
		t.Error(err)
	}
	t.Log(cidr)                 //10.25.10.1 ip的地址
	t.Log(ipNet)                //10.25.10.0/24-------------->ipNet.IP和ipNet.Mask就是这个ip的地址和掩码
	t.Log(ipNet.IP)             //10.25.10.0
	t.Log(ipNet.Mask)           //ffffff00 255.255.255.0
	t.Log(ipNet.Contains(cidr)) //true 子网掩码里面包含这个ip地址
	t.Log(ipNet.Network())      //ip+net
	t.Log(ipNet.String())       //10.25.10.0/24
}

func TestIpRange(t *testing.T) {
	num := 1025
	//计算返回num最近的2的幂次方

	fmt.Println(num)
}

// netip.Addr 单个ip的比较
func TestAddrCompare(t *testing.T) {
	parseAddr, err := netip.ParseAddr("10.25.10.1")
	if err != nil {
		t.Error(err)
	}
	//next下一个ip地址
	next := parseAddr.Next()
	t.Log(next)
	//prev上一个ip地址
	prev := parseAddr.Prev()
	t.Log(prev)
	//相等为0，大于为1，小于为-1
	t.Log(parseAddr.Compare(next))
	t.Log(parseAddr.Compare(prev))
}

func TestCIDRContains(t *testing.T) {
	prefix, err := netip.ParsePrefix("10.25.10.1/24")
	if err != nil {
		t.Error(err)
	}
	t.Log(prefix)
	addr, err := netip.ParseAddr("10.25.10.1")

	t.Log(addr)
	pre := netip.PrefixFrom(addr, 20)
	t.Log(pre)
	t.Log(pre.Addr())
	t.Log(prefix.Contains(pre.Addr()))
	t.Log(pre.Contains(prefix.Addr()))
}
