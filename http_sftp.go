package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//resp, err := http.Get("http://localhost:8080/api/v2/token")
	//if err == nil {
	//	fmt.Println(resp)
	//}
	//client := &http.Client{}
	//生成要访问的url
	url := "http://10.25.17.2:8080/api/v2/token"
	//提交请求
	request, err := http.NewRequest("GET", url, nil)

	//增加header选项
	request.Header.Add("Authorization", "Basic YWRtaW46MTIzNDU2")
	request.Header.Add("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	client := http.Client{}
	response, _ := client.Do(request)
	defer response.Body.Close()
	fmt.Println(response)
	all, err := ioutil.ReadAll(response.Body)
	fmt.Println(all)
	fmt.Println(string(all))

}
