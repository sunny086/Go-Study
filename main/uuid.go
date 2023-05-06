package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func main() {
	s := uuid.NewV4().String()
	fmt.Println(s)
}
