package day19_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day19"
	"github.com/mdwhatcott/testing/should"
)

func TestDay19(t *testing.T) {
	t.Parallel()

	should.So(t, day19.Part1(), should.Equal, nil)
	should.So(t, day19.Part2(), should.Equal, nil)
}
