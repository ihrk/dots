package text

import (
	"image"

	"github.com/ihrk/dots"
)

type Font interface {
	GetCP(int, int) dots.CodePoint
	LoadChar(rune)
	Bounds() image.Rectangle
}

func DisplayString(f Font, s string, p *dots.DotImage) {
	picR := p.Rect
	offW := 0
	for _, r := range s {
		f.LoadChar(r)
		fontR := f.Bounds()
		fontR = fontR.Add(image.Pt(offW, 0))

		r := picR.Intersect(fontR)
		if r.Empty() {
			break
		}

		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				p.Cps[p.CpOffset(x, y)] = f.GetCP(x-r.Min.X, y-r.Min.Y)
			}
		}

		offW += r.Dx()
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
