package dots

import (
	"image"
	"io"
	"testing"
)

var testR = image.Rect(0, 0, 40, 10)

func BenchmarkRender(b *testing.B) {
	p := NewImage(testR)

	for i := 0; i < b.N; i++ {
		var rd Renderer
		rd.Render(io.Discard, p)
	}
}

func BenchmarkStringer(b *testing.B) {
	p := NewImage(testR)

	for i := 0; i < b.N; i++ {
		_ = p.String()
	}
}

func BenchmarkClear(b *testing.B) {
	p := NewImage(testR)

	for i := 0; i < b.N; i++ {
		_ = p.Clear()
	}
}

func BenchmarkReverseByX(b *testing.B) {
	p := NewImage(testR)

	for i := 0; i < b.N; i++ {
		_ = p.ReverseByX()
	}
}

func BenchmarkReverseByY(b *testing.B) {
	p := NewImage(testR)

	for i := 0; i < b.N; i++ {
		_ = p.ReverseByY()
	}
}
