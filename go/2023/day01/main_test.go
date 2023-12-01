package day01

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/must/strconvmust"
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
	should.So(t, Part1(sampleLines1), should.Equal, 142)
	should.So(t, Part1(inputLines), should.Equal, 55538)

	should.So(t, Part2(sampleLines2), should.Equal, 281)
	should.So(t, Part2(inputLines), should.Equal, 54875)
}

func Part1(lines []string) (result int) {
	for _, line := range lines {
		var first, second string
		for _, char := range line {
			if unicode.IsDigit(char) {
				if first == "" {
					first = string(char)
				}
				second = string(char)
			}
		}
		result += strconvmust.Atoi(first + second)
	}
	return result
}

func Part2(lines []string) (result int) {
	for _, line := range lines {
		var first, second string
		for c, char := range line {
			if unicode.IsDigit(char) {
				if first == "" {
					first = string(char)
				}
				second = string(char)
				continue
			}
			sub := line[c:]
			digit, ok := extractDigit(sub)
			if !ok {
				continue
			}
			if first == "" {
				first = fmt.Sprint(digit)
			}
			second = fmt.Sprint(digit)
		}

		result += strconvmust.Atoi(first + second)
	}
	return result
}

func extractDigit(s string) (int, bool) {
	for key, val := range numbers {
		if strings.HasPrefix(s, key) {
			return val, true
		}
	}
	return 0, false
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
