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
		th   uint
		mode string
	)

	flag.StringVar(&url, "url", defaultURL, "")
	flag.StringVar(&mode, "mode", "dither", "")
	flag.UintVar(&th, "th", dots.DefaultThreshold, "")

	flag.Parse()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	src, err := png.Decode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	p := dots.NewImagePix(src.Bounds())

	switch mode {
	case "dither":
		dots.ErrorDiffusionDithering(src, p, dots.Atkinson)
	case "threshold":
		dots.Thresholding(src, p, uint8(th))
	case "ordered":
		dots.OrderedDithering(src, p, dots.GenerateThresholdMap(1))
	}

	os.Stdout.WriteString(p.String())
}
