package main

import (
	"fmt"
	"reflect"
)

type PeopleReflect struct {
	Name string
}

func (p *PeopleReflect) Eat() {
	fmt.Println("people eat")
}

func main() {
	of := reflect.ValueOf(&PeopleReflect{})
	of.MethodByName("Eat").Call([]reflect.Value{})
}
