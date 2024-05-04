package main

import (
	// "fmt"
	"log"
	"math/rand"
	"sort"

	// "os"
	"strings"

	// "github.com/go-echarts/go-echarts/v2/charts"
	// "github.com/go-echarts/go-echarts/v2/opts"
	"go-hep.org/x/hep/hplot"
	// "golang.org/x/tools/go/analysis/passes/sortslice"
	"gonum.org/v1/plot"

	// "gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// orig problem https://twitter.com/littmath/status/1786479662861877742
func litt_coins_may3() [2]int {
	flips := make([]uint8, 100)
	for i, _ := range flips {
		flips[i] = uint8("HT"[rand.Uint32()%2])
	}
	flips_str := string(flips)

	// i0 := strings.Index(flips_str, strings.Replace("HTHTH HTHTH HTHTH", " ", "", -1))
	// i1 := strings.Index(flips_str, strings.Replace("HTTTH HTTTH HTTTH", " ", "", -1))
	i0 := strings.Index(flips_str, strings.Replace("HTHTH", " ", "", -1))
	i1 := strings.Index(flips_str, strings.Replace("HTTTH", " ", "", -1))

	// fmt.Println(i0, i1)
	return [2]int{i0, i1}
}

func intslices_to_XYer(xs, ys []int) plotter.XYer {
	xs0 := make([]float64, len(xs))
	ys0 := make([]float64, len(ys))
	for i := 0; i < len(xs); i++ {
		xs0[i] = float64(xs[i])
		ys0[i] = float64(ys[i])
	}
	data := hplot.ZipXY(xs0, ys0)
	return data
}

// [lo, hi) exclusive on the top
func gen_range(lo, hi int) []int {
	if lo > hi {
		return nil
	}
	result := make([]int, hi-lo)
	for i := range result {
		result[i] = lo + i
	}
	return result
}

func main() {
	N := 10_000
	xs := make([]int, N)
	ys := make([]int, N)
	for i := 0; i < N; i++ {
		r := litt_coins_may3()
		xs[i] = r[0]
		ys[i] = r[1]
	}

	p := plot.New()
	p.Title.Text = "Coins"
	p.X.Label.Text = "Index"
	p.Y.Label.Text = "[0..100]"

	sort.Ints(xs)
	sort.Ints(ys)
	// ExampleScatters_plotutil(p, xs, gen_range(0, len(xs)))
	// ExampleScatters_plotutil(p, ys, gen_range(0, len(ys)))
	data0 := intslices_to_XYer(xs, gen_range(0, len(xs)))
	data1 := intslices_to_XYer(ys, gen_range(0, len(ys)))

	var err error

	err = plotutil.AddScatters(p, "data 0", data0, "data 1", data1)
	if err != nil {
		log.Fatalf("could not create scatters: %+v", err)
	}

	err = p.Save(20*vg.Centimeter, 10*vg.Centimeter, "scatter.pdf")
	if err != nil {
		log.Fatalf("could not save scatter plot: %+v", err)
	}

	// hplot.ZipXY(xs, ys)
	// chart := charts.NewScatter()
	// chart.AddSeries("HTHTH", scat0)
	// chart.AddSeries("HTTTH", scat1)
	// f, _ := os.Create("littchart.html")
	// chart.Render(f)

}
