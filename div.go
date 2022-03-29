package dots

// Both floor() and ceil()
// work correctly only on positive divider.

// floor rounds division result down, for example:
// floor(-5, 2) = -3
// floor( 5, 2) =  2
func floor(a, b int) int {
	q := a / b
	if a%b < 0 {
		q--
	}

	return q
}

// ceil rounds division result up, for example:
// floor(-5, 2) = -2
// floor( 5, 2) =  3
func ceil(a, b int) int {
	q := a / b
	if a%b > 0 {
		q++
	}

	return q
}
