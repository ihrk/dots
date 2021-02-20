package font8

import (
	"image"

	"github.com/ihrk/dots"
)

func (f *Font8) LoadChar(r rune) {
	if int(r) > len(chars) {
		r = ' '
	}
	f.src = chars[r]
}

func (f *Font8) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(f.src[0]), 3)
}

func (f *Font8) GetCP(x, y int) dots.CodePoint {
	return dots.FromRune(f.src[y][x])
}
