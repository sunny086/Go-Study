package chart

import (
	"bytes"
	"github.com/wcharczuk/go-chart/v2"
	"os"
	"testing"
)

func TestPieChart(t *testing.T) {
	var err error
	// 设置中文字体
	// 创建数据序列
	data := []chart.Value{
		{Value: 50, Label: "苹果"},
		{Value: 25, Label: "Banana"},
		{Value: 25, Label: "Orange"},
	}
	font, err := chart.GetDefaultFont()
	// 创建图表对象
	pie := chart.PieChart{
		Width:  512,
		Height: 256,
		Values: data,
		Font:   font,
	}

	// 生成图表文件
	outputBuf2 := new(bytes.Buffer)
	pie.Render(chart.PNG, outputBuf2)
	//生成图片
	err = os.WriteFile("pie.png", outputBuf2.Bytes(), 0644)
	if err != nil {
		t.Fatal(err)
	}
}
