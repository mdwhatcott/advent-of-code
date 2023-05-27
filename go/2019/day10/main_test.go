package advent_test

import (
	"testing"

	day10 "github.com/mdwhatcott/advent-of-code/go/2019/day10"
	"github.com/mdwhatcott/testing/should"
)

func TestDay10(t *testing.T) {
	should.So(t, day10.Part1(), should.Equal, 253)
	should.So(t, day10.Part2(), should.Equal, 815)
}
