package dcp

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	str := "2024-01-01"
	parse, _ := time.Parse("2006-01-02", str)
	t.Log(parse)
	year, month, day := parse.Date()
	t.Log(year, month, day)

	month = month - 1
	if month == 0 {
		year = year - 1
		month = 12
	}

	//if month == 1 && day == 1 {
	//	year = year - 1
	//	month = 12
	//} else {
	//	month = month - 1
	//}
	t.Log(year, month, day)
}
