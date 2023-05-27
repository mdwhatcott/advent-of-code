package advent_test

import (
	"testing"

	day04 "github.com/mdwhatcott/advent-of-code/go/2020/day04"
	"github.com/mdwhatcott/testing/should"
)

func TestDay04(t *testing.T) {
	should.So(t, day04.Part1(), should.Equal, 264)
	should.So(t, day04.Part2(), should.Equal, 224)
}
