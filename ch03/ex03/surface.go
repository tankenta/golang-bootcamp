package main

import (
	"fmt"
	"math"
)

const (
	width, height	= 600, 320
	cells			= 100
	xyrange			= 30.0
	xyscale			= width / 2 / xyrange
	zscale			= height * 0.4
	angle			= math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	zMin, zMax := math.Inf(1), math.Inf(-1)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, az, aok := corner(i+1, j)
			_, _, bz, bok := corner(i, j)
			_, _, cz, cok := corner(i, j+1)
			_, _, dz, dok := corner(i+1, j+1)
			if !aok || !bok || !cok || !dok {
				continue
			}
			zMin = math.Min(zMin, az)
			zMin = math.Min(zMin, bz)
			zMin = math.Min(zMin, cz)
			zMin = math.Min(zMin, dz)
			zMax = math.Max(zMax, az)
			zMax = math.Max(zMax, bz)
			zMax = math.Max(zMax, cz)
			zMax = math.Max(zMax, dz)
		}
	}

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, aok := corner(i+1, j)
			bx, by, bz, bok := corner(i, j)
			cx, cy, cz, cok := corner(i, j+1)
			dx, dy, dz, dok := corner(i+1, j+1)
			if !aok || !bok || !cok || !dok {
				continue
			}
			zNormalized := ((az+bz+cz+dz)/4 - zMin) / (zMax - zMin)
			red := uint8(zNormalized * 255)
			blue := uint8((1 - zNormalized) * 255)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='#%02x00%02x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, red, blue)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (sx, sy, z float64, ok bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z = f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, z, false
	}

	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
