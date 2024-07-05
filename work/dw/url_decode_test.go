package dw

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrlDecode1(t *testing.T) {
	encoded := "%E4%BF%A1%E6%81%AF"

	// 使用QueryUnescape（在Go 1.16之前）或PathUnescape（在Go 1.16及之后）进行解码
	decoded, err := url.QueryUnescape(encoded) // 在Go 1.16及之后，建议使用PathUnescape
	if err != nil {
		fmt.Println("Error decoding URL:", err)
		return
	}

	fmt.Println("Decoded string:", decoded)

	unescape, err := url.QueryUnescape(decoded)
	if err != nil {
		fmt.Println("Error decoding URL:", err)
		return
	}
	fmt.Println("Unescaped string:", unescape)

}

func TestUrlDecode2(t *testing.T) {
	encoded := "%E4%BF%A1%E6%81%AF"

	// 使用QueryUnescape（在Go 1.16之前）或PathUnescape（在Go 1.16及之后）进行解码
	decoded, err := url.PathUnescape(encoded) // 在Go 1.16及之后，建议使用 PathUnescape
	if err != nil {
		fmt.Println("Error decoding URL:", err)
		return
	}

	fmt.Println("Decoded string:", decoded)

	unescape, err := url.PathUnescape(decoded)
	if err != nil {
		fmt.Println("Error decoding URL:", err)
		return
	}
	fmt.Println("Unescaped string:", unescape)

}
