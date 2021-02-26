package main

import (
	"os"
	"time"

	"github.com/ihrk/dots"
	"github.com/ihrk/dots/text"
	"github.com/ihrk/dots/text/font8"
)

const (
	timeFormat = "15:04:05.0"
	clearTerm  = "\033[H\033[2J"
)

type clock struct {
	font  text.Font
	rd    *dots.Renderer
	frame *dots.DotImage
}

func (c clock) Run() {
	tick := time.NewTicker(100 * time.Millisecond)

	for stamp := range tick.C {
		c.frame.Clear()

		os.Stdout.WriteString(clearTerm)

		tm := stamp.Format(timeFormat)

		text.DisplayString(c.font, tm, c.frame)

		c.rd.Render(os.Stdout, c.frame)
	}
}

func main() {
	c := clock{
		font:  new(font8.Font8),
		rd:    new(dots.Renderer),
		frame: dots.NewDotImage(3*len([]rune(timeFormat)), 3),
	}

	c.Run()
}
