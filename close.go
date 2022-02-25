package main

import "fmt"

func main() {

	c := a()
	fmt.Println(&c)
	fmt.Printf("c: %p\n:", c)
	c()
	c()
	c()

}

func a() func() int {
	i := 0
	fmt.Printf("a function i: %p\n", &i)
	b := func() int {
		i++
		fmt.Printf("b function i: %p\n", &i)
		return i
	}
	return b

}
