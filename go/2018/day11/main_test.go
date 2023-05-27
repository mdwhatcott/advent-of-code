package day11_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day11"
	"github.com/mdwhatcott/testing/should"
)

func TestDay11(t *testing.T) {
	t.Parallel()

	should.So(t, day11.Part1(), should.Equal, "235,18")
	should.So(t, day11.Part2(), should.Equal, "236,227,12")
}
