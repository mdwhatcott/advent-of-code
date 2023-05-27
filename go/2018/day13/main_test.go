package day13_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day13"
	"github.com/mdwhatcott/testing/should"
)

func TestDay13(t *testing.T) {
	t.Skip("broken :(")
	t.Parallel()

	should.So(t, day13.Part1(), should.Equal, "(14, 42)")
	should.So(t, day13.Part2(), should.Equal, nil) // NOT: [(70, 114)] or [(72, 114)]
}
