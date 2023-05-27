package day04_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day04"
	"github.com/mdwhatcott/testing/should"
)

func TestDay04(t *testing.T) {
	t.Parallel()

	should.So(t, day04.Part1(), should.Equal, 35184)
	should.So(t, day04.Part2(), should.Equal, 37886)
}
