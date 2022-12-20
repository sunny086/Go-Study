package net

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"strings"
	"testing"
)

func TestIp2Int1(t *testing.T) {
	t.Log(Ip2Int1("10.25.10.1"))
	//两种计算逻辑一致
	t.Log(Ip2Int2("10.25.10.1"))
}

func Ip2Int1(ip string) uint32 {
	parseIP := net.ParseIP(ip)
	if parseIP.To4() != nil {
		return binary.BigEndian.Uint32(parseIP.To4())
	}
	return 0
}

func Ip2Int2(ip string) int64 {
	string2Int := func(in string) (out int) {
		out, _ = strconv.Atoi(in)
		return
	}

	if len(ip) == 0 {
		return 0
	}
	bits := strings.Split(ip, ".")
	if len(bits) < 4 {
		return 0
	}
	b0 := string2Int(bits[0])
	b1 := string2Int(bits[1])
	b2 := string2Int(bits[2])
	b3 := string2Int(bits[3])

	var sum int64
	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

func TestIp2Int2(t *testing.T) {
	// IP地址转换为uint32
	IP1 := net.ParseIP("192.168.8.44")
	IPUint32 := IPToUInt32(IP1)
	fmt.Println(IPUint32)

	// uint32转换为IP地址
	IP2 := UInt32ToIP(IPUint32)
	fmt.Println(IP2.String())
}

// UInt32ToIP uint32类型转换为IP
func UInt32ToIP(intIP uint32) net.IP {
	var bytes [4]byte
	bytes[0] = byte(intIP & 0xFF)
	bytes[1] = byte((intIP >> 8) & 0xFF)
	bytes[2] = byte((intIP >> 16) & 0xFF)
	bytes[3] = byte((intIP >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

// IPToUInt32 IP转换为uint32类型
func IPToUInt32(ipnr net.IP) uint32 {
	bits := strings.Split(ipnr.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum uint32

	sum += uint32(b0) << 24
	sum += uint32(b1) << 16
	sum += uint32(b2) << 8
	sum += uint32(b3)

	return sum
}
