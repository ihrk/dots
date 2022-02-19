package dots

import "image"

// ThresholdMap size is expected to be power of two.
// Valid map example:
// [0 2]
// [3 1]
type ThresholdMap [][]uint16

func OrderedDithering(src image.Image, thresholdMap ThresholdMap) *DotImage {
	l := len(thresholdMap)
	size := uint16(l * l)

	if size < 2 {
		panic("OrderedDithering: invalid threshold map size")
	}

	thBase := threshold16 / size * 2

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
				th := thresholdMap[y%l][x%l]
				g := gray16At(src, x, y)
				// Error diffusion and threshold
				// use '>=' (more or equal) because
				// they try to match closest to
				// original color, which is not
				// the case for ordered dithering,
				// so '>' (only if more) is used.
				if g > thBase*th {
					cp |= mask
				}
			}

			p.Cps[ix] = cp
		}
	}

	return p
}

func GenerateThresholdMap(n int) ThresholdMap {
	if n < 1 {
		return ThresholdMap{{0}}
	}

	if n > 7 {
		panic("GenerateThresholdMap: size is too big")
	}

	quarter := GenerateThresholdMap(n - 1)
	qlen := len(quarter)

	m := make(ThresholdMap, 2*qlen)

	for i := range m {
		m[i] = make([]uint16, 2*qlen)
	}

	for i := 0; i < qlen; i++ {
		for j := 0; j < qlen; j++ {
			m[i][j] = 4 * quarter[i][j]
		}
	}

	for i := 0; i < qlen; i++ {
		for j := 0; j < qlen; j++ {
			m[i+qlen][j] = 4*quarter[i][j] + 3
		}
	}

	for i := 0; i < qlen; i++ {
		for j := 0; j < qlen; j++ {
			m[i][j+qlen] = 4*quarter[i][j] + 2
		}
	}

	for i := 0; i < qlen; i++ {
		for j := 0; j < qlen; j++ {
			m[i+qlen][j+qlen] = 4*quarter[i][j] + 1
		}
	}

	return m
}
