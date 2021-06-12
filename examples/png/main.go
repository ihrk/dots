package main

import (
	"flag"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/ihrk/dots"
)

const defaultURL = "https://static-cdn.jtvnw.net/emoticons/v2/112291/default/dark/2.0"

func main() {
	var (
		url    string
		bg     uint
		th     uint
		dither bool
	)
	flag.StringVar(&url, "url", defaultURL, "")
	flag.BoolVar(&dither, "dither", false, "")
	flag.UintVar(&bg, "bg", dots.DefaultBackground, "")
	flag.UintVar(&th, "th", dots.DefaultThreshold, "")

	flag.Parse()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	png, err := png.Decode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var p *dots.DotImage

	if dither {
		p = dots.ErrDiffDithering(png, dots.Atkinson)
	} else {
		p = dots.Thresholding(png, dots.CodePoint(bg), uint8(th))
	}

	os.Stdout.WriteString(p.String())
}
