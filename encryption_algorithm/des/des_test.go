package des

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"testing"
)

/*
可逆
	对称： 文件加密和解密使用相同的密钥，即加密密钥也可以用作解密密钥
				 AES、DES、3DES、Blowfish、IDEA、RC4、RC5、RC6、HS256
	非对称： 公钥加密，私钥解密
				 RSA、DSA（数字签名用）、ECC（移动设备用）、RS256 (采用 SHA‐256 的 RSA 签名)
不可逆
	 一旦加密就不能反向解密得到密码原文
	 MD5、SHA、HMAC
*/

func TestDes(t *testing.T) {
	key := []byte("2fa6c1e9")
	str := "I love this beautiful world!"
	strEncrypted, err := Encrypt(str, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encrypted:", strEncrypted)
	strDecrypted, err := Decrypt(strEncrypted, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decrypted:", strDecrypted)
	//Output:
	//Encrypted: 5d2333b9fbbe5892379e6bcc25ffd1f3a51b6ffe4dc7af62beb28e1270d5daa1
	//Decrypted: I love this beautiful world!
}

func Encrypt(text string, key []byte) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

func Decrypt(decrypted string, key []byte) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
