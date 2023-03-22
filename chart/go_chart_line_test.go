package chart

import (
	"github.com/wcharczuk/go-chart"
	"os"
	"testing"
)

func TestLineChart1(t *testing.T) {
	xValues := []float64{1, 2, 3, 4, 5}
	yValues := []float64{1, 3, 2, 4, 5}

	graph := chart.Chart{
		// 设置图表的标题
		Title: "My Line Graph",
		Series: []chart.Series{
			// 添加折线数据
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: yValues,
			},
		},
	}

	// 将图表保存到文件
	f, _ := os.Create("line-chart1.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

func TestLineChart2(t *testing.T) {
	xValues := []float64{1, 2, 3, 4, 5}
	yValues := []float64{1, 3, 2, 4, 5}

	graph := chart.Chart{
		Title: "My Line Graph",
		XAxis: chart.XAxis{
			Name:      "X Axis",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
			Range: &chart.ContinuousRange{
				Min: 0,
				Max: 6,
			},
		},
		YAxis: chart.YAxis{
			Name:      "Y Axis",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
			Range: &chart.ContinuousRange{
				Min: 0,
				Max: 6,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: yValues,
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
			},
		},
	}

	f, _ := os.Create("line-chart2.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

func TestLineChart3(t *testing.T) {
	xValues := []float64{1, 2, 3, 4, 5}
	yValues := []float64{1, 3, 2, 4, 5}

	graph := chart.Chart{
		Title: "My Line Graph",
		XAxis: chart.XAxis{
			Name:      "X Axis",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
			Range: &chart.ContinuousRange{
				Min: 0,
				Max: 6,
			},
		},

		YAxis: chart.YAxis{
			Name:      "Y Axis",
			NameStyle: chart.StyleShow(),
			Style: chart.Style{
				Show: true,
			},
			Range: &chart.ContinuousRange{
				Min: 0,
				Max: 6,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: yValues,
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(0),
				},
			},
		},
	}

	f, _ := os.Create("line-chart3.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

func TestLineChart4(t *testing.T) {
}
