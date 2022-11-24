package string

import (
	"strings"
	"testing"
)

// 1.len的长度是按照字节来计算的，而不是字符
// 2.len的长度在切片里面是索引-1的，可以理解为左闭右开，开的位置正好是索引的位置=====>这里要区分长度和索引位置的关系，假设长度为5，索引位置为0-4，第五位索引位置是4
// 3.strings.Index()返回的是索引位置，而不是长度；切片里面传len的话就是+1后的值
func TestStringIndex(t *testing.T) {
	line := "Oct 17 17:12:11 ubuntu kernel: [  722.895541] Ip-Mac-BLackList set:IN=enp1s0 OUT= MAC=90:b8:e0:01:4b:05:8c:1c:da:42:5d:77:08:00 SRC=10.25.10.125 DST=10.25.10.126 LEN=60 TOS=0x00 PREC=0x00 TTL=64 ID=52507 DF PROTO=TCP SPT=38980 DPT=22 WINDOW=29200 RES=0x00 SYN URGP=0"
	//截图SRC和DST之间的ip
	start := strings.Index(line, "SRC=")
	end := strings.Index(line, "DST=")
	if start >= 0 && end >= 0 {
		ip := line[start+4 : end-1]
		println(ip)
	}
	//截图MAC和SRC之间的mac
	macStartIndex := strings.Index(line, "MAC=")
	macEndIndex := strings.Index(line, "SRC=")
	if macStartIndex >= 0 && macEndIndex >= 0 {
		mac := line[macStartIndex+4 : macStartIndex+4+17]
		println(mac)
	}
}
