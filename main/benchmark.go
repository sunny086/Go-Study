package main

import (
	"bytes"
	"fmt"
	"testing"
)

func BenchmarkStringOperation2(b *testing.B) {
	b.ResetTimer()
	str := ""
	for i := 0; i < b.N; i++ {
		str = fmt.Sprintf("%s%s", str, "golang")
	}
}

func BenchmarkStringOperation1(b *testing.B) {
	b.ResetTimer()
	str := ""
	for i := 0; i < b.N; i++ {
		str += "golang"
	}
}

func BenchmarkStringOperation3(b *testing.B) {
	b.ResetTimer()
	strBuf := bytes.NewBufferString("")
	for i := 0; i < b.N; i++ {
		strBuf.WriteString("golang")
	}
}

func BenchmarkStructOperation1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var demo = map[string]interface{}{}
		demo["Name"] = "Tom"
		demo["Age"] = 30
	}
}

func BenchmarkStructOperation2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var demo struct {
			Name string
			Age  int
		}
		demo.Name = "Tom"
		demo.Age = 30
	}
}
