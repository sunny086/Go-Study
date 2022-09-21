package refelct

import (
	"fmt"
	"reflect"
	"testing"
)

type user struct {
	Name string
	Age  int
}

func (u *user) ShowName() {
	fmt.Println(u.Name)
}

func (u *user) AddAge(addNum int) {
	fmt.Println("age add result:", u.Age+addNum)
}

func TestReflect_Method(t *testing.T) {
	u := &user{"lisi", 20}
	v := reflect.ValueOf(u)
	// 调用无参方法
	methodV := v.MethodByName("ShowName")
	methodV.Call(nil) // 或者传递一个空切片也可

	// 调用有参方法
	methodV2 := v.MethodByName("AddAge")
	args := []reflect.Value{reflect.ValueOf(30)} //
	methodV2.Call(args)
}
