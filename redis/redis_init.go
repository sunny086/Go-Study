package redis

import "github.com/go-redis/redis/v8"

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
