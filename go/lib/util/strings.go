package util

import "strings"

func RemoveAll(source string, removes ...string) string {
	for _, remove := range removes {
		source = strings.ReplaceAll(source, remove, "")
	}
	return source
}

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

// Levenshtein Distance Algorithm (Source: https://rosettacode.org/wiki/Levenshtein_distance#Go)
func Levenshtein(s, t string) int {
	if s == "" {
		return len(t)
	}
	if t == "" {
		return len(s)
	}
	if s[0] == t[0] {
		return Levenshtein(s[1:], t[1:])
	}
	a := Levenshtein(s[1:], t[1:])
	b := Levenshtein(s, t[1:])
	c := Levenshtein(s[1:], t)
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	return a + 1
}
