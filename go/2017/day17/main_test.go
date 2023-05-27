package day17_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day17"
	"github.com/mdwhatcott/testing/should"
)

func TestDay17(t *testing.T) {
	should.So(t, day17.Part1(), should.Equal, 777)
	should.So(t, day17.Part2(), should.Equal, 39289581)
}
