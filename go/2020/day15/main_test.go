package advent

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"
)

func TestPart1Examples(t *testing.T) {
	assert.Error(t).So(part1(2020, 0, 3, 6), should.Equal, 436)
}
