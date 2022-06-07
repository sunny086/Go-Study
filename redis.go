package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.25.10.216:6379",
		Password: "Netvine123#@!", // no password set
		DB:       3,               // use default DB
	})
	//_, err := client.Ping(context.Background()).Result()
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//val := client.Scan(context.Background(), 0, "trustHost:*", 500).Iterator().Val()
	result, _ := client.Keys("trustHost:*").Result()
	fmt.Println(result)
	//}
	for _, key := range result {
		var err error
		var ip64 int64
		var hashKeys []string
		//当前这个hash其实只有一个键值对 username和ip64，所以取的是数组第一个元素
		hashKeys, err = client.HKeys(key).Result()
		username := hashKeys[0]
		ip64, err = client.HGet(key, username).Int64()

		fmt.Println("=============")
		fmt.Println(err)
		fmt.Println(ip64)
		fmt.Println("=============")
	}
}
