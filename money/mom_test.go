package money

import (
	"fmt"
	"os"
	"strings"
)

// Transport 表示运输信息的结构体
type Transport struct {
	Date      string // 日期
	Applicant string // 申请用车人
	CarType   string // 车种
	License   string // 车牌号
	FromTo    string // 起止地
	Price     string // 运费金额（元）
	Remark    string // 备注
}

// 字符串数组
var strArr = []string{
	"夏鑫",
	"王静",
	"金铭",
	"丽娜",
	"李延琼",
	"小鱼",
	"小鱼",
	"韩粟",
	"康道林",
}

func main() {
	// 读取文件内容
	content, err := os.ReadFile("mom.txt")
	if err != nil {
		fmt.Println("读取文件时发生错误:", err)
		return
	}
	//替换文本中的一为-
	contentStr := []byte(strings.ReplaceAll(string(content), "一", "-"))
	//fmt.Println(string(contentStr))

	// 将文本按人名划分
	for _, name := range strArr {
		fmt.Println("------", name, "------")
		splitAndPrint(string(contentStr), name)
	}
}

// splitAndPrint 根据人名划分并打印运输信息
func splitAndPrint(content, name string) {
	// 根据人名划分文本
	parts := strings.Split(content, name)

	// 遍历每个划分部分
	for _, part := range parts[1:] { // 跳过第一个空字符串
		part = strings.TrimSpace(part) // 去除前后空白字符
		// 打印运输信息
		fmt.Println(part)
		fmt.Println()
	}
}
