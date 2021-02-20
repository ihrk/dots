package dots

import (
	"testing"
)

type nullWriter struct{}

func (nw nullWriter) Write(b []byte) (int, error) {
	return len(b), nil
}

func BenchmarkRender(b *testing.B) {
	p := NewPic(40, 10)

	for i := 0; i < b.N; i++ {
		var rd Renderer
		var nw nullWriter

		rd.Render(nw, p)
	}
}

func BenchmarkStringer(b *testing.B) {
	p := NewPic(40, 10)

	for i := 0; i < b.N; i++ {
		_ = p.String()
	}
}

func BenchmarkClear(b *testing.B) {
	p := NewPic(40, 10)

	for i := 0; i < b.N; i++ {
		_ = p.Clear()
	}
}

func BenchmarkReverseByX(b *testing.B) {
	p := NewPic(40, 10)

	for i := 0; i < b.N; i++ {
		_ = p.ReverseByX()
	}
}

func BenchmarkReverseByY(b *testing.B) {
	p := NewPic(40, 10)

	for i := 0; i < b.N; i++ {
		_ = p.ReverseByY()
	}
}
