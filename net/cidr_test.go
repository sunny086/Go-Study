package net

import (
	"net"
	"strconv"
	"strings"
	"testing"
)

func TestCIDR(t *testing.T) {
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
