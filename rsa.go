package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

//创建公钥与私钥
func GenRsaKey(bits int) (string, string, error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	w := bytes.NewBuffer([]byte(nil)) //bufio.NewWriter()
	err = pem.Encode(w, block)
	if err != nil {
		return "", "", err
	}
	prikey := string(w.Bytes())
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	w2 := bytes.NewBuffer(nil) //bufio.NewWriter()
	err = pem.Encode(w2, block)
	if err != nil {
		return "", "", err
	}
	pubkey := string(w2.Bytes())
	return pubkey, prikey, nil
}

// 加密
func RsaEncrypt(origData, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func main() {

	//key, s, err := GenRsaKey(128)
	//fmt.Println(key, s, err)
	open, err := os.Open("public_key.pem")
	fmt.Println(err)
	defer open.Close()
	pub, err := ioutil.ReadAll(open)
	fmt.Println(err)
	fmt.Println(pub)
	fmt.Println(string(pub))
	fmt.Println("=============加密=============")
	encrypt, err := RsaEncrypt([]byte("hello"), pub)
	fmt.Println(err)
	fmt.Println(encrypt)
	fmt.Println(string(encrypt))

}
