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
		y0 := float64(2*py+0)/(2*height)*(ymax-ymin) + ymin
		y1 := float64(2*py+1)/(2*height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x0 := float64(2*px+0)/(2*width)*(xmax-xmin) + xmin
			x1 := float64(2*px+1)/(2*width)*(xmax-xmin) + xmin
			r00, g00, b00, _ := mandelbrot(complex(x0, y0)).RGBA()
			r01, g01, b01, _ := mandelbrot(complex(x0, y1)).RGBA()
			r10, g10, b10, _ := mandelbrot(complex(x1, y0)).RGBA()
			r11, g11, b11, _ := mandelbrot(complex(x1, y1)).RGBA()
			r := uint8((r00+r01+r10+r11)/4/0x101)
			g := uint8((g00+g01+g10+g11)/4/0x101)
			b := uint8((b00+b01+b10+b11)/4/0x101)
			img.Set(px, py, color.RGBA{r, g, b, 255})
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
