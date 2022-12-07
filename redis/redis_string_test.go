package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
)

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.25.10.126:6379",
		Password: "Netvine123#@!", // no password set
		DB:       4,               // use default DB
	})
	return client
}

func TestString(t *testing.T) {
	client := NewClient()
	val := client.Get(context.Background(), "addressBlacklist:22.22.22.22").Val()
	t.Log("val:", val)
}
