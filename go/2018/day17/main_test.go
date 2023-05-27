package day17_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day17"
	"github.com/mdwhatcott/testing/should"
)

func TestDay17(t *testing.T) {
	t.Parallel()

	should.So(t, day17.Part1(), should.Equal, nil)
	should.So(t, day17.Part2(), should.Equal, nil)
}
