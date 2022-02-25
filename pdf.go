package main

import (
	"bytes"
	"fmt"
	"github.com/ledongthuc/pdf"
)

func main() {
	pdf.DebugOn = true
	//content, err := readPdf("D:\\IntelliJ IDEA 2021.3\\project\\talent-management\\server\\uploads\\file\\resume\\44fc5691f0dd51a20c4abee7eb7430eb_20220222170814.pdf") // Read local pdf file
	//content, err := readPdf("D:\\IntelliJ IDEA 2021.3\\project\\talent-management\\server\\uploads\\file\\resume\\acc414b55ab8c2aa3b2d2f0d9a99f4e9_20220222170900.pdf") // Read local pdf file
	//content, err := readPdf("D:\\IntelliJ IDEA 2021.3\\project\\talent-management\\server\\uploads\\file\\resume\\acc414b55ab8c2aa3b2d2f0d9a99f4e9_20220222170900_bb.pdf") // Read local pdf file
	content, err := readPdf("D:\\IntelliJ IDEA 2021.3\\project\\talent-management\\server\\uploads\\file\\resume\\123.pdf") // Read local pdf file
	fmt.Println(err)
	fmt.Println(content)
	return
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}
