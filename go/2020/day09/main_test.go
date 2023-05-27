package advent_test

import (
	"testing"

	day09 "github.com/mdwhatcott/advent-of-code/go/2020/day09"
	"github.com/mdwhatcott/testing/should"
)

func TestDay09(t *testing.T) {
	should.So(t, day09.Part1(), should.Equal, 23278925)
	should.So(t, day09.Part2(), should.Equal, 4011064)
}
