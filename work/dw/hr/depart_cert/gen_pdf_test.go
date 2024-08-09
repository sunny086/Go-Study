package depart_cert

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func TestGenPdf1(t *testing.T) {
	// 创建一个新的数据实例
	article := Article1{
		Title:   "this is a title",
		Content: "content",
	}

	// 定义 HTML 模板
	htmlTemplate := `
<!DOCTYPE html>
<html>
<head>
	<title>{{.Title}}</title>
</head>
<body>
	<h1>{{.Title}}</h1>
	<p>{{.Content}}</p>
</body>
</html>
`

	// 创建一个新的模板，并解析模板内容
	tmpl, err := template.New("article").Parse(htmlTemplate)
	if err != nil {
		fmt.Printf("无法解析 HTML 模板：%v\n", err)
		return
	}

	// 创建一个用于写入 HTML 文件的文件
	htmlFile, err := os.Create("example.html")
	if err != nil {
		fmt.Printf("无法创建 HTML 文件：%v\n", err)
		return
	}
	defer htmlFile.Close()

	// 使用数据渲染模板，并将结果写入 HTML 文件
	err = tmpl.Execute(htmlFile, article)
	if err != nil {
		fmt.Printf("无法生成 HTML 文件：%v\n", err)
		return
	}

	// 使用 wkhtmltopdf 将 HTML 转换为 PDF
	pdfFile := "example.pdf"
	err = runCommand("wkhtmltopdf", "example.html", pdfFile)
	if err != nil {
		fmt.Printf("HTML 转换为 PDF 失败：%v\n", err)
		return
	}

	fmt.Println("PDF 文件生成成功！")

}

// 定义数据结构
type Article1 struct {
	Title   string
	Content string
}

// runCommand 执行命令并返回输出结果
func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	err := cmd.Run()
	return err
}

func TestGenPdf2(t *testing.T) {

	type Employee struct {
		EmployeeName            string `json:"employee_name,omitempty"`
		IDNumber                string `json:"id_number,omitempty"`
		StartDate               string `json:"start_date,omitempty"`
		EndDate                 string `json:"end_date,omitempty"`
		Company                 string `json:"company,omitempty"`
		Position                string `json:"position,omitempty"`
		TerminationDate         string `json:"termination_date,omitempty"`
		CompanySignature        string `json:"company_signature,omitempty"`
		SignatureDate           string `json:"signature_date,omitempty"`
		EmployeeNameReceipt     string `json:"employee_name_receipt,omitempty"`
		IDNumberReceipt         string `json:"id_number_receipt,omitempty"`
		StartDateReceipt        string `json:"start_date_receipt,omitempty"`
		EndDateReceipt          string `json:"end_date_receipt,omitempty"`
		CompanyReceipt          string `json:"company_receipt,omitempty"`
		PositionReceipt         string `json:"position_receipt,omitempty"`
		TerminationDateReceipt  string `json:"termination_date_receipt,omitempty"`
		CompanyReceipt2         string `json:"company_receipt_2,omitempty"`
		SignatureDateReceipt    string `json:"signature_date_receipt,omitempty"`
		RecipientSignature      string `json:"recipient_signature,omitempty"`
		RecipientSignatureDate  string `json:"recipient_signature_date,omitempty"`
		PerformCompeteAgreement string `json:"perform_compete_agreement,omitempty"`
	}

	employee := Employee{
		EmployeeName:            "lisi",
		IDNumber:                "123456789012345678",
		StartDate:               "2020年1月1日",
		EndDate:                 "2022年12月31日",
		Company:                 "某公司",
		Position:                "软件工程师",
		TerminationDate:         "2023年1月1日",
		CompanySignature:        "某公司",
		SignatureDate:           "2023年1月1日",
		EmployeeNameReceipt:     "张三",
		IDNumberReceipt:         "123456789012345678",
		StartDateReceipt:        "2020年1月1日",
		EndDateReceipt:          "2022年12月31日",
		CompanyReceipt:          "某公司",
		PositionReceipt:         "软件工程师",
		TerminationDateReceipt:  "2023年1月1日",
		CompanyReceipt2:         "某公司",
		SignatureDateReceipt:    "2023年1月1日",
		RecipientSignature:      "李四",
		RecipientSignatureDate:  "2023年1月1日",
		PerformCompeteAgreement: "需要"}

	tmpl, err := template.ParseFiles("b.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, employee)
	if err != nil {
		panic(err)
	}

	htmlFile, err := os.Create("example.html")
	if err != nil {
		fmt.Printf("无法创建 HTML 文件：%v\n", err)
		return
	}

	// 使用数据渲染模板，并将结果写入 HTML 文件
	err = tmpl.Execute(htmlFile, employee)
	if err != nil {
		fmt.Printf("无法生成 HTML 文件：%v\n", err)
		return
	}
	htmlFile.Close()
	// 使用 wkhtmltopdf 将 HTML 转换为 PDF
	pdfFile := "example.pdf"
	err = runCommand("wkhtmltopdf", "example.html", pdfFile)
	if err != nil {
		fmt.Printf("HTML 转换为 PDF 失败：%v\n", err)
		return
	}
}

func TestBAse64(t *testing.T) {
	// logo.png转base64
	logoBytes, _ := ioutil.ReadFile("logo.png")
	toString := base64.StdEncoding.EncodeToString(logoBytes)
	fmt.Println(toString)
}
