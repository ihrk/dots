package dots

const (
	baseRune = '\u2800'
	runeSize = 3
)

const (
	blockWidth  = 2
	blockHeight = 4
	blockSize   = blockWidth * blockHeight
)

/*
Bit positions:
┏━━━┳━━━┓
┃ 0 ┃ 3 ┃
┣━━━╋━━━┫
┃ 1 ┃ 4 ┃
┣━━━╋━━━┫
┃ 2 ┃ 5 ┃
┣━━━╋━━━┫
┃ 6 ┃ 7 ┃
┗━━━┻━━━┛
*/
var bitPos = [blockSize]int{0, 3, 1, 4, 2, 5, 6, 7}

type CodePoint uint8

func CpToRune(cp CodePoint) rune {
	return baseRune + rune(cp)
}

func CpFlipH(cp CodePoint) CodePoint {
	h0 := cp >> 7 << 6
	h1 := cp >> 6 << 7

	b0 := cp << 2 >> 5
	b1 := cp << 5 >> 2

	return h0 | h1 | b0 | b1
}

func CpFlipV(cp CodePoint) CodePoint {
	b0 := (cp & (1 << 0)) << 6
	b6 := (cp & (1 << 6)) >> 6

	b1 := (cp & (1 << 1)) << 1
	b2 := (cp & (1 << 2)) >> 1

	b3 := (cp & (1 << 3)) << 4
	b7 := (cp & (1 << 7)) >> 4

	b4 := (cp & (1 << 4)) << 1
	b5 := (cp & (1 << 5)) >> 1

	return b0 | b1 | b2 | b3 | b4 | b5 | b6 | b7
}

//UB on random rune
func RuneToCp(r rune) CodePoint {
	return CodePoint(r - baseRune)
}
