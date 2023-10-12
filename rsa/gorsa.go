package rsa

import (
	"Go-Study/rsa/gorsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

// RSA公钥私钥产生
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
	encodePrivate := base64.StdEncoding.EncodeToString(block.Bytes)

	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	encodePublic := base64.StdEncoding.EncodeToString(block.Bytes)
	return encodePrivate, encodePublic, nil
}

// 公钥加密
func PublicEncrypt(data, publicKey string) (string, error) {

	grsa := gorsa.RSASecurity{}
	grsa.SetPublicKey(publicKey)

	rsadata, err := grsa.PubKeyENCTYPT([]byte(data))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(rsadata), nil
}

// 私钥加密
func PriKeyEncrypt(data, privateKey string) (string, error) {

	grsa := gorsa.RSASecurity{}
	grsa.SetPrivateKey(privateKey)

	rsadata, err := grsa.PriKeyENCTYPT([]byte(data))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(rsadata), nil
}

// 公钥解密
func PublicDecrypt(data, publicKey string) (string, error) {

	databs, _ := base64.StdEncoding.DecodeString(data)

	grsa := gorsa.RSASecurity{}
	grsa.SetPublicKey(publicKey)

	rsadata, err := grsa.PubKeyDECRYPT([]byte(databs))
	if err != nil {
		return "", err
	}

	return string(rsadata), nil

}

// 私钥解密
func PriKeyDecrypt(data, privateKey string) (string, error) {

	databs, _ := base64.StdEncoding.DecodeString(data)

	grsa := gorsa.RSASecurity{}
	grsa.SetPrivateKey(privateKey)

	rsadata, err := grsa.PriKeyDECRYPT([]byte(databs))
	if err != nil {
		return "", err
	}

	return string(rsadata), nil
}
