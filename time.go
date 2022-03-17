package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	duration, _ := time.ParseDuration("30m")
	add := now.Add(duration)
	fmt.Println(add)
}
