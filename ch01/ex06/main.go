package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0xaf, 0x52, 0xbf, 0xff},
	color.RGBA{0x35, 0x52, 0xba, 0xff},
	color.RGBA{0xa2, 0xcf, 0x6e, 0xff},
	color.RGBA{0xff, 0xcd, 0x38, 0xff},
	color.RGBA{0xff, 0x78, 0x4e, 0xff},
}

const (
	blackIndex = 0
	foreColors = 5
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles	= 5
		res		= 0.001
		size	= 100
		nframes	= 64
		delay	= 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	t_end := cycles*2*math.Pi
	t_changepoint := t_end / float64(foreColors)
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < t_end; t += res {
			x := math.Sin(t)
			t_y := t*freq + phase
			y := math.Sin(t_y)
			color_index := uint8(math.Floor(t/t_changepoint)) + 1
			img.SetColorIndex(
				size+int(x*size+0.5), size+int(y*size+0.5), color_index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
