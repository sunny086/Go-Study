package yaml

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

func TestReadYaml(t *testing.T) {
	// 读取 YAML 文件
	data, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}

	// 解析 YAML 文件
	var config CaptchaConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	// 打印配置信息
	fmt.Printf("key-long: %d\n", config.Captcha.KeyLong)
	fmt.Printf("img-width: %d\n", config.Captcha.ImgWidth)
	fmt.Printf("img-height: %d\n", config.Captcha.ImgHeight)
	fmt.Printf("enable: %t\n", config.Captcha.Enable)
}

type CaptchaConfig struct {
	Captcha Captcha `yaml:"captcha"`
}

type Captcha struct {
	KeyLong   int  `yaml:"key-long"`
	ImgWidth  int  `yaml:"img-width"`
	ImgHeight int  `yaml:"img-height"`
	Enable    bool `yaml:"enable"`
}
