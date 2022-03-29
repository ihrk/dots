package dots

import (
	"image"
	"io"
	"testing"
)

var testR = image.Rect(0, 0, 1920, 1080)

func BenchmarkRender(b *testing.B) {
	p := NewImagePix(testR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var rd Renderer
		rd.Render(io.Discard, p)
	}
}

func BenchmarkRenderBuffered(b *testing.B) {
	p := NewImagePix(testR)
	rd := NewRenderer(make([]byte, p.ByteLen()))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rd.Render(io.Discard, p)
	}
}

func BenchmarkStringer(b *testing.B) {
	p := NewImagePix(testR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.String()
	}
}

func BenchmarkClear(b *testing.B) {
	p := NewImagePix(testR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.Clear()
	}
}

func BenchmarkFill(b *testing.B) {
	p := NewImagePix(testR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.Fill(1)
	}
}

func BenchmarkFlipH(b *testing.B) {
	p := NewImagePix(testR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.FlipH()
	}
}

func BenchmarkFlipV(b *testing.B) {
	p := NewImagePix(testR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.FlipV()
	}
}

func BenchmarkDithering(b *testing.B) {
	src := image.NewNRGBA(testR)
	dst := NewImagePix(testR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ErrorDiffusionDithering(src, dst, Atkinson)
	}
}

func BenchmarkDitheringBuffered(b *testing.B) {
	src := image.NewNRGBA(testR)
	dst := NewImagePix(testR)

	var d Ditherer
	d.checkBuf(testR.Dx() * testR.Dy())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.ErrorDiffusion(src, dst, Atkinson)
	}
}

func BenchmarkThresholding(b *testing.B) {
	src := image.NewNRGBA(testR)
	dst := NewImagePix(testR)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Thresholding(src, dst, 128)
	}
}

func BenchmarkOrdered(b *testing.B) {
	src := image.NewNRGBA(testR)
	dst := NewImagePix(testR)

	thMap := GenerateThresholdMap(1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OrderedDithering(src, dst, thMap)
	}
}
