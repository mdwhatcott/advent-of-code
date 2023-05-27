package day25_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day25"
	"github.com/mdwhatcott/testing/should"
)

func TestDay25(t *testing.T) {
	t.Parallel()

	should.So(t, day25.Part1(), should.Equal, nil)
	should.So(t, day25.Part2(), should.Equal, nil)
}
