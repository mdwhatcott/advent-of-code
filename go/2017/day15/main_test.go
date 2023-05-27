package day15_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day15"
	"github.com/mdwhatcott/testing/should"
)

func TestDay15(t *testing.T) {
	should.So(t, day15.Part1(), should.Equal, 592)
	should.So(t, day15.Part2(), should.Equal, 320)
}
