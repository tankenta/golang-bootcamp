package main

import (
	"log"
	"net/http"
	"sync"
	"strconv"
	"io"
	"fmt"
	"math"
)

var mu sync.Mutex

var params = map[string]int{
	"width": 600,
	"height": 320,
	"cells": 100,
	"red": 0,
	"green": 0,
	"blue": 0,
}
const (
	xyrange = 30.0
	angle	= math.Pi / 6
)
var xyscale = float64(params["width"]) / 2.0 / xyrange
var zscale = float64(params["height"]) * 0.4
var sin30, cos30 = math.Sin(angle), math.Cos(angle)

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
		params[k], err = strconv.Atoi(v[0])
		if err != nil {
			log.Print(err)
		}
	}
	mu.Unlock()
	w.Header().Set("Content-Type", "image/svg+xml")
	writeSVG(w)
}

func writeSVG(out io.Writer) {
	mu.Lock()
	width	:= params["width"]
	height	:= params["height"]
	cells	:= params["cells"]
	red		:= uint8(params["red"])
	green	:= uint8(params["green"])
	blue	:= uint8(params["blue"])
	mu.Unlock()

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: #%02x%02x%02x; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", red, green, blue, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aok := corner(width, height, cells, i+1, j)
			bx, by, bok := corner(width, height, cells, i, j)
			cx, cy, cok := corner(width, height, cells, i, j+1)
			dx, dy, dok := corner(width, height, cells, i+1, j+1)
			if !aok || !bok || !cok || !dok {
				continue
			}
			fmt.Fprintf(out,
				"<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(width, height, cells, i, j int) (sx, sy float64, ok bool) {
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)
	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, false
	}

	sx = float64(width)/2 + (x-y)*cos30*xyscale
	sy = float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
