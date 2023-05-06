package main

import "testing"

func TestNewMakeCompileGrepCall(t *testing.T) {
	// Compile the program
	//go build -o compile
	//new make 区别
	//go tool compile -S .\compile.go | grep CALL
	s := new(string)
	t.Log(s)
}
