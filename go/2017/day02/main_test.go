package day02_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day02"
	"github.com/mdwhatcott/testing/should"
)

func TestDay02(t *testing.T) {
	should.So(t, day02.Part1(), should.Equal, 34925)
	should.So(t, day02.Part2(), should.Equal, 221)
}
