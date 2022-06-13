package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("main....")
}

func init() {
	log.Println("init")
}

func init() {
	fmt.Println("init....")
}
