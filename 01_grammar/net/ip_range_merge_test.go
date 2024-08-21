package net

import (
	"encoding/binary"
	"fmt"
	"net"
	"net/netip"
	"sort"
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
func TestMergeIp1(t *testing.T) {
	r := MergeIPs2([]string{
		"10.25.10.1/24",
		"10.25.10.1/23",
		"10.25.10.10/25",
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

func MergeIPs2(ips []string) []netip.Prefix {
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
			fmt.Println(err)
			addr, err := netip.ParseAddr(ip)
			if err != nil {
				fmt.Println(err)
			}
			pre = netip.PrefixFrom(addr, addr.BitLen())
		}
		um(pre)
	}

	var r []netip.Prefix
	for k, _ := range m {
		r = append(r, k)
	}
	return r
}

func TestMergeIp2(t *testing.T) {
	//区分单个ip ip段和cidr形式的ip
	source := []string{
		"10.25.10.1",
		"10.25.10.2",
		"10.25.10.3",
		"10.25.10.1",
		"10.25.10.2",
		"10.25.10.3",
		"10.25.10.1/24",
		"10.25.11.1",
		"10.25.11.2",
		"10.25.11.3",
		"10.25.11.4",
		"10.25.11.5",
		"10.25.11.6",
		"10.25.11.7",
		"10.25.11.8",
		"10.25.11.9",
		"10.25.11.10",
		"10.25.11.1-10.25.11.8",
		"10.25.17.10-10.25.17.11",
		"10.25.16.1/255.255.255.0",
		"10.25.11.1-10.25.11.8",
		"10.25.17.10-10.25.17.11",
		"10.25.16.1/255.255.255.0",
		"10.25.11.1/24",
	}
	singleIpList, ipRangeList, cidrList, err := DivideIp(source)
	if err != nil {
		t.Log(err)
		return
	}
	//三个ip进行去重
	ipRangeList = ipRangeList[:RemoveDuplicateData(ipRangeList)]
	cidrList = cidrList[:RemoveDuplicateData(cidrList)]
	//ip段转成单个ip
	ipRange2IpList, err := IPRange2Ip(ipRangeList)
	singleIpList = append(singleIpList, ipRange2IpList...)
	//合并之后再去重
	singleIpList = singleIpList[:RemoveDuplicateData(singleIpList)]
	//cidr合并去重
	cidrDuplicateList, ipNetList, err := RemoveDuplicateCIDR(cidrList)
	//判断每个ip是否在cidr中
	var singleIpDuplicateList []string
	for _, ip := range singleIpList {
		var b bool
		for _, ipNet := range ipNetList {
			if ipNet.Contains(net.ParseIP(ip)) {
				b = true
				//去除重复的单个ip
				continue
			}
		}
		if !b {
			singleIpDuplicateList = append(singleIpDuplicateList, ip)
		}
	}
	t.Log(cidrDuplicateList)
	t.Log(singleIpDuplicateList)
	res := append(cidrDuplicateList, singleIpDuplicateList...)
	t.Log(res)
}

// 测试IP去重取并集
func TestDivideIP(t *testing.T) {
	//区分单个ip ip段和cidr形式的ip
	source := []string{
		"10.25.10.1",
		"10.25.10.2",
		"10.25.10.3",
		"10.25.10.1",
		"10.25.10.2",
		"10.25.10.3",
		"10.25.10.1/24",
		"10.25.11.1",
		"10.25.11.2",
		"10.25.11.3",
		"10.25.11.4",
		"10.25.11.5",
		"10.25.11.6",
		"10.25.11.7",
		"10.25.11.8",
		"10.25.11.9",
		"10.25.11.10",
		"10.25.11.1-10.25.11.8",
		"10.25.17.10-10.25.17.70",
		"10.25.16.1/255.255.255.0",
	}
	singleIpList, ipRangeList, cidrList, err := DivideIp(source)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log(singleIpList)
	t.Log(ipRangeList)
	t.Log(cidrList)

}

func DivideIp(source []string) (singleIpList, ipRangeList, cidrList []string, err error) {
	//var prefix netip.Prefix
	for _, ip := range source {
		if strings.Contains(ip, "-") {
			ipRangeList = append(ipRangeList, ip)
		} else if strings.Contains(ip, "/") {
			cidr, err := ParseCIDR(ip)
			if err != nil {
				return nil, nil, nil, err
			}
			cidrList = append(cidrList, cidr)
		} else {
			singleIpList = append(singleIpList, ip)
		}
	}
	return
}

// 处理255.255.255.0后缀的掩码，统一返回cidr形式的掩码
func ParseCIDR(cidrIp string) (cidrParseIp string, err error) {
	ipSlice := strings.Split(cidrIp, "/")
	ip := ipSlice[0]
	cidr := ipSlice[1]
	if strings.Contains(cidr, ".") {
		//10.25.10.1/255.255.255.0形式的ip转换成cidr形式
		cidrBitSize := SubnetMask2CIDR(cidr)
		_, ipNet, err := net.ParseCIDR(ip + "/" + strconv.Itoa(cidrBitSize))
		if err != nil {
			return "", err
		}
		cidrParseIp = ipNet.String()
	} else {
		//10.25.10.1/24形式的ip 不需要处理
		cidrParseIp = cidrIp
	}
	return
}

func RemoveDuplicateData(strList []string) int {
	sort.Strings(strList)
	//如果是空切片，那就返回0
	if len(strList) == 0 {
		return 0
	}
	//用两个标记来比较相邻位置的值
	//当一样的话，那就不管继续
	//当不一样的时候，就把right指向的值赋值给left下一位
	left, right := 0, 1
	for ; right < len(strList); right++ {
		if strList[left] == strList[right] {
			continue
		}
		left++
		strList[left] = strList[right]
	}
	return left + 1
}

func TestIpRange2IpList(t *testing.T) {
	source := []string{
		"10.25.10.1-10.25.10.3",
		"10.25.17.1-10.25.17.2",
	}
	ipList, err := IPRange2Ip(source)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(ipList)
}

// ip段转换成单个ip
func IPRange2Ip(ipRangeList []string) (ipList []string, err error) {
	for _, ipRange := range ipRangeList {
		ipRangeSlice := strings.Split(ipRange, "-")
		startIp := ipRangeSlice[0]
		endIp := ipRangeSlice[1]
		startNum := Ip2Int(startIp)
		endNum := Ip2Int(endIp)
		addr, err := netip.ParseAddr(startIp)
		if err != nil {
			return nil, err
		}
		for i := startNum; i <= endNum; i++ {
			//数字转换成ip
			ipList = append(ipList, addr.String())
			addr = addr.Next()
		}
	}
	return
}

func Ip2Int(ip string) uint32 {
	parseIP := net.ParseIP(ip)
	if parseIP.To4() != nil {
		return binary.BigEndian.Uint32(parseIP.To4())
	}
	return 0
}

func TestCIDRContains1(t *testing.T) {
	source := []string{
		"10.25.10.1/24",   //10.25.10.1-10.25.10.254
		"10.25.10.10/24",  //10.25.10.1-10.25.10.254
		"10.25.11.128/25", //10.25.11.129-10.25.11.255
		"10.25.10.128/23", //10.25.10.1-10.25.11.254
	}
	cidrStrList, netList, err := RemoveDuplicateCIDR(source)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(cidrStrList)
	t.Log(netList)
}

func RemoveDuplicateCIDR(cidrList []string) (cidrDuplicateList []string, cidrDuplicateIpNetList []net.IPNet, err error) {
	var netIpList []net.IPNet
	for _, cidr := range cidrList {
		_, ipNet, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, nil, err
		}
		netIpList = append(netIpList, *ipNet)
	}
	left := CompareAndRemoveDuplicateCIDR(netIpList)
	netIpList = netIpList[:left]
	//翻转切片
	for i := len(netIpList)/2 - 1; i >= 0; i-- {
		opp := len(netIpList) - 1 - i
		netIpList[i], netIpList[opp] = netIpList[opp], netIpList[i]
	}
	left = CompareAndRemoveDuplicateCIDR(netIpList)
	netIpList = netIpList[:left]
	for _, netIp := range netIpList {
		cidrDuplicateList = append(cidrDuplicateList, netIp.String())
		cidrDuplicateIpNetList = append(cidrDuplicateIpNetList, netIp)
	}
	return
}

func CompareAndRemoveDuplicateCIDR(netIpList []net.IPNet) int {
	left, right := 0, 1
	for ; right < len(netIpList); right++ {
		ones1, _ := netIpList[left].Mask.Size()
		ones2, _ := netIpList[right].Mask.Size()
		if ones1 <= ones2 && netIpList[left].Contains(netIpList[right].IP) {
			continue
		}
		left++
		netIpList[left] = netIpList[right]
	}
	return left + 1
}
