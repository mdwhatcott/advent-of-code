package advent_test

import (
	"testing"

	day01 "github.com/mdwhatcott/advent-of-code/go/2020/day01"
	"github.com/mdwhatcott/testing/should"
)

func TestDay01(t *testing.T) {
	should.So(t, day01.Part1(), should.Equal, 1015476)
	should.So(t, day01.Part2(), should.Equal, 200878544)
}
