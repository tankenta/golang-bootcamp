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
			img.Set(px, py, solve(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func solve(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	const threshold = 0.01

	f := func(z complex128) complex128 { return cmplx.Pow(z, 4) - 1 }
	fz := f(z)
	for n := uint8(0); n < iterations; n++ {
		z = z - fz/(4 * cmplx.Pow(z, 3))
		fz = f(z)
		if cmplx.Abs(fz) < threshold {
			cbcr := root2CbCr(z)
			return color.YCbCr{255 - contrast*n, cbcr.Cb, cbcr.Cr}
		}
	}
	return color.Black
}

func root2CbCr(z complex128) color.YCbCr {
	var roots = []complex128{1+0i, 0+1i, -1+0i, 0-1i}
	var palette = []color.YCbCr{
		color.YCbCr{0,   0,   0},
		color.YCbCr{0,   0, 255},
		color.YCbCr{0, 255,   0},
		color.YCbCr{0, 255, 255},
	}

	min := math.Inf(1)
	argmin := 0
	for i, root := range roots {
		distance := cmplx.Abs(z - root)
		if distance < min {
			min = distance
			argmin = i
		}
	}
	return palette[argmin]
}
