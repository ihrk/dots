package dots

const (
	baseRune = '\u2800'
	baseSize = 3
)

type CodePoint uint8

func (cp CodePoint) Rune() rune {
	return baseRune + rune(cp)
}

func (cp CodePoint) revX() CodePoint {
	h0 := cp >> 7 << 6
	h1 := cp << 1 >> 7 << 7
	b0 := cp << 2 >> 5
	b1 := cp << 5 >> 2
	return h0 | h1 | b0 | b1
}

func (cp CodePoint) revY() CodePoint {
	b0 := (cp & (1 << 0)) << 6
	b1 := (cp & (1 << 1)) << 1
	b2 := (cp & (1 << 2)) >> 1
	b3 := (cp & (1 << 3)) << 4
	b4 := (cp & (1 << 4)) << 1
	b5 := (cp & (1 << 5)) >> 1
	b6 := (cp & (1 << 6)) >> 6
	b7 := (cp & (1 << 7)) >> 4
	return b0 | b1 | b2 | b3 | b4 | b5 | b6 | b7
}

//UB on random rune
func FromRune(r rune) CodePoint {
	return CodePoint(r - baseRune)
}

func XOR(old CodePoint, new CodePoint) CodePoint {
	return old ^ new
}

func OR(old CodePoint, new CodePoint) CodePoint {
	return old | new
}

func AND(old CodePoint, new CodePoint) CodePoint {
	return old & new
}

func ANDNOT(old CodePoint, new CodePoint) CodePoint {
	return old &^ new
}

func NEWONLY(old CodePoint, new CodePoint) CodePoint {
	return new
}
