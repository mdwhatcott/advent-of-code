package day08_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day08"
	"github.com/mdwhatcott/testing/should"
)

func TestDay08(t *testing.T) {
	t.Parallel()

	should.So(t, day08.Part1(), should.Equal, 38567)
	should.So(t, day08.Part2(), should.Equal, 24453)
}
