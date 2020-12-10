package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -1, -1, +1, +1
		width, height          = 1024 * 4, 1024 * 4
		x_scale                = (xmax - xmin) / width
		y_scale                = (ymax - ymin) / height
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			subPixels := make([]color.Color, 0)
			for i := -1.0; i <= 1; i++ {
				for j := -1.0; j <= 1; j++ {
					z := complex(x+i*x_scale, y+j*y_scale)
					subPixels = append(subPixels, mandelbrot(z))
				}
			}
			img.Set(px, py, avarage(subPixels))
		}
	}
	png.Encode(os.Stdout, img)
}

func avarage(colors []color.Color) color.Color {
	var r, g, b uint32
	n := uint32(len(colors))
	for _, cl := range colors {
		r_, g_, b_, _ := cl.RGBA()
		r += r_ / n
		g += g_ / n
		b += b_ / n
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return genColor(iterations, n)
		}
	}
	return color.Black
}

func genColor(iterations int, n uint8) color.Color {
	// 数が小さければ青、数が大きいほど青 -> 緑 -> 赤にする
	// iterationsを2分割する
	itr := float64(iterations)
	th := float64(itr / 2)
	nf := float64(n)
	var cl color.RGBA

	if nf < th {
		cl = color.RGBA{0, uint8(255 * (nf / th)), uint8(255 * (1 - nf/th)), 255}
	} else {
		x := nf - th
		th := itr - th
		cl = color.RGBA{uint8(255 * (x / th)), uint8(255 * (1 - x/th)), 0, 255}
	}
	return cl
}
