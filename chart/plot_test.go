package chart

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"log"
	"testing"
)

func TestPlot(t *testing.T) {
	// mock data
	x := []float64{0, 1, 2, 3, 4, 5}
	y1 := []float64{0, 2, 4, 6, 8, 10}
	y2 := []float64{1, 3, 5, 7, 9, 11}

	// create plot
	p := plot.New()

	// set plot title
	p.Title.Text = "Line Chart"

	// create line series1
	pts1 := make(plotter.XYs, len(x))
	for i := range pts1 {
		pts1[i].X = x[i]
		pts1[i].Y = y1[i]
	}
	series1, err := plotter.NewLine(pts1)
	if err != nil {
		log.Fatalf("Failed to create series1: %v", err)
	}
	series1.LineStyle.Width = vg.Points(2)
	series1.LineStyle.Color = color.RGBA{R: 255, A: 255}

	// create line series2
	pts2 := make(plotter.XYs, len(x))
	for i := range pts2 {
		pts2[i].X = x[i]
		pts2[i].Y = y2[i]
	}
	series2, err := plotter.NewLine(pts2)
	if err != nil {
		log.Fatalf("Failed to create series2: %v", err)
	}
	series2.LineStyle.Width = vg.Points(2)
	series2.LineStyle.Color = color.RGBA{B: 255, A: 255}

	// add series to plot
	p.Add(series1, series2)

	// set plot legend
	p.Legend.Add("Series 1", series1)
	p.Legend.Add("Series 2", series2)
	p.Legend.ThumbnailWidth = vg.Inch * 0.5

	// set plot axis labels
	p.X.Label.Text = "X-Axis"
	p.Y.Label.Text = "Y-Axis"

	// plot
	if err := p.Save(10*vg.Centimeter, 10*vg.Centimeter, "chart.png"); err != nil {
		log.Fatalf("Failed to plot chart: %v", err)
	}
	fmt.Println("Chart saved to chart.png")
}
