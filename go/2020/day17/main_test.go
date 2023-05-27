package advent_test

import (
	"testing"

	day17 "github.com/mdwhatcott/advent-of-code/go/2020/day17"
	"github.com/mdwhatcott/testing/should"
)

func TestDay17(t *testing.T) {
	should.So(t, day17.Part1(), should.Equal, 265)
	should.So(t, day17.Part2(), should.Equal, 1936)
}
