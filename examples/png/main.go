package main

import (
	"flag"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/ihrk/dots"
)

const defaultURL = "https://static-cdn.jtvnw.net/emoticons/v2/112291/default/dark/3.0"

func main() {
	var (
		url  string
		bg   uint
		th   uint
		mode string
	)

	flag.StringVar(&url, "url", defaultURL, "")
	flag.StringVar(&mode, "mode", "dither", "")
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

	switch mode {
	case "dither":
		p = dots.ErrDiffDithering(png, dots.Atkinson)
	case "threshold":
		p = dots.Thresholding(png, dots.CodePoint(bg), uint8(th))
	case "ordered":
		p = dots.OrderedDithering(png, dots.GenerateThresholdMap(1))
	}

	os.Stdout.WriteString(p.String())
}
