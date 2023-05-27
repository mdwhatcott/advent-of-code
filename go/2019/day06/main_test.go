package advent_test

import (
	"testing"

	day06 "github.com/mdwhatcott/advent-of-code/go/2019/day06"
	"github.com/mdwhatcott/testing/should"
)

func TestDay06(t *testing.T) {
	should.So(t, day06.Part1(), should.Equal, 295936)
	should.So(t, day06.Part2(), should.Equal, 457)
}
