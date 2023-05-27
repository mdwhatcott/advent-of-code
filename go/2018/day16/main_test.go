package day16_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day16"
	"github.com/mdwhatcott/testing/should"
)

func TestDay16(t *testing.T) {
	t.Parallel()

	should.So(t, day16.Part1(), should.Equal, nil)
	should.So(t, day16.Part2(), should.Equal, nil)
}
