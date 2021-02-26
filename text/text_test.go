package text

import (
	"testing"

	"github.com/ihrk/dots"
	"github.com/ihrk/dots/text/font8"
)

const testStr = "2"

func BenchmarkDisplay(b *testing.B) {
	var (
		f   = new(font8.Font8)
		img = dots.NewDotImage(len(testStr)*3, 3)
	)

	for i := 0; i < b.N; i++ {
		DisplayString(f, testStr, img)
	}
}
