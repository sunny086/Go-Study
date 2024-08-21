package refelct

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Name  string
	Age   int
	Score int
}

func (u *User) ShowName() {
	fmt.Println(u.Name)
}

func (u *User) AddAge(age int) {
	fmt.Println("age add result:", u.Age+age)
}

func (u *User) AddAgeAndScore(age, score int, str string) {
	u.Age = u.Age + age
	u.Score = u.Score + score
	fmt.Println("score add result:", u.Score+score)
	fmt.Println("age add result:", u.Age+age)
	fmt.Println("str:", str)
}

// TestReflect_Execute_Method 执行结构体的方法
func TestReflect_Execute_Method(t *testing.T) {
	u := &User{"lisi", 20, 90}
	v := reflect.ValueOf(u)
	// 调用无参方法
	methodV := v.MethodByName("ShowName")
	methodV.Call(nil) // 或者传递一个空切片也可
	// 调用有参方法
	methodV2 := v.MethodByName("AddAge")
	args := []reflect.Value{reflect.ValueOf(30)} //
	methodV2.Call(args)
	// 调用有参方法 参数需要按顺序传递
	methodV3 := v.MethodByName("AddAgeAndScore")
	args2 := []reflect.Value{reflect.ValueOf(30), reflect.ValueOf(10), reflect.ValueOf("调用有参方法 参数需要按顺序传递")}
	methodV3.Call(args2)
	fmt.Println(v)
}

// TestReflect_GetAllMethod 获取结构体的所有方法
func TestReflect_GetAllMethod(t *testing.T) {
	v := reflect.TypeOf(&User{})
	// 获取所有方法
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		marshal, err := json.Marshal(method)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(marshal))
		}
		fmt.Println("method name:", method.Name)
	}
}
