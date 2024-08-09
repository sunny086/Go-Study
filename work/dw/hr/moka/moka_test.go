package moka

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestMokaCreateRecruit(t *testing.T) {

}

func CreateRecruit() error {
	url := "https://api-staging-3.mokahr.com/api-platform/v1/headcount/?currentHireMode=1"
	apiKey := "mDQGQOWGBKQc3XSyxHp5WTmgmWSUBwYs"
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:", apiKey)))

	paramStr := `{"number":"DW-ZPSQ-20240801-00004-1","jobName":"数字化组件研发团队","type":"planned","needNumber":5,"departmentCode":"code1","commitment":"fulltime","description":"工作职责\u0026任职要求","status":"draft","startDate":"2024-08-05T13:30:11.812Z","completeDate":"2024-08-10T00:00:00.000Z","minSalary":10,"maxSalary":20,"education":"本科"}`
	// 转map

	jsonData, err := json.Marshal(paramStr)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	return nil
}

func TestPost2(t *testing.T) {
	url := "https://api-staging-3.mokahr.com/api-platform/v1/headcount/?currentHireMode=1"
	method := "POST"

	payload := strings.NewReader(`{
    "number": "DW-ZPSQ-20240801-00004-1",
    "jobName": "数字化组件研发团队",
    "type": "planned",
    "needNumber": 5,
    "departmentCode": "code1",
    "commitment": "fulltime",
    "description": "工作职责\u0026任职要求",
    "status": "draft",
    "startDate": "2024-08-05T13:30:11.812Z",
    "completeDate": "2024-08-10T00:00:00.000Z",
    "minSalary": 10,
    "maxSalary": 20,
    "education": "本科"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Basic bURRR1FPV0dCS1FjM1hTeXhIcDVXVG1nbVdTVUJ3WXM6")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
