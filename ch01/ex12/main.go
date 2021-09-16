package main

import (
	"log"
	"net/http"
	"sync"
	"strconv"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
)

var mu sync.Mutex

const res = 0.001
var params = map[string]int{
	"cycles": 5,
	"size": 100,
	"nframes": 64,
	"delay": 8,
}

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
	lissajous(w)
}

func lissajous(out io.Writer) {
	mu.Lock()
	cycles	:= params["cycles"]
	size	:= params["size"]
	nframes := params["nframes"]
	delay	:= params["delay"]
	mu.Unlock()

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	tEnd := float64(cycles)*2*math.Pi
	tChangepoint := tEnd / float64(foreColors)
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < tEnd; t += res {
			x := math.Sin(t)
			tY := t*freq + phase
			y := math.Sin(tY)
			colorIndex := uint8(math.Floor(t/tChangepoint)) + 1
			img.SetColorIndex(
				size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
