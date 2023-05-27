package advent_test

import (
	"testing"

	day11 "github.com/mdwhatcott/advent-of-code/go/2020/day11"
	"github.com/mdwhatcott/testing/should"
)

func TestDay11(t *testing.T) {
	should.So(t, day11.Part1(), should.Equal, 2438)
	should.So(t, day11.Part2(), should.Equal, 2174)
}
