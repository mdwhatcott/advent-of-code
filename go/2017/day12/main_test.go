package day12_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day12"
	"github.com/mdwhatcott/testing/should"
)

func TestDay12(t *testing.T) {
	part1, part2 := day12.Answers()
	should.So(t, part1, should.Equal, 115)
	should.So(t, part2, should.Equal, 221)
}
