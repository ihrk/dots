package dots

import (
	"image"
	"image/color"
)

var (
	map2x2 = [4]uint8{0, 128, 192, 64}
)

func OrderedDithering(src image.Image) *DotImage {
	srcR := src.Bounds()
	w := srcR.Dx() / blockWidth
	h := srcR.Dy() / blockHeight
	p := NewImage(image.Rect(0, 0, w, h))

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ix := i*w + j

			x0 := srcR.Min.X + j*blockWidth
			y0 := srcR.Min.Y + i*blockHeight

			var cp CodePoint

			for k := 0; k < blockSize; k++ {
				x, y := x0+k%2, y0+k/2

				mask := CodePoint(1 << bitPos[k])
				c := src.At(x, y)
				g := color.GrayModel.Convert(c).(color.Gray).Y
				if g > map2x2[k%4] {
					cp |= mask
				} else {
					cp &= ^mask
				}
			}

			p.Cps[ix] = cp
		}
	}

	return p
}

const (
	threshold16 = 1 << 15
	white16     = 0xffff
)

func ErrDiffDithering(src image.Image) *DotImage {
	srcR := src.Bounds()
	srcW := srcR.Dx()
	srcH := srcR.Dy()
	w := srcR.Dx() / blockWidth
	h := srcR.Dy() / blockHeight
	p := NewImage(image.Rect(0, 0, w, h))

	buf := make([]int32, srcW*srcH)

	for y := srcR.Min.Y; y < srcR.Max.Y; y++ {
		off := (y - srcR.Min.Y) * srcW
		for x := srcR.Min.X; x < srcR.Max.X; x++ {
			ix := off + (x - srcR.Min.X)
			g := color.Gray16Model.Convert(src.At(x, y)).(color.Gray16).Y
			old := int32(g) + buf[ix]
			var new int32
			if old >= threshold16 {
				new = white16
			}
			querr := old - new
			buf[ix] = new
			if x+1 < srcR.Max.X {
				t := ix + 1
				buf[t] = buf[t] + querr*7/16
			}
			if y+1 < srcR.Max.Y {
				t := ix + srcW
				if x-1 >= srcR.Min.X {
					buf[t-1] = buf[t-1] + querr*3/16
				}

				buf[t] = buf[t] + querr*5/16

				if x+1 < srcR.Max.X {
					buf[t+1] = buf[t+1] + querr*1/16
				}
			}
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			var cp CodePoint
			ix := i*w + j

			x0 := srcR.Min.X + j*blockWidth
			y0 := srcR.Min.Y + i*blockHeight

			for k := 0; k < blockSize; k++ {
				x, y := x0+k%2, y0+k/2
				ix := y*srcW + x
				mask := CodePoint(1 << bitPos[k])
				if buf[ix] > 0 {
					cp |= mask
				} else {
					cp &= ^mask
				}
			}

			p.Cps[ix] = cp
		}
	}

	return p
}
