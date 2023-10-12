package rsa

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestRsaEncrypt(t *testing.T) {
	//key, s, err := GenRsaKey(1024)
	//fmt.Println(key)
	//fmt.Println(s)
	//fmt.Println(err)
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

func TestReaGenerateAndEncrypt(t *testing.T) {
	publicKey, privateKey, err := GenRsaKey(1024)
	if err != nil {
		t.Error(err)
	}
	t.Log(publicKey)
	t.Log(privateKey)
	encrypt, err := RsaEncrypt([]byte("hello"), []byte(publicKey))
	if err != nil {
		t.Error(err)
	}
	t.Log(string(encrypt))

}

func TestRsa(t *testing.T) {
	secret := "jiFhl43b4B1Kh05idTGji2UcTjSdVFVe2ekZH9Uo21ot+df/YGIH+TjqOCsiyE9P9hc0+u20mkYVX1tyIT78OjzLYNk20xaxPlM0SWhDXrVP14q2l2OF7YJl5K6SiO89NtdN1vTK/Ch4JGJyYfeL6K0ZgHIoqm9VYHOOXr6ZzjI="
	pubKey := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDQpLMJqlrpj+ftN+anPTmktgiLhMguQt3QCvO1MC4ADKUeKYKaHL6zGDliyQnMd4drX/alJ9oEoK5pM10hh4tKhYhMgXu/YoI7xcgl50Viwoa8WbzBqknWW/3LpKVvfHOPGIbGPthMcV63rLb3Aao2TFinC0/rqv8lA9H3t2VrawIDAQAB"
	decrypt, err := PublicDecrypt(secret, pubKey)
	if err != nil {
		t.Error(err)
	}
	t.Log(decrypt)

	//获取decrypt的hash值
	h := sha256.New()
	h.Write([]byte(decrypt))
	hash := h.Sum(nil)
	t.Log(hex.EncodeToString(hash))
}
