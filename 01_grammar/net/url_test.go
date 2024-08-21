package net

import (
	"fmt"
	"net/url"
	"testing"
)

func TestNetUrl(t *testing.T) {
	escape := url.QueryEscape("123" + "策略.xlsx")

	fmt.Println(escape)
}
