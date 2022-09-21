package refelct

import (
	"fmt"
	"reflect"
	"testing"
)

func add(name string, age int) {
	fmt.Printf("name is %s, age is %d \n", name, age)
}

func TestReflect_Func(t *testing.T) {
	funcValue := reflect.ValueOf(add)
	params := []reflect.Value{reflect.ValueOf("lisi"), reflect.ValueOf(20)}
	reList := funcValue.Call(params)
	fmt.Println(reList)
}
