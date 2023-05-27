package day14_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day14"
	"github.com/mdwhatcott/testing/should"
)

func TestDay14(t *testing.T) {
	t.Parallel()

	should.So(t, day14.Part1(), should.Equal, nil)
	should.So(t, day14.Part2(), should.Equal, nil)
}
