package a

import (
	"Go-Study/cycle_dependence/test1/b_iface"
	"fmt"
)

func FuncA() {
	b_iface.GetB().FuncB()
}

type BImpl struct{}

func (b *BImpl) FuncB() {
	fmt.Println("BImpl.FuncB()")
}

func init() {
	b_iface.SetB(&BImpl{})
}
