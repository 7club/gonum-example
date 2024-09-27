package main

import (
	"fmt"
	"gonum.org/v1/gonum/optimize"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"math"
)

func main() {
	points2 := plotter.XYs{
		plotter.XY{
			X: float64(5),
			Y: float64(10),
		},
		plotter.XY{
			X: float64(10),
			Y: float64(8),
		},
		plotter.XY{
			X: float64(20),
			Y: float64(6),
		},
		plotter.XY{
			X: float64(50),
			Y: float64(4),
		},
		plotter.XY{
			X: float64(100),
			Y: float64(2.2),
		},
		plotter.XY{
			X: float64(200),
			Y: float64(1.5),
		},
		plotter.XY{
			X: float64(240),
			Y: float64(1.4),
		},
		plotter.XY{
			X: float64(280),
			Y: float64(1.27),
		},
		plotter.XY{
			X: float64(300),
			Y: float64(1.2),
		},
		plotter.XY{
			X: float64(350),
			Y: float64(1.15),
		},
		plotter.XY{
			X: float64(400),
			Y: float64(1.12),
		},
		plotter.XY{
			X: float64(500),
			Y: float64(1.1),
		},
	}

	result, err := optimize.Minimize(optimize.Problem{
		Func: func(x []float64) float64 {
			if len(x) != 2 {
				panic("illegal x")
			}
			n := x[0]
			p := x[1]
			var sum float64
			for _, point := range points2 {
				y := n * math.Pow(p, point.X)
				sum += math.Abs(y - point.Y)
			}
			return sum
		},
	}, []float64{1, 1}, &optimize.Settings{}, &optimize.NelderMead{})
	if err != nil {
		panic(err)
	}

	fa, fb := result.X[0], result.X[1]
	fmt.Println(fa, fb)
	points3 := plotter.XYs{}
	for i := 5; i <= 500; i += 10 {
		points3 = append(points3, plotter.XY{
			X: float64(i),
			Y: fa * math.Pow(fb, float64(i)), // N(t)=n(p^t)
		})
	}

	plt := plot.New()
	plt.Y.Min, plt.X.Min, plt.Y.Max, plt.X.Max = 0, 0, 100, 100

	if err := plotutil.AddLinePoints(plt,
		"line2", points2,
		"line3", points3,
	); err != nil {
		panic(err)
	}

	if err := plt.Save(5*vg.Inch, 5*vg.Inch, "11-optimize-fit.png"); err != nil {
		panic(err)
	}
}
