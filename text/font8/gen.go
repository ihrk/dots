package font8

func (f *Font8) LoadChar(r rune) {
	if int(r) > len(chars) {
		r = ' '
	}
	f.src = chars[r]
}
