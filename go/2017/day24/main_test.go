package day24_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day24"
	"github.com/mdwhatcott/testing/should"
)

func TestDay24(t *testing.T) {
	should.So(t, day24.Part1(), should.Equal, 1511)
	should.So(t, day24.Part2(), should.Equal, 1471)
}
