package refelct

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect_ValueOf01(t *testing.T) {
	//ValueOf()：获取变量的值，即pair中的 value
	//TypeOf()：获取变量的类型，即pair中的 concrete type
	type Person struct {
		Name string
		Age  int
	}
	p := Person{"lisi", 13}
	personValueOf := reflect.ValueOf(p)
	fmt.Println(personValueOf)        // {lisi 13}  变量的值
	fmt.Println(personValueOf.Type()) // refelct.Person 变量类型的对象名
}
