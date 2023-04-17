package main

import (
	_ "embed"
	"fmt"
)

//go:embed config.yaml
var configData []byte

func main() {
	fmt.Println(string(configData))
}

type CaptchaConfig struct {
	KeyLong   int  `yaml:"key-long"`
	ImgWidth  int  `yaml:"img-width"`
	ImgHeight int  `yaml:"img-height"`
	Enable    bool `yaml:"enable"`
}
