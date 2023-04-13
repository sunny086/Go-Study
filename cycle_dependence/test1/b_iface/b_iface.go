package b_iface

type BInterface interface {
	FuncB()
}

var b BInterface

func SetB(inB BInterface) {
	b = inB
}

func GetB() BInterface {
	return b
}
