package main

import (
	"GoTest/cycle_dependence/test1/a"
	"GoTest/cycle_dependence/test1/b"
)

func main() {
	a.FuncA()
	b.FuncB()
}
