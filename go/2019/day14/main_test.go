package advent_test

import (
	"testing"

	day14 "github.com/mdwhatcott/advent-of-code/go/2019/day14"
	"github.com/mdwhatcott/testing/should"
)

func TestDay14(t *testing.T) {
	should.So(t, day14.Part1(), should.Equal, 198984)
	should.So(t, day14.Part2(), should.Equal, 7659732)
}
