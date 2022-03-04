package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	interfaces, _ := net.Interfaces()
	fmt.Println(interfaces)
	s, _ := json.Marshal(interfaces)
	a := string(s)
	fmt.Println(a)
}
