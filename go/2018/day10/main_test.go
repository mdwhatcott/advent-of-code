package day10_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day10"
	"github.com/mdwhatcott/testing/should"
)

func TestDay10(t *testing.T) {
	t.Parallel()

	should.So(t, day10.Part1(), should.Equal, "ERCXLAJL")
	should.So(t, day10.Part2(), should.Equal, 10813)
}
