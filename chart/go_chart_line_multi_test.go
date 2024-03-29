package chart

import (
	"github.com/wcharczuk/go-chart/v2"
	"os"
	"testing"
	"time"
)

func TimeArr() []time.Time {
	var arr []time.Time
	for i := 0; i < 11; i++ {
		arr = append(arr, time.Now().Add(time.Hour*time.Duration(i)))
	}
	return arr
}

func TestMultiLineChart(t *testing.T) {
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:           "time",
			TickPosition:   chart.TickPositionBetweenTicks,
			ValueFormatter: chart.TimeValueFormatterWithFormat("2006-01-02 15:04:05"),
		},

		Series: []chart.Series{
			chart.TimeSeries{
				Name:    "conn",
				Style:   chart.StyleTextDefaults(),
				XValues: TimeArr(),
				YValues: []float64{133, 109, 123, 141, 143, 136, 115, 117, 117, 131, 108},
			},
			chart.TimeSeries{
				Name:    "new",
				Style:   chart.StyleTextDefaults(),
				YAxis:   chart.YAxisSecondary,
				XValues: TimeArr(),
				YValues: []float64{0, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1},
			},
		},
	}
	f, _ := os.Create("multi_line.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
