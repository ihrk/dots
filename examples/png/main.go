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
	var url string

	flag.StringVar(&url, "url", defaultURL, "")

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	opts := dots.NewOpts(149, 128)

	p, err := dots.RenderPNG(resp.Body, opts)
	if err != nil {
		log.Println(err)
		return
	}

	os.Stdout.WriteString(p.String())
}
