package advent_test

import (
	"testing"

	day16 "github.com/mdwhatcott/advent-of-code/go/2019/day16"
	"github.com/mdwhatcott/testing/should"
)

func TestDay16Part1(t *testing.T) {
	should.So(t, day16.Part1(), should.Equal, "30369587")
}
func TestDay16Part2(t *testing.T) {
	should.So(t, day16.Part2(), should.Equal, nil)
}
