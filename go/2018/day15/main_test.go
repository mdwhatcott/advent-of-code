package day15_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day15"
	"github.com/mdwhatcott/testing/should"
)

func TestDay15(t *testing.T) {
	t.Parallel()

	should.So(t, day15.Part1(), should.Equal, nil)
	should.So(t, day15.Part2(), should.Equal, nil)
}
