package dots

import (
	"image"
)

// ThresholdMap valid example:
// [0 2]
// [3 1]
type ThresholdMap struct {
	Base uint16
	Rows [][]uint16
}

func OrderedDithering(src image.Image, dst *Image, thMap ThresholdMap) {
	l := len(thMap.Rows)

	r := PixRectToCpRect(src.Bounds()).Intersect(dst.CpRect)

	srcRGBA64 := toRGBA64(src)

	for row := r.Min.Y; row < r.Max.Y; row++ {
		y0 := row * blockHeight
		for col := r.Min.X; col < r.Max.X; col++ {
			var cp CodePoint

			x0 := col * blockWidth

			for k := 0; k < blockSize; k++ {
				x, y := x0+k%2, y0+k/2

				threshold := thMap.Base * thMap.Rows[y%l][x%l]
				// Error diffusion and threshold
				// use '>=' (more or equal) because
				// they try to match closest to
				// original color, which is not
				// the case for ordered dithering,
				// so '>' (only if more) is used.
				if gray16At(srcRGBA64, x, y) > threshold {
					cp |= CodePoint(1 << bitPos[k])
				}
			}

			dst.Cps[dst.CpOffset(col, row)] = cp
		}
	}
}

func GenerateThresholdMap(n int) ThresholdMap {
	if n < 2 {
		return ThresholdMap{
			Base: threshold16 / 2,
			Rows: [][]uint16{
				{0, 2},
				{3, 1},
			},
		}
	}

	if n > 8 {
		panic("GenerateThresholdMap: size is too big")
	}

	quarter := GenerateThresholdMap(n - 1)
	qlen := len(quarter.Rows)

	m := ThresholdMap{
		Base: quarter.Base / 4,
		Rows: make([][]uint16, qlen*2),
	}

	for i := range m.Rows {
		m.Rows[i] = make([]uint16, 2*qlen)
	}

	for i := 0; i < qlen; i++ {
		for j := 0; j < qlen; j++ {
			m.Rows[i][j] = 4 * quarter.Rows[i][j]
		}
	}

	for i := 0; i < qlen; i++ {
		for j := 0; j < qlen; j++ {
			m.Rows[i+qlen][j] = 4*quarter.Rows[i][j] + 3
		}
	}

	for i := 0; i < qlen; i++ {
		for j := 0; j < qlen; j++ {
			m.Rows[i][j+qlen] = 4*quarter.Rows[i][j] + 2
		}
	}

	for i := 0; i < qlen; i++ {
		for j := 0; j < qlen; j++ {
			m.Rows[i+qlen][j+qlen] = 4*quarter.Rows[i][j] + 1
		}
	}

	return m
}
