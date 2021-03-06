package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"log"
	"net/http"
	"os"
	"fmt"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", lissajousHandler)
	log.Fatal(http.ListenAndServe(GetPort(), nil))
}

func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
		fmt.Println("$PORT must be set")
	}
	return ":" + port
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	lissajous(w)	
}

func lissajous(out io.Writer) {
	const (
		cycles = 5
		res = 0.001
		size = 500
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
