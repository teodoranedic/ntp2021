package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/teodoranedic/ntp2021/parallel"
	"github.com/teodoranedic/ntp2021/sequential"
	"github.com/teodoranedic/ntp2021/types"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var cpu_count = []int{1, 2, 3, 4}
var body_count = []int{200, 150, 100, 50}
var cpu_count_reversed = []int{4, 3, 2, 1}

// fixed 100 bodies, 4000 iterations
var p_pct = 0.99
var s_pct = 0.01

func strong_scaling() {
	speedup := make([]float64, 4)
	amdahl_speedup := make([]float64, 4)
	results := make([][]float64, 30)
	for i := range results {
		results[i] = make([]float64, 4)
	}
	for i := 0; i < 30; i++ {
		t1, _, _ := sequential.Sequential(types.N, types.Planets)
		for j := range cpu_count {
			tn := parallel.Parallel(types.N, types.Planets, cpu_count[j])
			su := float64(t1/time.Millisecond) / float64(tn/time.Millisecond) // speedup
			results[i][j] = su
			if i == 0 {
				amdahl_speedup[j] = 1 / (s_pct + p_pct/float64(cpu_count[j]))
			}
		}
	}

	sum1, sum2, sum3, sum4 := 0.0, 0.0, 0.0, 0.0
	for r := range results {
		sum1 += results[r][0]
		sum2 += results[r][1]
		sum3 += results[r][2]
		sum4 += results[r][3]
	}

	speedup[0] = sum1 / 30
	speedup[1] = sum2 / 30
	speedup[2] = sum3 / 30
	speedup[3] = sum4 / 30

	fmt.Println(speedup)
	fmt.Println(amdahl_speedup)

	plot_results(speedup, amdahl_speedup)

}

func weak_scaling() {
	speedup := make([]float64, 4)
	gustafson_speedup := make([]float64, 4)
	results := make([][]float64, 30)
	for i := range results {
		results[i] = make([]float64, 4)
	}
	for i := 0; i < 30; i++ {
		for j := range cpu_count {
			types.Planets = body_count[j]
			t1, p1, p2 := sequential.Sequential(types.N, types.Planets)
			tn := parallel.Parallel(types.N, types.Planets, cpu_count_reversed[j])
			su := float64(t1/time.Millisecond) / float64(tn/time.Millisecond) // speedup
			results[i][j] = su
			if i == 0 {
				gustafson_speedup[j] = float64(p2/time.Millisecond)/float64(t1/time.Millisecond) + float64(p1/time.Millisecond)/float64(t1/time.Millisecond)*float64(cpu_count_reversed[j])
			}
		}
	}

	sum1, sum2, sum3, sum4 := 0.0, 0.0, 0.0, 0.0
	for r := range results {
		sum1 += results[r][0]
		sum2 += results[r][1]
		sum3 += results[r][2]
		sum4 += results[r][3]
	}

	speedup[0] = sum1 / 30
	speedup[1] = sum2 / 30
	speedup[2] = sum3 / 30
	speedup[3] = sum4 / 30

	for i, j := 0, len(speedup)-1; i < j; i, j = i+1, j-1 {
		speedup[i], speedup[j] = speedup[j], speedup[i]
		gustafson_speedup[i], gustafson_speedup[j] = gustafson_speedup[j], gustafson_speedup[i]
	}

	plot_results(speedup, gustafson_speedup)
}

func plot_results(speedup, max_speedup []float64) {
	speedup_points := make(plotter.XYs, len(speedup))
	max_points := make(plotter.XYs, len(max_speedup))

	for i := range speedup {
		speedup_points[i].X = float64(cpu_count[i])
		speedup_points[i].Y = speedup[i]
		max_points[i].X = float64(cpu_count[i])
		max_points[i].Y = max_speedup[i]
	}

	p := plot.New()

	p.Title.Text = "Weak scaling"
	p.X.Label.Text = "CPU cores"
	p.Y.Label.Text = "Speedup"

	err := plotutil.AddLinePoints(p,
		"Real speedup", speedup_points,
		"Max speedup", max_points,
	)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "resources/weak_scaling.png"); err != nil {
		panic(err)
	}
}

func main() {
	p := parallel.Parallel(types.N, types.Planets, runtime.NumCPU())
	s, p1, p2 := sequential.Sequential(types.N, types.Planets)

	fmt.Println(p)
	fmt.Println(s, p1, p2)

	// strong_scaling()
	// weak_scaling()
}
