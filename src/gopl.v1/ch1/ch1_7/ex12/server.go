package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.RGBA{0x0F, 0xDF, 0x0C, 0xFF}, color.RGBA{0x1C, 0x1E, 0xB3, 0xFF}}

func main() {
	http.HandleFunc("/lissajous", lissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func lissajous(out http.ResponseWriter, r *http.Request) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}

	cycles := r.Form["cycles"][0]
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i, idx := 0, 1; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		cycles, err := strconv.Atoi(cycles)
		if err != nil {
			fmt.Fprintf(out, "Change params %s", cycles)
		}
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(idx))
		}
		// 遍历后几种颜色
		idx++
		if idx > len(palette) {
			idx = 1
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
