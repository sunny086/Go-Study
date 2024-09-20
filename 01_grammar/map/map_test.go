package _map

import (
	"fmt"
	"sync"
	"testing"
)

func Test01(t *testing.T) {

	m := make(map[int]string)
	m[1] = "a"
	m[2] = "b"
	m[3] = "c"
	s := m[4]
	fmt.Println(s)

}
func Test02(t *testing.T) {
	s := sync.Map{}

	s.Store("1", "123")

	actual, loaded := s.LoadOrStore("1", "a")

	fmt.Println(actual, loaded)

}
