package advent_test

import (
	"testing"

	day16 "github.com/mdwhatcott/advent-of-code/go/2019/day16"
	"github.com/mdwhatcott/testing/should"
)

func TestDay16(t *testing.T) {
	should.So(t, day16.Part1(), should.Equal, "30369587")
	should.So(t, day16.Part2(), should.Equal, nil)
}
