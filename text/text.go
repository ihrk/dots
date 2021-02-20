package text

import "github.com/ihrk/dots"

type Font interface {
	GetCP(int, int) dots.CodePoint
	LoadChar(rune)
	dots.Sizer
}

func DisplayString(f Font, s string, img *dots.Image) {
	iw, ih := img.Size()
	offW := 0
	for _, r := range s {
		if offW >= iw {
			break
		}

		f.LoadChar(r)
		w, h := f.Size()

		w = min(iw-offW, w)
		h = min(ih, h)

		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				img.SetCP(offW+j, i, f.GetCP(j, i))
			}
		}

		offW += w
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
