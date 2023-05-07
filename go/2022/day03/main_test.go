package day03

import (
	"testing"
	"unicode"

	"github.com/mdwhatcott/go-collections/set"
	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

var (
	inputLines  = util.InputLines()
	sampleLines = []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
)

func TestDay03(t *testing.T) {
	should.So(t, Part1(sampleLines), should.Equal, 157)
	should.So(t, Part1(inputLines), should.Equal, 8018)

	should.So(t, Part2(sampleLines), should.Equal, 70)
	should.So(t, Part2(inputLines), should.Equal, 2518)
}
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
