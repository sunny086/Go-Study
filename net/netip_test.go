package net

import (
	"net/netip"
	"testing"
)

// netip.ParsePrefix不能解析255.255.255.0这种后缀的掩码
func TestParsePrefix1(t *testing.T) {
	_, err := netip.ParsePrefix("10.25.10.1/255.255.255.0")
	if err != nil {
		t.Error(err) //netip.ParsePrefix("10.25.10.1/255.255.255.0"): bad bits after slash: "255.255.255.0"
	}
}
