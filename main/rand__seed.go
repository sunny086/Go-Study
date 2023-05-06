package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nano := time.Now().UnixNano()
	fmt.Println(nano)
	rand.Seed(time.Now().UnixNano())
	randInsert := rand.Intn(1000) // [0, 4)
	fmt.Println(randInsert)
}
