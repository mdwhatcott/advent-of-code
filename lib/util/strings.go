package util

func Anagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aLetters := make(map[rune]int)
	bLetters := make(map[rune]int)

	for _, c := range a {
		aLetters[c]++
	}
	for _, c := range b {
		bLetters[c]++
	}
	for c, n := range aLetters {
		if bLetters[c] != n {
			return false
		}
	}
	for c, n := range bLetters {
		if aLetters[c] != n {
			return false
		}
	}

	return len(aLetters) == len(bLetters)
}
