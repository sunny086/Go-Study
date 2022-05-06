package main

import (
	"fmt"
	"net"
)

func main() {

	mask := net.CIDRMask(24, 32)

	fmt.Println(mask)

	cidr, ipNet, err := net.ParseCIDR("255.255.0.0/9")
	fmt.Println(cidr, ipNet, err)

}
