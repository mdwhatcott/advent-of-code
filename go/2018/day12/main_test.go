package day12_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day12"
	"github.com/mdwhatcott/testing/should"
)

func TestDay12(t *testing.T) {
	t.Parallel()

	should.So(t, day12.Part1(), should.Equal, 3494)
	should.So(t, day12.Part2(), should.Equal, 2850000002454)
}
