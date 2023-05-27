package advent_test

import (
	"testing"

	day13 "github.com/mdwhatcott/advent-of-code/go/2019/day13"
	"github.com/mdwhatcott/testing/should"
)

func TestDay13(t *testing.T) {
	should.So(t, day13.Part1(), should.Equal, 226)
	should.So(t, day13.Part2(), should.Equal, 10800)
}
