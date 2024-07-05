package dw

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	unix := time.Now().Unix()
	t.Log(unix)
	t.Log(time.Unix(unix, 0).Format("2006-01-02 15:04:05"))
	milli := time.Now().UnixMilli()
	t.Log(milli)
	t.Log(time.UnixMilli(milli).Format("2006-01-02 15:04:05"))
}
