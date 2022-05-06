package main

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"strconv"
	"strings"
)

func main() {
	var CIDR = "^(?:(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\/([1-9]|[1-2]\\d|3[0-2])$"

	regexp2 := regexp2.MustCompile(CIDR, 0)
	matchString, err := regexp2.MatchString("255.128.0.0/9")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(matchString)
	}

	fmt.Println("==========================")

	ipSlice := strings.Split("255.128.0.0/9", `/`)
	ip := ipSlice[0]
	cidr := ipSlice[1]
	ipInt := ip2Int(ip)
	cidrInt, _ := strconv.Atoi(cidr)
	fmt.Println(ipInt)
	endIp := ipInt + (1 << uint32(32-cidrInt)) - 2
	fmt.Println(endIp)

	fmt.Println(ip2Int("255.128.0.1"))
	fmt.Println(ip2Int("255.255.255.254"))

}

func ip2Int(ip string) int64 {
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

func string2Int(in string) (out int) {
	out, _ = strconv.Atoi(in)
	return
}
