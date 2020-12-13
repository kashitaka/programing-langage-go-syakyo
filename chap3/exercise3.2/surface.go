package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zsacle        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if math.IsNaN(ax + ay + bx + by + cx + cy + dx + dy) {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := saddle(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zsacle
	return sx, sy
}

func f(x, y float64) float64 {
	// x , y = 0, 0 のとき
	// r = 0, Sin(0) = 0 となる. 0 / 0 = NaN が返る
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func eggPack(x, y float64) float64 {
	// 卵パック
	// x, yに関してのsin, cosの合成
	return (math.Sin(x) + math.Cos(y)) / 10
}

func mogulRamp(x, y float64) float64 {
	// モーグルこぶ
	// sin,cosをxとyの関数に
	return (math.Sin(x+y) + math.Cos(y+y)) / 10
}

func saddle(x, y float64) float64 {
	// Saddle ... ? 🤔
	return math.Sin(x/xyrange*math.Pi+math.Pi/2) / 2 *
		(math.Cos(y/xyrange*4*math.Pi+math.Pi) + math.Cos(0))
}