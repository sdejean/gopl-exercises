// Server returns lissajous(w) gif
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

var palette = []color.Color{color.White, color.Black}

const (
	DefaultCycles  = 5     // number of complete 'x' oscillator revolutions
	DefaultRes     = 0.001 // angular resolution
	DefaultSize    = 100   // image canvas covers [-size...+size]
	DefaultNframes = 64    // number of animation frames
	DefaultDelay   = 8     // delay between frames in 10ms units
	whiteIndex     = 0     // first color in palette
	blackIndex     = 1     // next color in palette
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	var (
		cycles  = DefaultCycles
		size    = DefaultSize
		nframes = DefaultNframes
		delay   = DefaultDelay
	)
	var err error
	if err = r.ParseForm(); err != nil {
		log.Print(err)
	}
	if r.Form["cycles"] != nil {
		cycles, err = strconv.Atoi(strings.Join(r.Form["cycles"], ""))
		if err != nil {
			log.Print(err)
		}
	}
	if r.Form["size"] != nil {
		size, err = strconv.Atoi(strings.Join(r.Form["size"], ""))
		if err != nil {
			log.Print(err)
		}
	}
	if r.Form["nframes"] != nil {
		nframes, err = strconv.Atoi(strings.Join(r.Form["nframes"], ""))
		if err != nil {
			log.Print(err)
		}
	}
	if r.Form["delay"] != nil {
		delay, err = strconv.Atoi(strings.Join(r.Form["delay"], ""))
		if err != nil {
			log.Print(err)
		}
	}

	lissajous(w, cycles, size, nframes, delay)
}

func lissajous(out io.Writer, cycles int, size int, nframes int, delay int) {

	freq := rand.Float64() * 3.0 // relative frequency of 'y' oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += DefaultRes {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				size+int(x*float64(size)+0.5),
				size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
