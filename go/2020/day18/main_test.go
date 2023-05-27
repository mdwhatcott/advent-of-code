package advent_test

import (
	"testing"

	day18 "github.com/mdwhatcott/advent-of-code/go/2020/day18"
	"github.com/mdwhatcott/testing/should"
)

func TestDay18(t *testing.T) {
	should.So(t, day18.Part1(), should.Equal, 6923486965641)
	should.So(t, day18.Part2(), should.Equal, 70722650566361)
}
