package array

import (
	"fmt"
	"testing"
)

func TestArrayPointer(t *testing.T) {
	var arr = [5]int{1, 2, 3, 4, 5}
	modifyArr1(arr)
	fmt.Println(arr)
	modifyArr2(&arr)
	fmt.Println(arr)
	//[1 2 3 4 5]
	//[1 20 3 4 5]
}

// 传递的是变量副本，不会改变原数组
func modifyArr1(a [5]int) {
	a[1] = 20
}

// 传递的是指针，会改变原数组
func modifyArr2(a *[5]int) {
	a[1] = 20
}

type Person struct {
	Name string
	Age  int
}

//func Test01(t *testing.T) {
//	var p1 = Person{"Tom", 20}
//	var p2 = Person{
//		Name: "Go",
//		Age:  21,
//	}
//	//var ps = []Person{p1, p2}
//	//for i := 0; i < len(ps); i++ {
//	//	ps[i].Age = 22
//	//}
//	//fmt.Println(ps)
//
//	var ps []*Person
//	ps = append(ps, &p1)
//	ps = append(ps, &p2)
//	for i := 0; i < len(ps); i++ {
//		ps[i].Age = 22
//
//	}
//	fmt.Println(ps)
//
//	//var p1 = Person{"Tom", 20}
//	//var p2 = Person{
//	//	Name: "Go",
//	//	Age:  21,
//	//}
//	var ps2 = []Person{p1, p2}
//	for i, _ := range ps2 {
//		ps2[i].Age = 23
//	}
//	fmt.Println(ps2)
//
//}

// 直接声明对象数组 调用函数传递 通过下标可以修改
func Test01(t *testing.T) {
	var p1 = Person{"Tom", 20}
	var p2 = Person{
		Name: "Go",
		Age:  21,
	}
	var ps = []Person{p1, p2}
	modifyPerson1(ps)
	fmt.Println(ps)
}

// 直接数组传递过来 通过下标修改 可以修改
func modifyPerson1(ps []Person) {
	for i := 0; i < len(ps); i++ {
		ps[i].Age = 99
	}
	fmt.Println(ps)
}

func Test02(t *testing.T) {
	var p1 = Person{"Tom", 20}
	var p2 = Person{
		Name: "Go",
		Age:  21,
	}
	var ps = []Person{p1, p2}
	modifyPerson2(ps)
	fmt.Println(ps)
}

// 直接数组传递过来 range遍历不使用下标 修改失败
func modifyPerson2(ps []Person) {
	for _, p := range ps {
		p.Age = 99
	}
	fmt.Println(ps)
}

func Test03(t *testing.T) {
	var p1 = Person{"Tom", 20}
	var p2 = Person{
		Name: "Go",
		Age:  21,
	}
	//var ps = []*Person{&p1, &p2}
	var ps []*Person
	ps = append(ps, &p1)
	ps = append(ps, &p2)
	modifyPerson3(ps)
	fmt.Println(ps)
}

// 传递指针数组 range遍历不使用下标 修改成功
func modifyPerson3(ps []*Person) {
	//for i := 0; i < len(ps); i++ {
	//	ps[i].Age = 99
	//}
	//fmt.Println(ps)

	for _, p := range ps {
		p.Age = 99
	}
	fmt.Println(ps)
}
