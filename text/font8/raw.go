package font8

import "github.com/ihrk/dots"

//Font8 is three-line monospace font with
//height of symbols at 8.
type Font8 struct {
	src [3][]rune
}

func (f *Font8) Size() (int, int) {
	return len(f.src[0]), 3
}

func (f *Font8) GetCP(x, y int) dots.CodePoint {
	return dots.FromRune(f.src[y][x])
}

func (f *Font8) FillImage(img *dots.Image) {
	h, w := img.Size()
	for i := range f.src {
		if i >= h {
			break
		}
		col := 0
		for _, r := range f.src[i] {
			if col >= w {
				break
			}
			img.SetCP(col, i, dots.FromRune(r))
			col++
		}
	}
}

var chars = [128][3][]rune{
	'A': [3][]rune{
		[]rune("⠀⡰⡀"),
		[]rune("⢰⠥⢵"),
		[]rune("⠈⠀⠈"),
	},
	'a': [3][]rune{
		[]rune("⠀⣀⡀"),
		[]rune("⢠⠒⣺"),
		[]rune("⠀⠉⠈"),
	},
	'B': [3][]rune{
		[]rune("⢰⠒⢢"),
		[]rune("⢸⠉⢱"),
		[]rune("⠈⠉⠁"),
	},
	'b': [3][]rune{
		[]rune("⢰⢀⡀"),
		[]rune("⢸⡁⢸"),
		[]rune("⠈⠈⠁"),
	},
	'C': [3][]rune{
		[]rune("⢠⠒⠢"),
		[]rune("⢸⠀⢀"),
		[]rune("⠀⠉⠁"),
	},
	'c': [3][]rune{
		[]rune("⠀⣀⡀"),
		[]rune("⢸⠀⢈"),
		[]rune("⠀⠉⠁"),
	},
	'D': [3][]rune{
		[]rune("⢰⠒⢄"),
		[]rune("⢸⠀⡸"),
		[]rune("⠈⠉⠀"),
	},
	'd': [3][]rune{
		[]rune("⠀⣀⢰"),
		[]rune("⢸⠀⣹"),
		[]rune("⠀⠉⠈"),
	},
	'E': [3][]rune{
		[]rune("⢰⠒⠒"),
		[]rune("⢸⠉⠉"),
		[]rune("⠈⠉⠉"),
	},
	'e': [3][]rune{
		[]rune("⠀⣀⡀"),
		[]rune("⢸⠒⢚"),
		[]rune("⠀⠉⠁"),
	},
	'F': [3][]rune{
		[]rune("⢰⠒⠒"),
		[]rune("⢸⠉⠉"),
		[]rune("⠈⠀⠀"),
	},
	'f': [3][]rune{
		[]rune("⢀⣠⣒"),
		[]rune("⠀⢸⠀"),
		[]rune("⠀⠈⠀"),
	},
	'G': [3][]rune{
		[]rune("⢠⠒⠢"),
		[]rune("⢸⠀⢲"),
		[]rune("⠀⠉⠁"),
	},
	'g': [3][]rune{
		[]rune("⠀⣀⢀"),
		[]rune("⢸⠀⣹"),
		[]rune("⠀⠭⠜"),
	},
	'H': [3][]rune{
		[]rune("⢰⠀⢰"),
		[]rune("⢸⠉⢹"),
		[]rune("⠈⠀⠈"),
	},
	'h': [3][]rune{
		[]rune("⢰⢀⡀"),
		[]rune("⢸⠁⢸"),
		[]rune("⠈⠀⠈"),
	},
	'I': [3][]rune{
		[]rune("⠐⢲⠒"),
		[]rune("⠀⢸⠀"),
		[]rune("⠈⠉⠉"),
	},
	'i': [3][]rune{
		[]rune("⠀⣐⠀"),
		[]rune("⠀⢸⠀"),
		[]rune("⠀⠀⠉"),
	},
	'J': [3][]rune{
		[]rune("⠀⠒⡖"),
		[]rune("⢀⠀⡇"),
		[]rune("⠀⠉⠀"),
	},
	'j': [3][]rune{
		[]rune("⠀⢀⡂"),
		[]rune("⠀⠀⡇"),
		[]rune("⠀⠤⠃"),
	},
	'K': [3][]rune{
		[]rune("⢰⢀⠔"),
		[]rune("⢸⠣⡀"),
		[]rune("⠈⠀⠈"),
	},
	'k': [3][]rune{
		[]rune("⢰⠀⣀"),
		[]rune("⢸⠪⡀"),
		[]rune("⠈⠀⠈"),
	},
	'L': [3][]rune{
		[]rune("⢰⠀⠀"),
		[]rune("⢸⠀⠀"),
		[]rune("⠈⠉⠉"),
	},
	'l': [3][]rune{
		[]rune("⠀⢲⠀"),
		[]rune("⠀⢸⠀"),
		[]rune("⠀⠀⠉"),
	},
	'M': [3][]rune{
		[]rune("⢰⡀⣰"),
		[]rune("⢸⠈⢸"),
		[]rune("⠈⠀⠈"),
	},
	'm': [3][]rune{
		[]rune("⢀⣀⡀"),
		[]rune("⢸⢸⢸"),
		[]rune("⠈⠈⠈"),
	},
	'N': [3][]rune{
		[]rune("⢰⡄⢰"),
		[]rune("⢸⠘⣼"),
		[]rune("⠈⠀⠈"),
	},
	'n': [3][]rune{
		[]rune("⢀⢀⡀"),
		[]rune("⢸⠁⢸"),
		[]rune("⠈⠀⠈"),
	},
	'O': [3][]rune{
		[]rune("⢠⠒⢢"),
		[]rune("⢸⠀⢸"),
		[]rune("⠀⠉⠁"),
	},
	'o': [3][]rune{
		[]rune("⠀⣀⡀"),
		[]rune("⢸⠀⢸"),
		[]rune("⠀⠉⠁"),
	},
	'P': [3][]rune{
		[]rune("⢰⠒⢢"),
		[]rune("⢸⠒⠊"),
		[]rune("⠈⠀⠀"),
	},
	'p': [3][]rune{
		[]rune("⢀⢀⡀"),
		[]rune("⢸⡁⢸"),
		[]rune("⠸⠈⠁"),
	},
	'Q': [3][]rune{
		[]rune("⢠⠒⢢"),
		[]rune("⢸⠀⢸"),
		[]rune("⠀⠉⠑"),
	},
	'q': [3][]rune{
		[]rune("⠀⣀⢀"),
		[]rune("⢸⠀⣹"),
		[]rune("⠀⠉⠸"),
	},
	'R': [3][]rune{
		[]rune("⢰⠒⢢"),
		[]rune("⢸⠒⢎"),
		[]rune("⠈⠀⠈"),
	},
	'r': [3][]rune{
		[]rune("⢀⢀⡀"),
		[]rune("⢸⠁⠈"),
		[]rune("⠈⠀⠀"),
	},
	'S': [3][]rune{
		[]rune("⢠⠒⠢"),
		[]rune("⢀⠉⢢"),
		[]rune("⠀⠉⠁"),
	},
	's': [3][]rune{
		[]rune("⠀⣀⡀"),
		[]rune("⢈⠒⢢"),
		[]rune("⠀⠉⠁"),
	},
	'T': [3][]rune{
		[]rune("⠐⢲⠒"),
		[]rune("⠀⢸⠀"),
		[]rune("⠀⠈⠀"),
	},
	't': [3][]rune{
		[]rune("⢀⣰⣀"),
		[]rune("⠀⢸⠀"),
		[]rune("⠀⠀⠉"),
	},
	'U': [3][]rune{
		[]rune("⢰⠀⢰"),
		[]rune("⢸⠀⢸"),
		[]rune("⠀⠉⠁"),
	},
	'u': [3][]rune{
		[]rune("⢀⠀⢀"),
		[]rune("⢸⠀⣸"),
		[]rune("⠀⠉⠈"),
	},

	'V': [3][]rune{
		[]rune("⢰⠀⢰"),
		[]rune("⠈⢆⠎"),
		[]rune("⠀⠈⠀"),
	},

	'v': [3][]rune{
		[]rune("⢀⠀⢀"),
		[]rune("⠈⢆⠎"),
		[]rune("⠀⠈⠀"),
	},

	'W': [3][]rune{
		[]rune("⢰⢀⢰"),
		[]rune("⠘⡜⡜"),
		[]rune("⠀⠁⠁"),
	},

	'w': [3][]rune{
		[]rune("⢀⠀⢀"),
		[]rune("⠸⡰⡸"),
		[]rune("⠀⠁⠁"),
	},

	'X': [3][]rune{
		[]rune("⠰⡀⡰"),
		[]rune("⢀⠜⢄"),
		[]rune("⠈⠀⠈"),
	},
	'x': [3][]rune{
		[]rune("⢀⠀⢀"),
		[]rune("⠀⡱⡁"),
		[]rune("⠈⠀⠈"),
	},
	'Y': [3][]rune{
		[]rune("⠰⡀⡰"),
		[]rune("⠀⢸⠀"),
		[]rune("⠀⠈⠀"),
	},
	'y': [3][]rune{
		[]rune("⢀⠀⢀"),
		[]rune("⠈⢆⠎"),
		[]rune("⠠⠜⠀"),
	},
	'Z': [3][]rune{
		[]rune("⠐⢒⠖"),
		[]rune("⢀⠎⠀"),
		[]rune("⠈⠉⠉"),
	},
	'z': [3][]rune{
		[]rune("⢀⣀⣀"),
		[]rune("⢀⠔⠁"),
		[]rune("⠈⠉⠉"),
	},

	'_': [3][]rune{
		[]rune("⠀⠀⠀"),
		[]rune("⠀⠀⠀"),
		[]rune("⠤⠤⠤"),
	},

	' ': [3][]rune{
		[]rune("⠀⠀⠀"),
		[]rune("⠀⠀⠀"),
		[]rune("⠀⠀⠀"),
	},

	'.': [3][]rune{
		[]rune("⠀⠀⠀"),
		[]rune("⠀⢀⠀"),
		[]rune("⠀⠈⠀"),
	},

	'/': [3][]rune{
		[]rune("⠀⠀⡰"),
		[]rune("⠀⡰⠁"),
		[]rune("⠐⠁⠀"),
	},

	'\\': [3][]rune{
		[]rune("⠰⡀⠀"),
		[]rune("⠀⠱⡀"),
		[]rune("⠀⠀⠑"),
	},

	'|': [3][]rune{
		[]rune("⠀⢰⠀"),
		[]rune("⠀⢸⠀"),
		[]rune("⠀⠸⠀"),
	},
	':': [3][]rune{
		[]rune("⠀⢀⠀"),
		[]rune("⠀⢈⠀"),
		[]rune("⠀⠈⠀"),
	},

	',': [3][]rune{
		[]rune("⠀⠀⠀"),
		[]rune("⠀⢀⠀"),
		[]rune("⠀⠜⠀"),
	},

	'!': [3][]rune{
		[]rune("⠀⢰⠀"),
		[]rune("⠀⢘⠀"),
		[]rune("⠀⠈⠀"),
	},

	'?': [3][]rune{
		[]rune("⠠⠒⡄"),
		[]rune("⠀⢘⠀"),
		[]rune("⠀⠈⠀"),
	},

	'#': [3][]rune{
		[]rune("⢀⣔⣔"),
		[]rune("⢤⢧⠧"),
		[]rune("⠈⠈⠀"),
	},

	'+': [3][]rune{
		[]rune("⠀⢀⠀"),
		[]rune("⠐⢺⠒"),
		[]rune("⠀⠀⠀"),
	},

	'-': [3][]rune{
		[]rune("⠀⠀⠀"),
		[]rune("⠐⠒⠒"),
		[]rune("⠀⠀⠀"),
	},

	'@': [3][]rune{
		[]rune("⢀⠤⢄"),
		[]rune("⡇⡎⢹"),
		[]rune("⠑⠬⠍"),
	},

	'=': [3][]rune{
		[]rune("⠀⠀⠀"),
		[]rune("⠨⠭⠭"),
		[]rune("⠀⠀⠀"),
	},

	/* 	'🖤' : [3][]rune{
		[]rune("⢠⣶⣦⣠⣶⣦"),
		[]rune("⠘⢿⣿⣿⣿⠟"),
		[]rune("⠀⠀⠙⠟⠁⠀"),
	}, */

	'0': [3][]rune{
		[]rune("⢠⢒⢢"),
		[]rune("⢸⠸⢸"),
		[]rune("⠀⠉⠁"),
	},
	'1': [3][]rune{
		[]rune("⢀⢴⠀"),
		[]rune("⠀⢸⠀"),
		[]rune("⠈⠉⠉"),
	},
	'2': [3][]rune{
		[]rune("⠠⠒⢢"),
		[]rune("⠀⡠⠊"),
		[]rune("⠈⠉⠉"),
	},
	'3': [3][]rune{
		[]rune("⠠⠒⢢"),
		[]rune("⢀⠈⢱"),
		[]rune("⠀⠉⠁"),
	},
	'4': [3][]rune{
		[]rune("⠀⡠⡆"),
		[]rune("⠰⠥⡧"),
		[]rune("⠀⠀⠁"),
	},
	'5': [3][]rune{
		[]rune("⢰⠒⠒"),
		[]rune("⢈⠉⢱"),
		[]rune("⠀⠉⠁"),
	},
	'6': [3][]rune{
		[]rune("⢠⠒⠢"),
		[]rune("⢸⠉⢱"),
		[]rune("⠀⠉⠁"),
	},
	'7': [3][]rune{
		[]rune("⠐⠒⡲"),
		[]rune("⠀⡰⠁"),
		[]rune("⠀⠁⠀"),
	},
	'8': [3][]rune{
		[]rune("⢠⠒⢢"),
		[]rune("⢰⠉⢱"),
		[]rune("⠀⠉⠁"),
	},
	'9': [3][]rune{
		[]rune("⢠⠒⢢"),
		[]rune("⢈⠒⢺"),
		[]rune("⠀⠉⠁"),
	},
}
