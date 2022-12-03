package day03

import (
	"unicode"

	"github.com/mdwhatcott/go-collections/set"
)

func Part1(lines []string) (result int) {
	for _, line := range lines {
		chars := []rune(line)
		a := set.From(chars[:len(chars)/2]...)
		b := set.From(chars[len(chars)/2:]...)
		letter := a.Intersection(b).Slice()[0]
		result += priority(letter)
	}
	return result
}
func Part2(lines []string) (result int) {
	for x := 0; x < len(lines); x += 3 {
		a := set.From([]rune(lines[x+0])...)
		b := set.From([]rune(lines[x+1])...)
		c := set.From([]rune(lines[x+2])...)
		letter := a.Intersection(b).Intersection(c).Slice()[0]
		result += priority(letter)
	}
	return result
}
func priority(letter rune) int {
	if unicode.IsLower(letter) {
		return int(letter-'a') + 1
	} else {
		return int(letter-'A') + 1 + 26
	}
}
