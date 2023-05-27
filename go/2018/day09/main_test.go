package day09_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day09"
	"github.com/mdwhatcott/testing/should"
)

func TestDay09(t *testing.T) {
	t.Parallel()

	should.So(t, day09.Part1(), should.Equal, 398730)
	should.So(t, day09.Part2(), should.Equal, 3349635509)
}
