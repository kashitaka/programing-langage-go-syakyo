package main

import (
	"bytes"
	"fmt"
	"math"
	"sync"
	"time"
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
	start := time.Now()
	var wg sync.WaitGroup
	ch := make(chan string)
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go draw(ch)
	}
	go func() {
		for {
			_, ok := <-ch
			if ok {
				wg.Done()
			}
		}
	}()

	// 待ち受け
	wg.Wait()
	close(ch)
	fmt.Println("done")
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}

func draw(ch chan<- string) {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(&buf, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(&buf, "</svg>")
	ch <- buf.String()
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zsacle
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
