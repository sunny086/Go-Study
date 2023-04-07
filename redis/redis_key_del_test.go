package redis

import (
	"context"
	"testing"
)

func TestRedisKeyDel(t *testing.T) {
	client.Set(context.Background(), "test:key1", "value1", 0)
	client.Set(context.Background(), "test:key2", "value1", 0)
	client.Set(context.Background(), "test:key3", "value1", 0)
	client.Set(context.Background(), "test:key4", "value1", 0)
	//删除test前缀的key
	keys, _ := client.Keys(context.Background(), "test:*").Result()
	client.Del(context.Background(), keys...)
	for _, key := range keys {
		t.Log("key:", key)
	}
}
