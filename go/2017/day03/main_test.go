package day03_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day03"
	"github.com/mdwhatcott/testing/should"
)

func TestDay03(t *testing.T) {
	should.So(t, day03.Part1(), should.Equal, 438)
	should.So(t, day03.Part2(), should.Equal, 266330)
}
