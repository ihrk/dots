package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/ihrk/dots"
)

const defaultURL = "https://static-cdn.jtvnw.net/emoticons/v2/112291/default/dark/2.0"

func main() {
	var (
		url string
		bg  uint
		th  uint
	)

	flag.StringVar(&url, "url", defaultURL, "")
	flag.UintVar(&bg, "bg", dots.DefBg, "")
	flag.UintVar(&th, "th", dots.DefTh, "")

	flag.Parse()

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	opts := dots.NewOpts(dots.CodePoint(bg), uint8(th))

	p, err := dots.RenderPNG(resp.Body, opts)
	if err != nil {
		log.Println(err)
		return
	}

	os.Stdout.WriteString(p.String())
}
