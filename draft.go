package dots

import (
	"image"
	"image/color"
)

const (
	max16            = 0xffff
	defaultThreshold = 128
)

type DotImage struct {
	Dots   []bool
	Stride int
	Rect   image.Rectangle
}

func (p *DotImage) Bounds() image.Rectangle {
	return p.Rect
}

func (p *DotImage) ColorModel() color.Model {
	return DotModel
}

func (p *DotImage) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return color.NRGBA64{}
	}

	i := p.PixOffset(x, y)
	return DotC{p.Dots[i]}
}

func (p *DotImage) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*1
}

type DotC struct {
	isActive bool
}

func (c DotC) RGBA() (r, g, b, a uint32) {
	if c.isActive {
		return max16, max16, max16, max16
	}
	return 0, 0, 0, max16
}

var DotModel color.Model = color.ModelFunc(dotModel)

func dotModel(c color.Color) color.Color {
	if _, ok := c.(DotC); ok {
		return c
	}

	r, g, b, a := c.RGBA()
	return DotC{r == max16 && g == max16 && b == max16 && a == max16}
}
