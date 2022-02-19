package dots

import (
	"image"
	"io"
	"testing"
)

var testR = image.Rect(0, 0, 40, 10)

func BenchmarkRender(b *testing.B) {
	p := NewImage(testR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var rd Renderer
		rd.Render(io.Discard, p)
	}
}

func BenchmarkRenderBuffered(b *testing.B) {
	p := NewImage(testR)
	rd := NewRenderer(make([]byte, p.ByteLen()))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rd.Render(io.Discard, p)
	}
}

func BenchmarkStringer(b *testing.B) {
	p := NewImage(testR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.String()
	}
}

func BenchmarkClear(b *testing.B) {
	p := NewImage(testR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.Clear()
	}
}

func BenchmarkFlipH(b *testing.B) {
	p := NewImage(testR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.FlipH()
	}
}

func BenchmarkFlipV(b *testing.B) {
	p := NewImage(testR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.FlipV()
	}
}

func BenchmarkDithering(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ErrDiffDithering(testR, Atkinson)
	}
}

func BenchmarkDitheringGray16(b *testing.B) {
	gray := image.NewGray16(testR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ErrDiffDithering(gray, Atkinson)
	}
}

func BenchmarkThresholdingGray16(b *testing.B) {
	gray := image.NewGray16(testR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Thresholding(gray, 0, 128)
	}
}

func BenchmarkOrderedGray16(b *testing.B) {
	gray := image.NewGray16(testR)

	thMap := GenerateThresholdMap(1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OrderedDithering(gray, thMap)
	}
}
