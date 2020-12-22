package advent

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestPart1Examples(t *testing.T) {
	assertions.New(t).So(part1(2020, 0, 3, 6), should.Equal, 436)
}
