package advent_test

import (
	"testing"

	day07 "github.com/mdwhatcott/advent-of-code/go/2019/day07"
	"github.com/mdwhatcott/testing/should"
)

func TestDay07(t *testing.T) {
	should.So(t, day07.Part1(), should.Equal, 75228)
	should.So(t, day07.Part2(), should.Equal, 79846026)
}
