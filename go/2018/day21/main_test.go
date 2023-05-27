package day21_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day21"
	"github.com/mdwhatcott/testing/should"
)

func TestDay21(t *testing.T) {
	t.Parallel()

	should.So(t, day21.Part1(), should.Equal, nil)
	should.So(t, day21.Part2(), should.Equal, nil)
}
