package dots

import (
	"testing"
)

type nullWriter struct{}

func (nw nullWriter) Write(b []byte) (int, error) {
	return len(b), nil
}

func BenchmarkRender(b *testing.B) {
	img := NewImage(40, 10)

	for i := 0; i < b.N; i++ {
		var rd Renderer
		var nw nullWriter

		rd.Render(nw, img)
	}
}

func BenchmarkStringer(b *testing.B) {
	img := NewImage(40, 10)

	for i := 0; i < b.N; i++ {
		_ = img.String()
	}
}

func BenchmarkClear(b *testing.B) {
	img := NewImage(40, 10)

	for i := 0; i < b.N; i++ {
		_ = img.Clear()
	}
}

func BenchmarkReverseByX(b *testing.B) {
	img := NewImage(40, 10)

	for i := 0; i < b.N; i++ {
		_ = img.ReverseByX()
	}
}

func BenchmarkReverseByY(b *testing.B) {
	img := NewImage(40, 10)

	for i := 0; i < b.N; i++ {
		_ = img.ReverseByY()
	}
}
