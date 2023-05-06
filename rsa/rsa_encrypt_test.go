package rsa

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestRsaEncrypt(t *testing.T) {
	//key, s, err := GenRsaKey(128)
	//fmt.Println(key, s, err)
	open, err := os.Open("./public_key.pem")
	defer open.Close()
	if err != nil {
		t.Error(err)
	}
	pub, err := io.ReadAll(open)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(pub))
	fmt.Println("=============加密=============")
	encrypt, err := RsaEncrypt([]byte("hello"), pub)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(encrypt))
}
