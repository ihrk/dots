package font8

import (
	"testing"
)

func BenchmarkLoad(b *testing.B) {
	var f Font8
	for i := 0; i < b.N; i++ {
		f.LoadChar('0')
	}
}
