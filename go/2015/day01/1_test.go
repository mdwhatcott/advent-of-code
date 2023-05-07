package advent

import (
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func Test1(t *testing.T) {
	input := util.InputString()
	should.So(t, endingFloor(input), should.Equal, 232)
	should.So(t, basement(input), should.Equal, 1783)
}
