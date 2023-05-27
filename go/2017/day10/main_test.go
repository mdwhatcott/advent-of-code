package day10_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day10"
	"github.com/mdwhatcott/testing/should"
)

func TestDay10(t *testing.T) {
	should.So(t, day10.Part1(), should.Equal, 6952)
	should.So(t, day10.Part2(), should.Equal, "28e7c4360520718a5dc811d3942cf1fd")
}
