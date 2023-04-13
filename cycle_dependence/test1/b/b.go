package b

import (
	"GoTest/cycle_dependence/test1/a_iface"
	"fmt"
)

func FuncB() {
	a_iface.GetA().FuncA()
}

type AImpl struct{}

func (a *AImpl) FuncA() {
	fmt.Println("AImpl.FuncA()")
}

func init() {
	a_iface.SetA(&AImpl{})
}
