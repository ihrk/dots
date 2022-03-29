package dots

import (
	"image"
	"image/color"
)

type rgba64 struct {
	image.Image
}

func (p *rgba64) RGBA64At(x, y int) color.RGBA64 {
	r, g, b, a := p.At(x, y).RGBA()
	return color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)}
}

func toRGBA64(src image.Image) image.RGBA64Image {
	if p, ok := src.(image.RGBA64Image); ok {
		return p
	}

	return &rgba64{src}
}
