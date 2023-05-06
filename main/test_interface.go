package main

type test01 interface {
	fuc1()
	fuc2()
	fuc3()
}

type t1 struct {
}

func (*t1) fuc1() {
	println("fuc1")
}

func (*t1) fuc2() {
	println("fuc2")
}

func (*t1) fuc3() {
	println("fuc3")
}

type t2 struct {
}

func (*t2) fuc1() {
	println("fuc1")
}

func (*t2) fuc2() {
	println("fuc2")
}

func (*t2) fuc3() {
	println("fuc3")
}

func NewTest01() test01 {
	var a int = 1
	if a == 1 {

		return &t1{}
	} else {
		return &t2{}
	}
}

func main() {
	t := NewTest01()
	t.fuc2()
}
