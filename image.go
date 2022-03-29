package dots

import (
	"image"
	"image/color"
	"unsafe"
)

type Image struct {
	CpRect image.Rectangle
	Stride int
	Cps    []CodePoint
}

func PixRectToCpRect(r image.Rectangle) image.Rectangle {
	r.Min.X = floor(r.Min.X, blockWidth)
	r.Max.X = ceil(r.Max.X, blockWidth)

	r.Min.Y = floor(r.Min.Y, blockHeight)
	r.Max.Y = ceil(r.Max.Y, blockHeight)

	return r
}

// NewImage creates empty image with given size in CodePoints.
func NewImage(r image.Rectangle) *Image {
	return &Image{
		Cps:    make([]CodePoint, r.Dx()*r.Dy()),
		CpRect: r,
		Stride: r.Dx(),
	}
}

// NewImagePix creates empty image with given size in pixels.
func NewImagePix(r image.Rectangle) *Image {
	return NewImage(PixRectToCpRect(r))
}

// Bounds returns image size in pixels.
func (p *Image) Bounds() image.Rectangle {
	r := p.CpRect

	r.Min.X *= blockWidth
	r.Max.X *= blockWidth

	r.Min.Y *= blockHeight
	r.Max.Y *= blockHeight

	return r
}

type Color struct {
	White bool
}

func (c Color) RGBA() (r, g, b, a uint32) {
	if c.White {
		return white16, white16, white16, white16
	}

	return 0, 0, 0, white16
}

func (p *Image) At(px, py int) color.Color {
	return p.DotAt(px, py)
}

func (p *Image) RGBA64At(x, y int) color.RGBA64 {
	r, g, b, a := p.DotAt(x, y).RGBA()
	return color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)}
}

func (p *Image) DotAt(px, py int) Color {

	x, bitX := px/blockWidth, px%blockWidth
	y, bitY := py/blockHeight, py%blockHeight

	if !(image.Point{x, y}.In(p.CpRect)) {
		return Color{}
	}

	cp := p.CpAt(x, y)

	bitIx := bitPos[bitX+bitY*blockWidth]

	return Color{cp&(1<<bitIx) != 0}
}

func (p *Image) ColorModel() color.Model {
	return ColorModel
}

var ColorModel = color.ModelFunc(colorModel)

func colorModel(c color.Color) color.Color {
	if _, ok := c.(Color); ok {
		return c
	}

	r, g, b, _ := c.RGBA()

	return Color{gray16(r, g, b) >= threshold16}
}

func (p *Image) CpOffset(x, y int) int {
	return (y-p.CpRect.Min.Y)*p.Stride + (x - p.CpRect.Min.X)
}

func (p *Image) CpAt(x, y int) CodePoint {
	return p.Cps[p.CpOffset(x, y)]
}

func (p *Image) Clear() *Image {
	if p.CpRect.Dx() == p.Stride {
		// fast path for a full width image
		for i := range p.Cps {
			p.Cps[i] = 0
		}
	} else {
		p.Fill(0)
	}

	return p
}

func (p *Image) Fill(cp CodePoint) *Image {
	dx := p.CpRect.Dx()
	dy := p.CpRect.Dy()

	for row := 0; row < dy; row++ {
		for col := 0; col < dx; col++ {
			p.Cps[row*p.Stride+col] = cp
		}
	}

	return p
}

// SubImage returns a picture inside of r in CodePoints.
// The returned value is shared with original picture.
func (p *Image) SubImage(r image.Rectangle) *Image {
	r = r.Intersect(p.CpRect)

	return &Image{
		Cps:    p.Cps[p.CpOffset(r.Min.X, r.Min.Y):p.CpOffset(r.Max.X, r.Max.Y)],
		Stride: p.Stride,
		CpRect: r,
	}
}

// SubImagePix returns a picture inside of r in pixels.
// The returned value is shared with original picture.
func (p *Image) SubImagePix(r image.Rectangle) *Image {
	return p.SubImage(PixRectToCpRect(r))
}

func (p *Image) FlipBits() *Image {
	for i := range p.Cps {
		p.Cps[i] = ^p.Cps[i]
	}
	return p
}

func (p *Image) FlipH() *Image {
	r := p.CpRect
	dx := r.Dx()
	dy := r.Dy()
	centerX := dx / 2

	for row := 0; row < dy; row++ {
		for col := 0; col < centerX; col++ {
			ix1 := row*p.Stride + col
			ix2 := row*p.Stride + dx - col - 1
			p.Cps[ix1], p.Cps[ix2] = p.Cps[ix2], p.Cps[ix1]
		}
	}

	for i := range p.Cps {
		p.Cps[i] = CpFlipH(p.Cps[i])
	}

	return p
}

func (p *Image) FlipV() *Image {
	r := p.CpRect
	dx := r.Dx()
	dy := r.Dy()
	centerY := dy / 2

	for row := 0; row < centerY; row++ {
		for col := 0; col < dx; col++ {
			ix1 := row*p.Stride + col
			ix2 := (dy-row-1)*p.Stride + col
			p.Cps[ix1], p.Cps[ix2] = p.Cps[ix2], p.Cps[ix1]
		}
	}

	for i := range p.Cps {
		p.Cps[i] = CpFlipV(p.Cps[i])
	}

	return p
}

// ByteLen returns number of bytes required to render image.
func (p *Image) ByteLen() int {
	return (p.CpRect.Dx()*runeSize + 1) * p.CpRect.Dy()
}

func (p *Image) String() string {
	buf := make([]byte, p.ByteLen())

	p.read(buf)

	return *(*string)(unsafe.Pointer(&buf))
}

const (
	tx    = 0b10000000
	maskx = 0b00111111
)

func (p *Image) read(buf []byte) {
	r := p.CpRect
	dx := r.Dx()
	dy := r.Dy()
	rowSize := dx*runeSize + 1

	for row := 0; row < dy; row++ {
		for col := 0; col < dx; col++ {
			cpb := byte(p.Cps[row*p.Stride+col])

			ix := row*rowSize + col*runeSize
			buf[ix+0] = 226
			buf[ix+1] = tx | (160|(cpb>>6))&maskx
			buf[ix+2] = tx | cpb&maskx
		}
		buf[(row+1)*rowSize-1] = '\n'
	}
}
