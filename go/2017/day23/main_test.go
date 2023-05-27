package day23_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day23"
	"github.com/mdwhatcott/testing/should"
)

func TestDay23(t *testing.T) {
	should.So(t, day23.Part1(), should.Equal, 6241)
	should.So(t, day23.Part2(), should.Equal, 909)
}
