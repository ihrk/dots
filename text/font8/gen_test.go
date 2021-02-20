package font8

import (
	"testing"

	"github.com/ihrk/dots"
)

func BenchmarkLoad(b *testing.B) {
	var f Font8
	for i := 0; i < b.N; i++ {
		f.LoadChar('0')
	}
}

func BenchmarkFill(b *testing.B) {
	var f Font8
	f.LoadChar('0')
	img := dots.NewImage(3, 3)

	for i := 0; i < b.N; i++ {
		f.FillImage(img)
	}
}

func BenchmarkDraw(b *testing.B) {
	var f Font8
	f.LoadChar('0')
	img := dots.NewImage(3, 3)
	f.FillImage(img)

	frame := dots.NewImage(10, 10)
	for i := 0; i < b.N; i++ {
		frame.DrawImage(3, 3, img)
	}
}

func BenchmarkDrawT(b *testing.B) {
	var f Font8
	f.LoadChar('0')
	img := dots.NewImage(3, 3)
	f.FillImage(img)

	frame := dots.NewImage(10, 10)
	for i := 0; i < b.N; i++ {
		frame.DrawImageTransform(3, 3, img, dots.AND)
	}
}
