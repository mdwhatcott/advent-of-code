package day20_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day20"
	"github.com/mdwhatcott/testing/should"
)

func TestDay20(t *testing.T) {
	should.So(t, day20.Part1(), should.Equal, 457)
	should.So(t, day20.Part2(), should.Equal, 448)
}
