package advent_test

import (
	"testing"

	day09 "github.com/mdwhatcott/advent-of-code/go/2019/day09"
	"github.com/mdwhatcott/testing/should"
)

func TestDay09(t *testing.T) {
	should.So(t, day09.Part1(), should.Equal, []int{3742852857})
	should.So(t, day09.Part2(), should.Equal, []int{73439})
}
