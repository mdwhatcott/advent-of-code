package day09_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day09"
	"github.com/mdwhatcott/testing/should"
)

func TestDay09(t *testing.T) {
	should.So(t, day09.Part1(), should.Equal, 14212)
	should.So(t, day09.Part2(), should.Equal, 6569)
}
