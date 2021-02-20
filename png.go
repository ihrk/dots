package dots

import (
	"image"
	"image/color"
	"image/png"
	"io"

	"os"
)

const (
	blockWidth  = 2
	blockHeight = 4
	blockSize   = blockWidth * blockHeight
)

var bitPos = [blockSize]int{0, 3, 1, 4, 2, 5, 6, 7}

type Opts struct {
	bg CodePoint //background
	th uint8     //threshold
}

func NewOpts(bg CodePoint, th uint8) *Opts {
	return &Opts{bg, th}
}

func (o *Opts) getThreshold() uint8 {
	if o == nil {
		return 128
	}
	return o.th
}

func (o *Opts) getBackground() CodePoint {
	if o == nil {
		return 149
	}
	return o.bg
}

func isTransparent(x, y int, img image.Image) bool {
	_, _, _, a := img.At(x, y).RGBA()
	return a == 0
}

func getBit(x, y int, th uint8, img image.Image) bool {
	c := color.GrayModel.Convert(img.At(x, y)).(color.Gray).Y

	return c >= th
}

func RenderPNG(r io.Reader, opts *Opts) (*Image, error) {
	pngImage, err := png.Decode(r)
	if err != nil {
		return nil, err
	}

	w := pngImage.Bounds().Max.X / 2
	h := pngImage.Bounds().Max.Y / 4
	img := NewImage(w, h)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ix := i*img.w + j

			x0 := j * 2
			y0 := i * 4

			cp := opts.getBackground()

			for k := 0; k < blockSize; k++ {
				x, y := x0+k%2, y0+k/2

				mask := CodePoint(1 << bitPos[k])

				if !isTransparent(x, y, pngImage) {
					if getBit(x, y, opts.getThreshold(), pngImage) {
						cp |= mask
					} else {
						cp &= ^mask
					}
				}

			}

			img.cps[ix] = cp
		}
	}

	return img, nil
}

func RenderPNGFromFile(path string, opts *Opts) (*Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return RenderPNG(file, opts)
}