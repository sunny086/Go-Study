package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()

	fmt.Println(interfaces)

	fmt.Println(err)

}
