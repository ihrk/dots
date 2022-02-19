package dots

import "image"

const (
	DefaultBackground = 0
	DefaultThreshold  = 128
)

func Thresholding(src image.Image, bg CodePoint, th uint8) *DotImage {
	srcR := src.Bounds()
	w := srcR.Dx() / blockWidth
	h := srcR.Dy() / blockHeight
	p := NewImage(image.Rect(0, 0, w, h))

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ix := i*w + j

			x0 := srcR.Min.X + j*blockWidth
			y0 := srcR.Min.Y + i*blockHeight

			cp := bg

			for k := 0; k < blockSize; k++ {
				x, y := x0+k%2, y0+k/2

				mask := CodePoint(1 << bitPos[k])
				c := src.At(x, y)
				_, _, _, alpha := c.RGBA()
				g := uint8(gray16At(src, x, y) >> 8)
				if alpha != 0 {
					if g >= th {
						cp |= mask
					} else {
						cp &= ^mask
					}
				}
			}

			p.Cps[ix] = cp
		}
	}

	return p
}
