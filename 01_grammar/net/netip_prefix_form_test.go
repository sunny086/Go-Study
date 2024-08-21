package net

import (
	"net/netip"
	"testing"
)

// 直接解析cidr字符串返回netip.Prefix，也可以使用Addr和bitSize进行PrefixForm返回netip.Prefix
func TestCompareParseFormAndParsePrefix(t *testing.T) {
	prefix, _ := netip.ParsePrefix("10.25.10.1/24")
	t.Log(prefix) //10.25.10.1/24
	//对比
	addr, _ := netip.ParseAddr("10.25.17.1")
	prefixFrom := netip.PrefixFrom(addr, 24)
	t.Log(prefixFrom) //10.25.17.1/24
}
