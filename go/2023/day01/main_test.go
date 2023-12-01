package day01

import (
	"strings"
	"testing"
	"unicode"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/testing/should"
)

var (
	inputLines   = inputs.Read(2023, 1).Lines()
	sampleLines1 = []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	sampleLines2 = []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
)

func Test(t *testing.T) {
	should.So(t, CalibrationSum(sampleLines1, nil), should.Equal, 142)
	should.So(t, CalibrationSum(inputLines, nil), should.Equal, 55538)

	should.So(t, CalibrationSum(sampleLines2, numbers), should.Equal, 281)
	should.So(t, CalibrationSum(inputLines, numbers), should.Equal, 54875)
}

func CalibrationSum(lines []string, replacements map[string]int) (result int) {
	for _, line := range lines {
		result += CalibrationValue(line, replacements)
	}
	return result
}
func CalibrationValue(s string, replacements map[string]int) int {
	return edgeDigit(s, replacements, false)*10 +
		edgeDigit(reverse(s), replacements, true)
}
func edgeDigit(s string, replacements map[string]int, backward bool) int {
	if len(s) == 0 {
		return 0
	}
	if i := digit(s, replacements, backward); i > 0 {
		return i
	}
	return edgeDigit(s[1:], replacements, backward)
}
func digit(s string, replacements map[string]int, reversed bool) int {
	if unicode.IsDigit(rune(s[0])) {
		return int(s[0] - '0')
	}
	for key, val := range replacements {
		if reversed {
			key = reverse(key)
		}
		if strings.HasPrefix(s, key) {
			return val
		}
	}
	return 0
}
func reverse(s string) string {
	return string(funcy.Reverse([]rune(s)))
}

var numbers = map[string]int{
	//"zero":  0, // not included
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
