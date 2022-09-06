package advent

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestPart1Examples(t *testing.T) {
	should.So(t, part1(2020, 0, 3, 6), should.Equal, 436)
}
