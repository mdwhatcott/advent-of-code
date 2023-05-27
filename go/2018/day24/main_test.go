package day24_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day24"
	"github.com/mdwhatcott/testing/should"
)

func TestDay24(t *testing.T) {
	t.Parallel()

	should.So(t, day24.Part1(), should.Equal, nil)
	should.So(t, day24.Part2(), should.Equal, nil)
}
