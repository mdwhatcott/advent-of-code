package day02_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day02"
	"github.com/mdwhatcott/testing/should"
)

func TestDay02(t *testing.T) {
	t.Parallel()

	should.So(t, day02.Part1(), should.Equal, 5976)
	should.So(t, day02.Part2(), should.Equal, "xretqmmonskvzupalfiwhcfdb")
}
