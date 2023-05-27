package advent_test

import (
	"testing"

	day01 "github.com/mdwhatcott/advent-of-code/go/2019/day01"
	"github.com/mdwhatcott/testing/should"
)

func TestDay01(t *testing.T) {
	should.So(t, day01.Part1(), should.Equal, 3563458)
	should.So(t, day01.Part2(), should.Equal, 5342292)
}
