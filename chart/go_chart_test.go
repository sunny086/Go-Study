package chart

import (
	"bytes"
	"encoding/base64"
	"github.com/wcharczuk/go-chart"
	"io"
	"os"
	"testing"
	"text/template"
)

type ReportData struct {
	UserName              string                // 当前用户
	SN                    string                // 设备SN号
	ReportTime            string                // 报告生成时间
	ProtocolTypeStatistic ProtocolTypeStatistic // 协议类型统计
	SessionStatistic      SessionStatistic      // 实时会话监控
	ProtocolCurrentFlow   ProtocolCurrentFlow   // 协议实时流量
	ProtocolFlowStatistic ProtocolFlowStatistic // 协议流量统计
}

type ProtocolTypeStatistic struct {
	ProtocolCount      int    //协议数量
	MaxProtocolName    string //最大协议名称
	MaxProtocolPercent string //最大协议百分比
	ImgData            string //图片数据
}

type SessionStatistic struct {
	CharsTimeDelta string  //时间范围
	Concs          []int64 // 并发数
	News           []int64 // 新建数
	ImgData        string  //图片数据
}

type ProtocolCurrentFlow struct {
	ImgData string //图片数据

}

type ProtocolFlowStatistic struct {
	CharsTimeDelta string //时间范围
	ImgData        string //图片数据

}

func TestGoChart(t *testing.T) {
	var reportData = ReportData{
		UserName: "admin",
	}
	templateFilePath := "template.xml"
	open, _ := os.Open(templateFilePath)
	templateContent, err := io.ReadAll(open)
	if err != nil {
		t.Fatal(err)
	}
	temp := template.New("report")
	parse, err := temp.Parse(string(templateContent))
	if err != nil {
		t.Fatal(err)
	}
	outputBuf := new(bytes.Buffer)

	//使用go-chart生成图表
	// 创建数据序列
	data := []chart.Value{
		{Value: 50, Label: "Apple"},
		{Value: 25, Label: "Banana"},
		{Value: 25, Label: "Orange"},
	}

	// 创建图表对象
	pie := chart.PieChart{
		Width:  512,
		Height: 256,
		Values: data,
	}

	// 生成图表文件
	outputBuf2 := new(bytes.Buffer)
	pie.Render(chart.PNG, outputBuf2)
	reportData.ProtocolTypeStatistic.ImgData = base64.StdEncoding.EncodeToString(outputBuf2.Bytes())
	err = parse.Execute(outputBuf, &reportData)
	err = os.WriteFile("report.doc", outputBuf.Bytes(), 0644)
	if err != nil {
		t.Fatal(err)
	}
}

func Test1(t *testing.T) {
	// 创建数据序列
	data := []chart.Value{
		{Value: 50, Label: "Apple"},
		{Value: 25, Label: "Banana"},
		{Value: 25, Label: "Orange"},
	}

	// 创建图表对象
	pie := chart.PieChart{
		Title:  "Fruit Distribution",
		Values: data,
	}

	// 生成图表文件
	outputBuf2 := new(bytes.Buffer)
	pie.Render(chart.PNG, outputBuf2)
	base64.StdEncoding.EncodeToString(outputBuf2.Bytes())
}
