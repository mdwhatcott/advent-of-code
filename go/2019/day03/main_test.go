package advent_test

import (
	"testing"

	day03 "github.com/mdwhatcott/advent-of-code/go/2019/day03"
	"github.com/mdwhatcott/testing/should"
)

func TestDay03(t *testing.T) {
	should.So(t, day03.Part1(), should.Equal, 2050)
	should.So(t, day03.Part2(), should.Equal, 21666)
}
