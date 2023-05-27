package day16_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2017/day16"
	"github.com/mdwhatcott/testing/should"
)

func TestDay16(t *testing.T) {
	should.So(t, day16.Part1(), should.Equal, "giadhmkpcnbfjelo")
	should.So(t, day16.Part2(), should.Equal, "njfgilbkcoemhpad")
}
