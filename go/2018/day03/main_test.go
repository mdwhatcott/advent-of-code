package day03_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day03"
	"github.com/mdwhatcott/testing/should"
)

func TestDay03(t *testing.T) {
	t.Parallel()

	should.So(t, day03.Part1(), should.Equal, 111266)
	should.So(t, day03.Part2(), should.Equal, 266)
}
