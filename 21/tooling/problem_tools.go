package tooling

func FindLetter(s *[]rune, letter string) int {
	for i, ch := range *s {
		if string(ch) == letter {
			return i
		}
	}
	return -1
}

func ReversePortion(s *[]rune, start, end int) {
	if start > end {
		start, end = end, start
	}
	if start == end {
		return
	}
	for i := 0; i < (end-start+1)/2; i++ {
		newStart, newEnd := start+i, end-i
		(*s)[newStart], (*s)[newEnd] = (*s)[newEnd], (*s)[newStart]
	}
}

func Rotate(s *[]rune, amt, start, end int, rotLeft bool) {
	if start > end {
		start, end = end, start
	}
	rotDistance := end - start + 1
	amt = amt % rotDistance
	if rotDistance == 1 || amt == 0 {
		return
	}

	replacement := make([]rune, rotDistance)

	if !rotLeft { // then go right
		amt = Mod(-1*amt, rotDistance)
	}

	for i := range replacement {
		position := (i + amt) % rotDistance
		replacement[i] = (*s)[start+position]
	}

	for i := range replacement {
		(*s)[start+i] = replacement[i]
	}
}

func SwapIdx(s *[]rune, a, b int) {
	(*s)[a], (*s)[b] = (*s)[b], (*s)[a]
}
