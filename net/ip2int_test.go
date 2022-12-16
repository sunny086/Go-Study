package net

import (
	"encoding/binary"
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
