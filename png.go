package dots

import (
	"image"
	"image/color"
	"image/png"
	"io"

	"os"
)

const (
	//DefBg is default background value.
	DefBg = 149
	//DefTh is default threshold value.
	DefTh = 128
)

type Opts struct {
	bg    CodePoint //background
	th    uint8     //threshold
	isVar bool      //if true sets threshold to average of pixelblock value
}

func NewOpts(bg CodePoint, th uint8, isVar bool) *Opts {
	return &Opts{bg, th, isVar}
}

func (o *Opts) getThreshold() uint8 {
	if o == nil {
		return DefTh
	}
	return o.th
}

func (o *Opts) getBackground() CodePoint {
	if o == nil {
		return DefBg
	}
	return o.bg
}

func (o *Opts) isVariable() bool {
	if o == nil {
		return false
	}
	return o.isVar
}

func isTransparent(x, y int, p image.Image) bool {
	_, _, _, a := p.At(x, y).RGBA()
	return a == 0
}

func getGray(x, y int, p image.Image) uint8 {
	return color.GrayModel.Convert(p.At(x, y)).(color.Gray).Y
}

func NewImageFromPNG(r io.Reader, opts *Opts) (*DotImage, error) {
	pngImage, err := png.Decode(r)
	if err != nil {
		return nil, err
	}

	w := pngImage.Bounds().Max.X / 2
	h := pngImage.Bounds().Max.Y / 4
	p := NewImage(image.Rect(0, 0, w, h))

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ix := i*w + j

			x0 := j * 2
			y0 := i * 4

			cp := opts.getBackground()

			th := opts.getThreshold()

			if opts.isVariable() {
				num, sum := 0, 0

				for k := 0; k < blockSize; k++ {
					x, y := x0+k%2, y0+k/2

					if !isTransparent(x, y, pngImage) {
						num++
						sum += int(getGray(x, y, pngImage))
					}
				}

				if num > 0 && uint8(sum/num) < th {
					th = uint8(sum / num)
				}
			}

			for k := 0; k < blockSize; k++ {
				x, y := x0+k%2, y0+k/2

				mask := CodePoint(1 << bitPos[k])

				if !isTransparent(x, y, pngImage) {
					if getGray(x, y, pngImage) >= th {
						cp |= mask
					} else {
						cp &= ^mask
					}
				}
			}

			p.Cps[ix] = cp
		}
	}

	return p, nil
}

func NewImageFromPNGFile(path string, opts *Opts) (*DotImage, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return NewImageFromPNG(file, opts)
}
