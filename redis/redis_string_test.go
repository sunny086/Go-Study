package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
)

//func NewClient() *redis.Client {
//	client := redis.NewClient(&redis.Options{
//		Addr:     "10.25.10.126:6379",
//		Password: "Netvine123#@!", // no password set
//		DB:       4,               // use default DB
//	})
//	return client
//}

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "10.25.10.126:6379",
		Password: "Netvine123#@!", // no password set
		DB:       4,               // use default DB
	})
}

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
