package main

import "fmt"

type Person1 interface {
	SayHello(name string)
	Run() string
}

type Hero struct {
	Name  string
	Speed int
}

func (h *Hero) Run() string {
	return "run"
}

func (h *Hero) SayHello(name string) {
	fmt.Println("hello", name)
}

func main() {
	var hero Hero
	hero.Name = "hero"
	hero.Speed = 100
	var person Person1 = &hero
	person.SayHello("hero")
}
