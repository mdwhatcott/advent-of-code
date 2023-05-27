package day22_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day22"
	"github.com/mdwhatcott/testing/should"
)

func TestDay22(t *testing.T) {
	t.Parallel()

	should.So(t, day22.Part1(), should.Equal, nil)
	should.So(t, day22.Part2(), should.Equal, nil)
}
