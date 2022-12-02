package advent

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestPart1(t *testing.T) {
	should.So(t, Part1(), should.Equal, 7720)
}
