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

//NewImage creates empty image with
//given size in CodePoints.
func NewImage(r image.Rectangle) *DotImage {
	return &DotImage{
		Cps:    make([]CodePoint, 1*r.Dx()*r.Dy()),
		CpRect: r,
		Stride: 1 * r.Dx(),
	}
}

//Bounds returns image size in pixels.
func (p *DotImage) Bounds() image.Rectangle {
	r := p.CpRect

	r.Min.X *= 2
	r.Max.X *= 2

	r.Min.Y *= 4
	r.Max.Y *= 4

	return r
}

//At uses index of pixel
func (p *DotImage) At(px, py int) color.Color {
	x, cpx := px/blockWidth, px%blockWidth
	y, cpy := py/blockHeight, py%blockHeight

	if p.Cps[p.CpOffset(x, y)].IsOn(cpx, cpy) {
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

//SubPic returns a picture inside of r.
//The returned value is shared with original picture.
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

func (p *DotImage) DrawImageTransform(
	p2 *DotImage,
	transform func(CodePoint, CodePoint) CodePoint) {
	r := p.CpRect.Intersect(p2.CpRect)
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			ix0 := p.CpOffset(x, y)
			ix2 := p2.CpOffset(x, y)
			p.Cps[ix0] = transform(p.Cps[ix0], p2.Cps[ix2])
		}
	}
}

func (p *DotImage) DrawImage(x, y int, p2 *DotImage) {
	p.DrawImageTransform(p2, NEWONLY)
}

func (p *DotImage) FlipBits() *DotImage {
	for i := range p.Cps {
		p.Cps[i] = ^p.Cps[i]
	}
	return p
}

func (p *DotImage) ReverseByX() *DotImage {
	r := p.CpRect
	centerX := (r.Min.X + r.Max.X) / 2

	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < centerX; x++ {
			ix1 := p.CpOffset(x, y)
			ix2 := p.CpOffset(r.Max.X-x-1, y)
			p.Cps[ix1], p.Cps[ix2] = p.Cps[ix2], p.Cps[ix1]
		}
	}

	for i := range p.Cps {
		p.Cps[i] = p.Cps[i].RevX()
	}

	return p
}

func (p *DotImage) ReverseByY() *DotImage {
	r := p.CpRect
	centerY := (r.Min.Y + r.Max.Y) / 2

	for y := r.Min.Y; y < centerY; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			ix1 := p.CpOffset(x, y)
			ix2 := p.CpOffset(x, r.Max.Y-y-1)
			p.Cps[ix1], p.Cps[ix2] = p.Cps[ix2], p.Cps[ix1]
		}
	}

	for i := range p.Cps {
		p.Cps[i] = p.Cps[i].RevY()
	}

	return p
}

//ByteLen returns number of bytes
//required to render image.
func (p *DotImage) ByteLen() int {
	return (3*p.CpRect.Dx() + 1) * p.CpRect.Dy()
}

func (p *DotImage) String() string {
	buf := make([]byte, p.ByteLen())

	p.read(buf)

	//using unsafe code for better performance
	return *(*string)(unsafe.Pointer(&buf))
}

func (p *DotImage) read(buf []byte) {
	r := p.CpRect
	dx := r.Dx()

	for y := r.Min.Y; y < r.Max.Y; y++ {
		line := y - r.Min.Y
		for x := r.Min.X; x < r.Max.X; x++ {
			col := x - r.Min.X
			cpb := byte(p.Cps[p.CpOffset(x, y)])

			ix := line*(3*dx+1) + 3*col
			buf[ix+0] = 226
			buf[ix+1] = tx | (160|(cpb>>6))&maskx
			buf[ix+2] = tx | cpb&maskx
		}
		buf[line*(3*dx+1)+3*dx] = '\n'
	}
}
