package day23_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day23"
	"github.com/mdwhatcott/testing/should"
)

func TestDay23(t *testing.T) {
	t.Parallel()

	should.So(t, day23.Part1(), should.Equal, nil)
	should.So(t, day23.Part2(), should.Equal, nil)
}
