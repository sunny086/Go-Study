package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	baseURL := "http://127.0.0.1/api/statistics/purchase_inventory_report?materiel_name=%E7%89%A9%E6%96%991&report_date=[%222024-06-01%22,%222024-06-16%22]&page_no=1&page_size=10"

	//// 设置查询参数
	//queryParams := url.Values{}
	//queryParams.Set("materiel_name", "物料1")                       // 假设直接传入非编码的字符串
	//queryParams.Set("report_date", `["2024-06-01","2024-06-16"]`) // 这里假设API可以接受JSON格式的日期数组作为查询参数（这通常不是标准做法）
	//queryParams.Set("page_no", "1")
	//queryParams.Set("page_size", "10")

	//// 构建完整的URL
	//urlWithQuery := baseURL + "?" + queryParams.Encode()

	// 发起GET请求
	resp, err := http.Get(baseURL)
	if err != nil {
		t.Fatalf("Error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected response status code: %d", resp.StatusCode)
	}

	// 读取响应体（如果需要的话）
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	// 打印响应体或进行其他处理
	fmt.Println(string(body))

}
