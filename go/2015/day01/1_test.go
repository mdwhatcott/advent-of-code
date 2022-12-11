package advent

import (
	"testing"

	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

func Test1(t *testing.T) {
	input := util.InputString()
	should.So(t, endingFloor(input), should.Equal, 232)
	should.So(t, basement(input), should.Equal, 1783)
}
