package day07_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day07"
	"github.com/mdwhatcott/testing/should"
)

func TestDay07(t *testing.T) {
	part1, part2 := day07.Answers()
	should.So(t, part1, should.Equal, "dgoocsw")
	should.So(t, part2, should.Equal, 1275)
}
