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
	pixR := src.Bounds().Intersect(dst.Bounds())

	pixDx := pixR.Dx()
	pixDy := pixR.Dy()

	d.checkBuf(pixDx * pixDy)

	for pixRow := 0; pixRow < pixDy; pixRow++ {
		pixY := pixR.Min.Y + pixRow
		for pixCol := 0; pixCol < pixDx; pixCol++ {
			pixX := pixR.Min.X + pixCol
			bufIx := pixRow*pixDx + pixCol

			oldValue := int32(gray16At(src, pixX, pixY)) + d.buf[bufIx]

			var newValue int32
			if oldValue >= threshold16 {
				newValue = white16
			}

			d.buf[bufIx] = newValue

			qErr := oldValue - newValue

			for i, diff := range k.Rows[0] {
				diffPixCol := pixCol + i + 1
				if diffPixCol == pixDx {
					break
				}

				d.buf[pixRow*pixDx+diffPixCol] += qErr * diff / k.Base
			}

			for j := 1; j < len(k.Rows); j++ {
				diffPixRow := pixRow + j
				if diffPixRow == pixDy {
					break
				}

				diffBufOffset := diffPixRow * pixDx
				pixColOffset := len(k.Rows[j]) / 2

				for i, diff := range k.Rows[j] {
					diffPixCol := pixCol + i - pixColOffset

					if diffPixCol < 0 {
						continue
					} else if diffPixCol == pixDx {
						break
					}

					d.buf[diffBufOffset+diffPixCol] += qErr * diff / k.Base
				}
			}
		}
	}

	dx := pixDx / blockWidth
	dy := pixDy / blockHeight

	for row := 0; row < dy; row++ {
		offset := row * dx
		y0 := pixR.Min.Y + row*blockHeight
		for col := 0; col < dx; col++ {
			ix := offset + col

			var cp CodePoint

			x0 := pixR.Min.X + col*blockWidth

			for k := 0; k < blockSize; k++ {
				x, y := x0+k%2, y0+k/2
				bufIx := y*pixDx + x
				mask := CodePoint(1 << bitPos[k])
				if d.buf[bufIx] > 0 {
					cp |= mask
				}
			}

			dst.Cps[ix] = cp
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

func gray16At(img image.Image, x, y int) uint16 {
	return color.Gray16Model.Convert(img.At(x, y)).(color.Gray16).Y
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
