package dots

import (
	"unsafe"
)

//Sizer types must have fixed
//number of characters per line and total number of lines.
type Sizer interface {
	//Size returns (width, height).
	Size() (int, int)
}

type Image struct {
	w   int
	h   int
	cps []CodePoint
}

//Size returns width and height.
func (img *Image) Size() (int, int) {
	return img.w, img.h
}

func (img *Image) Clear() *Image {
	for i := range img.cps {
		img.cps[i] = 0
	}
	return img
}

func (img *Image) Fill(cp CodePoint) *Image {
	for i := range img.cps {
		img.cps[i] = cp
	}
	return img
}

const (
	tx    = 0b10000000
	maskx = 0b00111111
)

func NewImage(width, height int) *Image {
	cps := make([]CodePoint, width*height)
	return &Image{width, height, cps}
}

func (img *Image) DrawImageTransform(
	x, y int,
	img2 *Image,
	transform func(CodePoint, CodePoint) CodePoint) {
	h := min(img2.h, img.h-y)
	w := min(img2.w, img.w-x)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ix0 := (i+y)*img.w + j + x
			ix2 := i*img2.w + j
			img.cps[ix0] = transform(img.cps[ix0], img2.cps[ix2])
		}
	}
}

func (img *Image) DrawImage(x, y int, img2 *Image) {
	img.DrawImageTransform(x, y, img2, NEWONLY)
}

//Clone image data from left top corner at (x,y)
//into another image.
func (img *Image) Clone(x, y int, img2 *Image) {
	h := min(img2.h, img.h-y)
	w := min(img2.w, img.w-x)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ix0 := (i+y)*img.w + j + x
			ix2 := i*img2.w + j
			img2.cps[ix2] = img.cps[ix0]
		}
	}
}

func (img *Image) FlipBits() *Image {
	for i := range img.cps {
		img.cps[i] = ^img.cps[i]
	}
	return img
}

func (img *Image) SetCP(x, y int, cp CodePoint) {
	ix := x + y*img.w
	img.cps[ix] = cp
}

func (img *Image) ReverseByX() *Image {
	for i := 0; i < img.h; i++ {
		off := i * img.w
		for j := 0; j < img.w/2; j++ {
			ix1 := off + j
			ix2 := off + img.w - j - 1
			img.cps[ix1], img.cps[ix2] = img.cps[ix2], img.cps[ix1]
		}
	}

	for i := range img.cps {
		img.cps[i] = img.cps[i].revX()
	}

	return img
}

func (img *Image) ReverseByY() *Image {
	for i := 0; i < img.h/2; i++ {
		off1 := i * img.w
		off2 := (img.h - i - 1) * img.w
		for j := 0; j < img.w; j++ {
			ix1 := off1 + j
			ix2 := off2 + j
			img.cps[ix1], img.cps[ix2] = img.cps[ix2], img.cps[ix1]
		}
	}

	for i := range img.cps {
		img.cps[i] = img.cps[i].revY()
	}

	return img
}

//ByteLen returns number of bytes
//required to render image.
func (img *Image) ByteLen() int {
	return (3*img.w + 1) * img.h
}

func (img *Image) String() string {
	buf := make([]byte, img.ByteLen())

	img.read(buf)

	//using unsafe code for better performance
	return *(*string)(unsafe.Pointer(&buf))
}

func (img *Image) read(buf []byte) {
	for i := 0; i < img.h; i++ {
		for j := 0; j < img.w; j++ {
			ix := i*img.w + j
			cpb := byte(img.cps[ix])
			buf[ix*3+i] = 226
			buf[ix*3+i+1] = tx | (160|(cpb>>6))&maskx
			buf[ix*3+i+2] = tx | cpb&maskx
		}
		buf[i*(3*img.w+1)+3*img.w] = '\n'
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
