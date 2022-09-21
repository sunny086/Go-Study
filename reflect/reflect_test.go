package refelct

import (
	"fmt"
	"reflect"
	"testing"
)

type SysDictDataGetPageReq struct {
	Id        int    `form:"id" search:"type:exact;column:dict_code;table:sys_dict_data" comment:""`
	DictLabel string `form:"dictLabel" search:"type:contains;column:dict_label;table:sys_dict_data" comment:""`
	DictValue string `form:"dictValue" search:"type:contains;column:dict_value;table:sys_dict_data" comment:""`
	DictType  string `form:"dictType" search:"type:contains;column:dict_type;table:sys_dict_data" comment:""`
	Status    string `form:"status" search:"type:exact;column:status;table:sys_dict_data" comment:""`
}

func reflectTest04() {
	var d SysDictDataGetPageReq
	s := reflect.TypeOf(d).String()
	fmt.Println(s)
	qType := reflect.TypeOf(&d).Elem()
	fmt.Println(qType)
	name := qType.Name()
	fmt.Println(name)
	fmt.Println(qType.NumField())
	fmt.Println(qType.NumMethod())

}

func reflectTest03() {

	type user struct {
		Name string
		Age  int `json:"age" id:"100"` // 结构体标签
	}

	s := user{
		Name: "zs",
		Age:  1,
	}

	typeOfUser := reflect.TypeOf(s)

	// 字段用法
	for i := 0; i < typeOfUser.NumField(); i++ { // NumField 当前结构体有多少个字段
		fieldType := typeOfUser.Field(i) // 获取每个字段
		fmt.Println(fieldType.Name, fieldType.Tag)
	}
	if userAge, ok := typeOfUser.FieldByName("Age"); ok {
		fmt.Println(userAge) // {Age  int json:"age" id:"100" 16 [1] false}
	}

	// 方法用法
	for i := 0; i < typeOfUser.NumMethod(); i++ {
		fieldType := typeOfUser.Method(i) // 获取每个字段
		fmt.Println(fieldType.Name)
	}

}

func reflectTest02() {
	var num int64 = 100
	// 设置值：指针传递
	ptrValue := reflect.ValueOf(&num)
	newValue := ptrValue.Elem()                // Elem()用于获取原始值的反射对象
	fmt.Println("type：", newValue.Type())      // int64
	fmt.Println("can set：", newValue.CanSet()) // true
	newValue.SetInt(200)

	// 获取值：值传递
	rValue := reflect.ValueOf(num)
	fmt.Println(rValue.Int())               // 方式一：200
	fmt.Println(rValue.Interface().(int64)) // 方式二：200

}

func TestReflect01(t *testing.T) {
	//ValueOf()：获取变量的值，即pair中的 value
	//TypeOf()：获取变量的类型，即pair中的 concrete type
	type Person struct {
		Name string
		Age  int
	}
	p := Person{"lisi", 13}

	fmt.Println(reflect.ValueOf(p))        // {lisi 13}  变量的值
	fmt.Println(reflect.ValueOf(p).Type()) // main.Person 变量类型的对象名

	fmt.Println(reflect.TypeOf(p)) //  main.Person	变量类型的对象名

	fmt.Println(reflect.TypeOf(p).Name()) // Person:变量类型对象的类型名
	fmt.Println(reflect.TypeOf(p).Kind()) // struct:变量类型对象的种类名

	fmt.Println(reflect.TypeOf(p).Name() == "Person")       // true
	fmt.Println(reflect.TypeOf(p).Kind() == reflect.Struct) //true

}

// 通过反射获取结构体的type类型（package.structName）和kind种类(struct、ptr....)
func TestReflect_TypeOf01(t *testing.T) {
	//TypeOf()：获取变量的类型，即pair中的 concrete type
	type Person struct {
		Name string
		Age  int
	}
	person := Person{"lisi", 13}
	//  person's type is refelct.Person ,kind is struct
	fmt.Printf("person's type is %s ,kind is %s \n", reflect.TypeOf(person), reflect.TypeOf(person).Kind())
	//  *person's type is *refelct.Person ,kind is ptr
	fmt.Printf("*person's type is %s ,kind is %s \n", reflect.TypeOf(&person), reflect.TypeOf(&person).Kind())
}

//对于指针类型的变量，可以使用Type.Elem获取到指针指向变量的真实类型对象
func TestReflect_TypeOf_Elem(t *testing.T) {
	//TypeOf()：获取变量的类型，即pair中的 concrete type
	type Person struct {
		Name string
		Age  int
	}
	person := Person{"lisi", 13}
	personTypeOf := reflect.TypeOf(&person)
	fmt.Printf("*person's type is %s ,kind is %s \n", personTypeOf, personTypeOf.Kind())
	elem := reflect.TypeOf(&person).Elem()
	//*person's type is refelct.Person ,kind is struct
	fmt.Printf("*person's type is %s ,kind is %s \n", elem, elem.Kind())
}
