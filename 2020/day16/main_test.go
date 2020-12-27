package advent

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestParseFieldDefinition(t *testing.T) {
	parsed := ParseFieldDefinition("departure location: 25-568 or 594-957")
	assertions.New(t).So(parsed.Name, should.Equal, "departure location")
}
