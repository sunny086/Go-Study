package a_iface

type AInterface interface {
	FuncA()
}

var a AInterface

func SetA(inA AInterface) {
	a = inA
}

func GetA() AInterface {
	return a
}
