package refelct

import (
	"fmt"
	"reflect"
	"testing"
)

type user struct {
	Name  string
	Age   int
	Score int
}

func (u *user) ShowName() {
	fmt.Println(u.Name)
}

func (u *user) AddAge(age int) {
	fmt.Println("age add result:", u.Age+age)
}

func (u *user) AddAgeAndScore(age, score int) {
	u.Age = u.Age + age
	u.Score = u.Score + score
	fmt.Println("score add result:", u.Score+score)
	fmt.Println("age add result:", u.Age+age)
}

func TestReflect_Method(t *testing.T) {
	u := &user{"lisi", 20, 90}
	v := reflect.ValueOf(u)
	// 调用无参方法
	methodV := v.MethodByName("ShowName")
	methodV.Call(nil) // 或者传递一个空切片也可
	// 调用有参方法
	methodV2 := v.MethodByName("AddAge")
	args := []reflect.Value{reflect.ValueOf(30)} //
	methodV2.Call(args)
	// 调用有参方法
	methodV3 := v.MethodByName("AddAgeAndScore")
	args2 := []reflect.Value{reflect.ValueOf(30), reflect.ValueOf(10)}
	methodV3.Call(args2)
	fmt.Println(v)
}
