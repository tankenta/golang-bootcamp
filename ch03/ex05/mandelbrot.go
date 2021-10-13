package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax	= -2, -2, +2, +2
		width, height			= 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			Cb, Cr := iter2CbCr(n, 10)
			return color.YCbCr{255 - contrast*n, Cb, Cr}
		}
	}
	return color.Black
}

func iter2CbCr(iters, itersPerCycle uint8) (uint8, uint8) {
	unitRad := 2 * math.Pi / float64(itersPerCycle)
	Cb := uint8(255 * math.Cos(float64(iters) * unitRad))
	Cr := uint8(255 * math.Sin(float64(iters) * unitRad))
	return Cb, Cr
}
