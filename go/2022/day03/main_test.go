package day03

import (
	"testing"

	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

var sampleLines = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func Test(t *testing.T) {
	inputLines := util.InputLines()

	should.So(t, Part1(sampleLines), should.Equal, 157)
	should.So(t, Part1(inputLines), should.Equal, 8018)

	should.So(t, Part2(sampleLines), should.Equal, 70)
	should.So(t, Part2(inputLines), should.Equal, 2518)
}
