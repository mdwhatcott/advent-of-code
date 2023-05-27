package day06_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day06"
	"github.com/mdwhatcott/testing/should"
)

func TestDay06(t *testing.T) {
	t.Parallel()

	should.So(t, day06.Part1(), should.Equal, 3604)
	should.So(t, day06.Part2(), should.Equal, 46563)
}
