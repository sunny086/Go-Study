package main

import (
	"flag"
	"fmt"
)

func main() {
	//go build之后可以运行加参数，下面的代码可以获取命令行参数
	var config string
	flag.StringVar(&config, "c", "abc", "choose config file.")
	flag.Parse()
	fmt.Println(config)
}
