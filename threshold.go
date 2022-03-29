package dots

import (
	"image"
)

const (
	DefaultThreshold = 128
)

func Thresholding(src image.Image, dst *Image, threshold8 uint8) {
	r := PixRectToCpRect(src.Bounds()).Intersect(dst.CpRect)

	srcRGBA64 := toRGBA64(src)

	threshold := uint16(threshold8)
	threshold |= threshold << 8

	for row := r.Min.Y; row < r.Max.Y; row++ {
		y0 := row * blockHeight
		for col := r.Min.X; col < r.Max.X; col++ {
			var cp CodePoint

			x0 := col * blockWidth

			for k := 0; k < blockSize; k++ {
				x, y := x0+k%2, y0+k/2

				if gray16At(srcRGBA64, x, y) >= threshold {
					cp |= CodePoint(1 << bitPos[k])
				}
			}

			dst.Cps[dst.CpOffset(col, row)] = cp
		}
	}
}

func gray16(r, g, b uint32) uint32 {
	y := (19595*r + 38470*g + 7471*b + 1<<15) >> 16

	return y
}

func gray16At(src image.RGBA64Image, x, y int) uint16 {
	c := src.RGBA64At(x, y)

	return uint16(gray16(uint32(c.R), uint32(c.G), uint32(c.B)))
}
