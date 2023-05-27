package day11_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day11"
	"github.com/mdwhatcott/testing/should"
)

func TestDay11(t *testing.T) {
	part1, part2 := day11.Answers()
	should.So(t, part1, should.Equal, 707)
	should.So(t, part2, should.Equal, 1490)
}
