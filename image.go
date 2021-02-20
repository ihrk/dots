package dots

import (
	"image"
	"unsafe"
)

type DotPic struct {
	Rect   image.Rectangle
	Stride int
	Cps    []CodePoint
}

func NewPic(x, y int) *DotPic {
	return &DotPic{
		Cps:    make([]CodePoint, x*y),
		Rect:   image.Rect(0, 0, x, y),
		Stride: x,
	}
}

func (p *DotPic) CpOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x - p.Rect.Min.X)
}

func (p *DotPic) CpAt(x, y int) CodePoint {
	return p.Cps[p.CpOffset(x, y)]
}

func (p *DotPic) Clear() *DotPic {
	for i := range p.Cps {
		p.Cps[i] = 0
	}
	return p
}

func (p *DotPic) Fill(cp CodePoint) *DotPic {
	for i := range p.Cps {
		p.Cps[i] = cp
	}
	return p
}

//SubPic returns a picture inside of r.
//The returned value is shared with original picture.
func (p *DotPic) SubPic(r image.Rectangle) *DotPic {
	r = r.Intersect(p.Rect)
	i := p.CpOffset(r.Min.X, r.Min.Y)
	return &DotPic{
		Cps:    p.Cps[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

const (
	tx    = 0b10000000
	maskx = 0b00111111
)

func (p *DotPic) DrawImageTransform(
	p2 *DotPic,
	transform func(CodePoint, CodePoint) CodePoint) {
	r := p.Rect.Intersect(p2.Rect)
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			ix0 := p.CpOffset(x, y)
			ix2 := p2.CpOffset(x, y)
			p.Cps[ix0] = transform(p.Cps[ix0], p2.Cps[ix2])
		}
	}
}

func (p *DotPic) DrawImage(x, y int, p2 *DotPic) {
	p.DrawImageTransform(p2, NEWONLY)
}

func (p *DotPic) FlipBits() *DotPic {
	for i := range p.Cps {
		p.Cps[i] = ^p.Cps[i]
	}
	return p
}

func (p *DotPic) ReverseByX() *DotPic {
	r := p.Rect
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

func (p *DotPic) ReverseByY() *DotPic {
	r := p.Rect
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
func (p *DotPic) ByteLen() int {
	return (3*p.Rect.Dx() + 1) * p.Rect.Dy()
}

func (p *DotPic) String() string {
	buf := make([]byte, p.ByteLen())

	p.read(buf)

	//using unsafe code for better performance
	return *(*string)(unsafe.Pointer(&buf))
}

func (p *DotPic) read(buf []byte) {
	r := p.Rect
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
