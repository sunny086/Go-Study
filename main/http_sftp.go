package main

import (
	"GoTest/sftp/insert/req"
	"GoTest/sftp/insert/res"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//FtpAuth()
	FtpUserAuthInsert()
}

func FtpUserAuthInsert() {
	var ftpUserInsertReq req.FtpUserInsertReq
	fmt.Println(ftpUserInsertReq)
	ftpUserInsertReq.Status = 1
	ftpUserInsertReq.Username = "xujs"
	ftpUserInsertReq.Password = "123456"
	ftpUserInsertReq.HomeDir = "/srv/sftpgo/data"
	ftpUserInsertReq.VirtualFolders = []req.VirtualFolders{
		{
			Name:        uuid.New().String(),
			MappedPath:  "/data/resource_library/share",
			Users:       []string{"xujs"},
			Filesystem:  req.Filesystem{Provider: 0},
			VirtualPath: "/共享目录",
		},
		{
			Name:        uuid.New().String(),
			MappedPath:  "/data/resource_library/xujs",
			Users:       []string{"xujs"},
			Filesystem:  req.Filesystem{Provider: 0},
			VirtualPath: "/个人目录",
		},
	}
	//map
	permission := make(map[string]interface{})
	permission["/"] = []string{"list"}
	permission["/共享目录"] = []string{"*"}
	permission["/个人目录"] = []string{"*"}

	ftpUserInsertReq.Permissions = permission
	ftpUserInsertReq.Filters.AllowedIP = []string{"0.0.0.0/0"}
	jsonStr, _ := json.Marshal(ftpUserInsertReq)
	fmt.Println(string(jsonStr))
	ftpUserInsertRes, err := FtpUserInsert(string(jsonStr))
	fmt.Println(ftpUserInsertRes)
	fmt.Println(err)
}

type FtpToken struct {
	AccessToken string `json:"access_token"` //令牌
	ExpiresAt   string `json:"expires_at"`   // 过期时间
}

func FtpAuth() (ftpToken FtpToken, err error) {
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
	//转成结构体
	err = json.Unmarshal(all, &ftpToken)
	fmt.Println(ftpToken.AccessToken)
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyIyZmFfcHJvdG9jb2xzIjpudWxsLCIyZmFfcmVxdWlyZWQiOmZhbHNlLCJhdWQiOlsiQVBJIiwiMTAuMjUuMTcuMjMwIl0sImV4cCI6MTY1ODkwNjEyOCwianRpIjoiY2JnZTJvN2l0dG1maTEwa3Q2bTAiLCJuYmYiOjE2NTg5MDQ4OTgsInBlcm1pc3Npb25zIjpbIioiXSwic3ViIjoiSTZpQytyd0lZVW9JcVoyMGFJVFo5VmNlbHpYcUVvZFFpNHpNdjF4cUhRST0iLCJ1c2VybmFtZSI6ImFkbWluIn0.AEHZd_X-0dYVzzXwBH3H_g9Y9JyYNDVYTpqALynYwzA
	return
}

func FtpUserInsert(params string) (ftpUserInsertRes res.FtpUserInsertRes, err error) {
	ftpAuth, _ := FtpAuth()
	// 1. 创建http客户端实例
	client := &http.Client{}
	// 2. 创建请求实例
	req, err := http.NewRequest("POST", "http://10.25.17.2:8080/api/v2/users", strings.NewReader(params))
	if err != nil {
		panic(err)
	}
	// 3. 设置请求头，可以设置多个
	req.Header.Set("Authorization", "Bearer "+ftpAuth.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	// 4. 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 5. 一次性读取请求到的数据
	all, err := ioutil.ReadAll(resp.Body)
	//转成结构体
	err = json.Unmarshal(all, &ftpUserInsertRes)
	return
}
