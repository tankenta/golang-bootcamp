package main

import (
	"log"
	"net/http"
	"sync"
	"strconv"
	"io"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
)

var mu sync.Mutex

var params = map[string]float64{
	"scale": 1.0,
	"xtrans": 0.0,
	"ytrans": 0.0,
}
const (
	xmin, ymin, xmax, ymax	= -2, -2, +2, +2
	width, height			= 1024, 1024
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	mu.Lock()
	for k, v := range r.Form {
		var err error
		params[k], err = strconv.ParseFloat(v[0], 64)
		if err != nil {
			log.Print(err)
		}
	}
	mu.Unlock()
	fractal(w)
}

func fractal(out io.Writer) {
	mu.Lock()
	scale	:= params["scale"]
	xtrans	:= params["xtrans"]
	ytrans	:= params["ytrans"]
	mu.Unlock()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := (float64(py)/height*(ymax-ymin) + ymin + ytrans)/scale
		for px := 0; px < width; px++ {
			x := (float64(px)/width*(xmax-xmin) + xmin + xtrans)/scale
			z := complex(x, y)
			img.Set(px, py, solve(z))
		}
	}
	png.Encode(out, img)
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
