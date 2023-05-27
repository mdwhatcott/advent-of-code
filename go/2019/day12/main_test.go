package advent_test

import (
	"testing"

	day12 "github.com/mdwhatcott/advent-of-code/go/2019/day12"
	"github.com/mdwhatcott/testing/should"
)

func TestDay12(t *testing.T) {
	should.So(t, day12.Part1(), should.Equal, 12644)
	should.So(t, day12.Part2(), should.Equal, 290314621566528)
}
