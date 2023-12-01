package day01

import (
	"strings"
	"testing"
	"unicode"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
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
	return first(s, replacements)*10 + last(s, replacements)
}
func first(s string, replacements map[string]int) int {
	for x := 0; x < len(s); x++ {
		if i := digit(s[x:], replacements); i > 0 {
			return i
		}
	}
	panic("boink")
}
func last(s string, replacements map[string]int) int {
	for x := len(s) - 1; x >= 0; x-- {
		if i := digit(s[x:], replacements); i > 0 {
			return i
		}
	}
	panic("boink")
}
func digit(s string, replacements map[string]int) int {
	if unicode.IsDigit(rune(s[0])) {
		return int(s[0] - '0')
	}
	for key, val := range replacements {
		if strings.HasPrefix(s, key) {
			return val
		}
	}
	return 0
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
