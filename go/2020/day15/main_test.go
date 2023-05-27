package advent_test

import (
	"testing"

	day15 "github.com/mdwhatcott/advent-of-code/go/2020/day15"
	"github.com/mdwhatcott/testing/should"
)

func TestDay15(t *testing.T) {
	should.So(t, day15.Part1(), should.Equal, 1025)
	should.So(t, day15.Part2(), should.Equal, 129262)
}
