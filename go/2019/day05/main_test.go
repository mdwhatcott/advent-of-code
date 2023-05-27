package advent_test

import (
	"testing"

	day05 "github.com/mdwhatcott/advent-of-code/go/2019/day05"
	"github.com/mdwhatcott/testing/should"
)

func TestDay05(t *testing.T) {
	should.So(t, day05.Part1(), should.Equal, 16348437)
	should.So(t, day05.Part2(), should.Equal, 6959377)
}
