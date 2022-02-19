package dots

import (
	"image"
	"image/color"
	"unsafe"
)

type DotImage struct {
	CpRect image.Rectangle
	Stride int
	Cps    []CodePoint
}

// NewImage creates empty image with given size in CodePoints.
func NewImage(r image.Rectangle) *DotImage {
	return &DotImage{
		Cps:    make([]CodePoint, r.Dx()*r.Dy()),
		CpRect: r,
		Stride: r.Dx(),
	}
}

// Bounds returns image size in pixels.
func (p *DotImage) Bounds() image.Rectangle {
	r := p.CpRect

	r.Min.X *= blockWidth
	r.Max.X *= blockWidth

	r.Min.Y *= blockHeight
	r.Max.Y *= blockHeight

	return r
}

// At uses index of pixel
func (p *DotImage) At(px, py int) color.Color {
	x, cpx := px/blockWidth, px%blockWidth
	y, cpy := py/blockHeight, py%blockHeight

	cp := p.CpAt(x, y)

	if cp&(1<<(bitPos[cpx+cpy*blockWidth])) != 0 {
		return color.White
	}

	return color.Black
}

func (p *DotImage) ColorModel() color.Model {
	return color.Gray16Model
}

func (p *DotImage) CpOffset(x, y int) int {
	return (y-p.CpRect.Min.Y)*p.Stride + (x - p.CpRect.Min.X)
}

func (p *DotImage) CpAt(x, y int) CodePoint {
	return p.Cps[p.CpOffset(x, y)]
}

func (p *DotImage) Clear() *DotImage {
	for i := range p.Cps {
		p.Cps[i] = 0
	}
	return p
}

func (p *DotImage) Fill(cp CodePoint) *DotImage {
	for i := range p.Cps {
		p.Cps[i] = cp
	}
	return p
}

// SubPic returns a picture inside of r.
// The returned value is shared with original picture.
func (p *DotImage) SubImage(r image.Rectangle) *DotImage {
	r = r.Intersect(p.CpRect)
	i := p.CpOffset(r.Min.X, r.Min.Y)
	return &DotImage{
		Cps:    p.Cps[i:],
		Stride: p.Stride,
		CpRect: r,
	}
}

const (
	tx    = 0b10000000
	maskx = 0b00111111
)

func (p *DotImage) FlipBits() *DotImage {
	for i := range p.Cps {
		p.Cps[i] = ^p.Cps[i]
	}
	return p
}

func (p *DotImage) FlipH() *DotImage {
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

func (p *DotImage) FlipV() *DotImage {
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
func (p *DotImage) ByteLen() int {
	return (p.CpRect.Dx()*runeSize + 1) * p.CpRect.Dy()
}

func (p *DotImage) String() string {
	buf := make([]byte, p.ByteLen())

	p.read(buf)

	return *(*string)(unsafe.Pointer(&buf))
}

func (p *DotImage) read(buf []byte) {
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
