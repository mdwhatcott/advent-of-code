package advent_test

import (
	"testing"

	day17 "github.com/mdwhatcott/advent-of-code/go/2019/day17"
	"github.com/mdwhatcott/testing/should"
)

func TestDay17(t *testing.T) {
	should.So(t, day17.Part1(), should.Equal, 7720)
	should.So(t, day17.Part2(), should.Equal, nil)
}
