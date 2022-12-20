package net

import (
	"net/netip"
	"testing"
)

// netip.ParsePrefix 不能解析255.255.255.0这种后缀的掩码
func TestParsePrefix1(t *testing.T) {
	_, err := netip.ParsePrefix("10.25.10.1/255.255.255.0")
	if err != nil {
		t.Error(err) //netip.ParsePrefix("10.25.10.1/255.255.255.0"): bad bits after slash: "255.255.255.0"
	}
}

// netip.ParsePrefix 可以解析标准后缀的掩码，可以获取到掩码的长度和掩码的ip地址以及字符串形式的掩码
func TestParsePrefix2(t *testing.T) {
	prefix, _ := netip.ParsePrefix("10.25.10.1/24")
	t.Log(prefix)          //10.25.10.1/24
	t.Log(prefix.Addr())   //10.25.10.1
	t.Log(prefix.Bits())   //24
	t.Log(prefix.String()) //10.25.10.1/24
}

// netip.Prefix 可以判断子网掩码书写的ip是否是单个ip 换句话说是否是32后缀的
func TestPrefix1(t *testing.T) {
	prefix, _ := netip.ParsePrefix("10.25.10.1/32")
	ip := prefix.IsSingleIP()
	t.Log(ip) //true
}
