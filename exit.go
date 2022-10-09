package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("hello world")
	runtime.Goexit()
	fmt.Println("hello world")
	fmt.Println("hello world")
	//hello world
	//hello world
	//fatal error: no goroutines (main called runtime.Goexit) - deadlock!
}
