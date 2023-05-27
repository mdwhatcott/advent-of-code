package day21_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day21"
	"github.com/mdwhatcott/testing/should"
)

func TestDay21(t *testing.T) {
	should.So(t, day21.Part1(), should.Equal, 110)
	should.So(t, day21.Part2(), should.Equal, 1277716)
}
