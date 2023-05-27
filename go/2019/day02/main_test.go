package advent_test

import (
	"testing"

	day02 "github.com/mdwhatcott/advent-of-code/go/2019/day02"
	"github.com/mdwhatcott/testing/should"
)

func TestDay02(t *testing.T) {
	should.So(t, day02.Part1(), should.Equal, 3101878)
	should.So(t, day02.Part2(), should.Equal, 8444)
}
