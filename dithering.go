package dots

import (
	"image"
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
		for i := 0; i < n; i++ {
			d.buf[i] = 0
		}

		return
	}
	d.buf = make([]int32, len(d.buf)+n)
}

func (d *Ditherer) ErrorDiffusion(src image.Image, dst *Image, k DiffusionKernel) {
	r := PixRectToCpRect(src.Bounds()).Intersect(dst.CpRect)

	pixDx := r.Dx() * blockWidth
	pixDy := r.Dy() * blockHeight

	pixMinX := r.Min.X * blockWidth
	pixMinY := r.Min.Y * blockHeight

	d.checkBuf(pixDx * pixDy)

	srcRGBA64 := toRGBA64(src)

	for pixRow := 0; pixRow < pixDy; pixRow++ {
		pixY := pixMinY + pixRow
		for pixCol := 0; pixCol < pixDx; pixCol++ {
			pixX := pixMinX + pixCol
			bufIx := pixRow*pixDx + pixCol

			qErr := int32(gray16At(srcRGBA64, pixX, pixY)) + d.buf[bufIx]

			if qErr >= threshold16 {
				qErr -= white16
				d.buf[bufIx] = 1
			} else {
				d.buf[bufIx] = 0
			}

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

	for row := r.Min.Y; row < r.Max.Y; row++ {
		y0 := row*blockHeight - pixMinY
		for col := r.Min.X; col < r.Max.X; col++ {
			var cp CodePoint

			x0 := col*blockWidth - pixMinX

			for k := 0; k < blockSize; k++ {
				x, y := x0+k%2, y0+k/2

				if d.buf[y*pixDx+x] > 0 {
					cp |= CodePoint(1 << bitPos[k])
				}
			}

			dst.Cps[dst.CpOffset(col, row)] = cp
		}
	}
}

func ErrorDiffusionDithering(src image.Image, dst *Image, k DiffusionKernel) {
	var d Ditherer

	d.ErrorDiffusion(src, dst, k)
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
