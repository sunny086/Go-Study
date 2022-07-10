package main

import "fmt"

func main() {

	// Compile the program
	//go build -o compile
	//new make 区别
	//go tool compile -S .\compile.go | grep CALL
	s := new(string)
	fmt.Println(s)
}
