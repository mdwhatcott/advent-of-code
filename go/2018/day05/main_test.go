package day05_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day05"
	"github.com/mdwhatcott/testing/should"
)

func TestDay05(t *testing.T) {
	t.Parallel()

	should.So(t, day05.Part1(), should.Equal, 11264)
	should.So(t, day05.Part2(), should.Equal, 4552)
}
