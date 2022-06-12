package main

import (
	"fmt"
	"runtime"
)

func main() {

	cpu := runtime.NumCPU
	fmt.Println("CPU:", cpu())

}
