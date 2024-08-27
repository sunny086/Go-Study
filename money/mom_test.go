package money

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// Transport 表示运输信息的结构体
type Transport struct {
	Date      string // 日期
	Applicant string // 申请用车人
	CarType   string // 车种
	License   string // 车牌号
	FromTo    string // 起止地
	Price     int    // 运费金额（元）
	Remark    string // 备注
}

func TestMom(t *testing.T) {
	file, err := os.Open("mom.txt")
	if err != nil {
		log.Fatalf("打开文件时发生错误: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var transport []Transport
	i := 0
	var currentTransport Transport

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {
			// 循环读取每一行 每读五次 就创建一个新的Transport对象
			switch i {
			case 0:
				// 日期
				currentTransport.Date = line
			case 1:
				// 申请用车人
				currentTransport.Applicant = line
			case 2:
				// 车种和车牌号
				carInfo := strings.Split(line, "，")
				if len(carInfo) > 1 {
					currentTransport.CarType = strings.TrimSpace(carInfo[0])
					// 提取车牌号
					licenseMatch := regexp.MustCompile(`车号(.+)`).FindStringSubmatch(carInfo[1])
					if len(licenseMatch) > 1 {
						currentTransport.License = strings.TrimSpace(licenseMatch[1])
					}
				}
			case 3:
				// 起止地
				currentTransport.FromTo = line
			case 4:
				// 运费金额（元）
				priceMatch := regexp.MustCompile(`运费(\d+)元`).FindStringSubmatch(line)
				if len(priceMatch) > 1 {
					price, err := strconv.Atoi(priceMatch[1])
					if err == nil {
						currentTransport.Price = price
					}
				}
			}

			i++
			// 每读取五行，就将Transport对象添加到切片中
			if i == 5 {
				transport = append(transport, currentTransport)
				// 重置currentTransport
				currentTransport = Transport{}
				i = 0
			}
		}
	}

	// 检查是否有读取错误
	if err := scanner.Err(); err != nil {
		log.Fatalf("读取文件时发生错误: %v", err)
	}

	// 输出读取的运输信息
	for _, t := range transport {
		fmt.Printf("Transport: %+v\n", t)
	}
}
