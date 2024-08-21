package main

import (
	"Go-Study/01_grammar/init1"
	"fmt"
	"log"
)

func main() {
	fmt.Println("main....")
	init1.Test()
}

func init() {
	log.Println("init")
}

func init() {
	fmt.Println("init....")
}
