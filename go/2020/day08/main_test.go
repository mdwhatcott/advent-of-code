package advent_test

import (
	"testing"

	day08 "github.com/mdwhatcott/advent-of-code/go/2020/day08"
	"github.com/mdwhatcott/testing/should"
)

func TestDay08(t *testing.T) {
	should.So(t, day08.Part1(), should.Equal, 1832)
	should.So(t, day08.Part2(), should.Equal, 662)
}
