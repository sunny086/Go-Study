package net

import (
	"fmt"
	"net/netip"
	"testing"
)

/*
10.25.10.1-10.25.10.10
10.25.10.15
10.25.10.1/24
10.25.11.1-10.25.11.3
10.25.11.7

合并后

10.25.10.1/24
10.25.11.1-10.25.11.3
10.25.11.7
首先预处理单IP，变成二元组。
然后对每一个二元组按left升序排序，left相同按right升序排序。
扫描线算法，从左到右枚举每一个区间左端点，对每一个区间左端点，向右扫描看是否有第二个区间的左端点在本区间内。
如果有一个，判断是否需要扩展本区间；如果有多个，找到被包含多个左断点的最大右端点，判断是否需要扩展本区间。如果没有，处理下一个左端点。
最后遍历结果二元组，把left right相同的二元组处理成一个数据。
递归处理结果二元组，直到没有重叠的区间。
*/
func TestMergeIp(t *testing.T) {
	r := MergeIPs([]string{
		"10.25.10.1",
		"10.25.10.10",
		"10.25.10.15",
		"10.25.10.1/24",
		"10.25.10.1/28",
		"10.25.11.1",
	})

	for _, v := range r {
		fmt.Println(v.String())
	}
}

func MergeIPs(ips []string) []netip.Prefix {
	var m = make(map[netip.Prefix]struct{})
	var um = func(n netip.Prefix) {
		for oln := range m {
			if n.Contains(oln.Addr()) {
				if n.Bits() < oln.Bits() {
					delete(m, oln)
				} else {
					return
				}
			}
		}
		m[n] = struct{}{}
	}

	for _, ip := range ips {
		pre, err := netip.ParsePrefix(ip)
		if err != nil {
			pre, err = netip.ParsePrefix(ip + "/32")
			if err != nil {
				panic("")
			}
		}
		um(pre)
	}

	var r []netip.Prefix
	for k, _ := range m {
		r = append(r, k)
	}
	return r
}
