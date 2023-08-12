package main

import (
	"Go-Study/cycle_dependence/test1/a"
	"Go-Study/cycle_dependence/test1/b"
)

func main() {
	a.FuncA()
	b.FuncB()
}
