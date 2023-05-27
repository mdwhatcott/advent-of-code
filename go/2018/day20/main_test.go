package day20_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day20"
	"github.com/mdwhatcott/testing/should"
)

func TestDay20(t *testing.T) {
	t.Parallel()

	should.So(t, day20.Part1(), should.Equal, nil)
	should.So(t, day20.Part2(), should.Equal, nil)
}
