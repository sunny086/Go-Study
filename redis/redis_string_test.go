package redis

import (
	"context"
	"testing"
)

func TestString(t *testing.T) {
	val := client.Get(context.Background(), "auxConfig:upperStorageLimit").Val()
	t.Log("val:", val)
}

func TestStringArgs(t *testing.T) {
	args := client.Get(context.Background(), "auxConfig:upperStorageLimit").Args()
	t.Log("args:", args)
}

func TestStringScan(t *testing.T) {
	var limit int64
	err := client.Get(context.Background(), "auxConfig:upperStorageLimit").Scan(&limit)
	if err != nil {
		t.Log("err:", err)
	}
	t.Log("limit:", limit)
}
