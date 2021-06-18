package dots

import (
	"image"
	"image/color"
)

const (
	threshold16 = 1 << 15
	white16     = 0xffff
)

type Ditherer struct {
	buf []int32
}

func (d *Ditherer) checkBuf(n int) {
	if n <= len(d.buf) {
		d.resetBuf(n)
		return
	}
	d.buf = make([]int32, len(d.buf)+n)
}

func (d *Ditherer) resetBuf(n int) {
	for i := 0; i < n; i++ {
		d.buf[i] = 0
	}
}

func (d *Ditherer) Dither(src image.Image, dst *DotImage, k DiffusionKernel) {
	rect := src.Bounds().Intersect(dst.Bounds())

	dx, dy := rect.Dx(), rect.Dy()

	d.checkBuf(dx * dy)

	for y := 0; y < dy; y++ {
		py := rect.Min.Y + y
		for x := 0; x < dx; x++ {
			px := rect.Min.X + x
			ix := y*dx + x

			old := getGrayScale(src, px, py) + d.buf[ix]
			var qErr int32
			d.buf[ix], qErr = convert(old)

			for i, diff := range k.Rows[0] {
				if px+i+1 >= rect.Max.X {
					break
				}

				d.buf[ix+i+1] += qErr * diff / k.Base
			}

			for j := 1; j < len(k.Rows); j++ {
				if py+j >= rect.Max.Y {
					break
				}
				for i, diff := range k.Rows[j] {
					g := i - len(k.Rows[j])/2
					if px+g < rect.Min.X {
						continue
					} else if px+g >= rect.Max.X {
						break
					}
					d.buf[ix+j*dx+g] += qErr * diff / k.Base
				}
			}
		}
	}

	cpDx, cpDy := dx/blockWidth, dy/blockHeight

	for i := 0; i < cpDy; i++ {
		for j := 0; j < cpDx; j++ {
			var cp CodePoint
			cpIx := i*cpDx + j

			x0 := rect.Min.X + j*blockWidth
			y0 := rect.Min.Y + i*blockHeight

			for k := 0; k < blockSize; k++ {
				x, y := x0+k%2, y0+k/2
				ix := y*dx + x
				mask := CodePoint(1 << bitPos[k])
				if d.buf[ix] > 0 {
					cp |= mask
				} else {
					cp &= ^mask
				}
			}

			dst.Cps[cpIx] = cp
		}
	}
}

func ErrDiffDithering(src image.Image, k DiffusionKernel) *DotImage {
	w := src.Bounds().Dx() / blockWidth
	h := src.Bounds().Dy() / blockHeight
	p := NewImage(image.Rect(0, 0, w, h))

	var d Ditherer

	d.Dither(src, p, k)

	return p
}

func getGrayScale(img image.Image, x, y int) int32 {
	return int32(color.Gray16Model.Convert(img.At(x, y)).(color.Gray16).Y)
}

func convert(old int32) (new, quantErr int32) {
	if old >= threshold16 {
		new = white16
	}
	quantErr = old - new
	return
}

type DiffusionKernel struct {
	Base int32
	Rows [][]int32
}

var FloydSteinberg = DiffusionKernel{
	Base: 16,
	Rows: [][]int32{
		{7},
		{3, 5, 1},
	},
}

var JarvisJudiceNinke = DiffusionKernel{
	Base: 48,
	Rows: [][]int32{
		{7, 5},
		{3, 5, 7, 5, 3},
		{1, 3, 5, 3, 1},
	},
}

var Atkinson = DiffusionKernel{
	Base: 8,
	Rows: [][]int32{
		{1, 1},
		{1, 1, 1},
		{1},
	},
}

var Burkes = DiffusionKernel{
	Base: 32,
	Rows: [][]int32{
		{8, 4},
		{2, 4, 8, 4, 2},
	},
}

var Sierra = DiffusionKernel{
	Base: 32,
	Rows: [][]int32{
		{5, 3},
		{2, 4, 5, 4, 2},
		{2, 3, 2},
	},
}

var SierraLite = DiffusionKernel{
	Base: 4,
	Rows: [][]int32{
		{2},
		{1, 1},
	},
}
