package aes

import (
	"testing"
	"time"
)

var char = "rTjdHPqMw3nJ97Mr"

func TestAesEncrypt(t *testing.T) {
	encrypt, _ := Base64AESECBEncrypt("110010198010102222", char)
	t.Log(encrypt)
}

func TestAESCBCDecrypt(t *testing.T) {
	decrypt, err := Base64AESECBDecrypt("VjEOPyPCJlIxwWlXS0ZWlTPHUPuoVkfMEJ5q9qBaPrY=", char)
	if err != nil {
		t.Error(err)
	}
	t.Log(decrypt)

	t.Log(time.Now().Year())
	t.Log(time.Now().Month())
	t.Log(time.Now().Day())
	t.Log(time.Now())

	parse, _ := time.Parse("2006-01-02", "2020-12-01")

	if parse.Day() == 1 {
		t.Log("1号")
	} else {
		t.Log("不是12号")
	}

}
