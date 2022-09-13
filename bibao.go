package main

var x int

func Increase() func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	increase := Increase()
	println(increase())
	println(increase())
	println(increase())
}
