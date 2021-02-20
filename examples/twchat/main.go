package main

import (
	"flag"
	"fmt"

	"github.com/ihrk/dots"
	"github.com/ihrk/dots/text"
	"github.com/ihrk/dots/text/font8"
)

func main() {
	var (
		msg       string
		isFlipped bool
	)

	flag.StringVar(&msg, "msg", "defaultval", "")
	flag.BoolVar(&isFlipped, "flip", false, "")

	flag.Parse()

	img := dots.NewImage(30, 3)
	text.DisplayString(new(font8.Font8), msg, img)
	if isFlipped {
		img.FlipBits()
	}

	fmt.Print(img)
}
