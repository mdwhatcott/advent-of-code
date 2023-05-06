package day14

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestEncodeBinary(t *testing.T) {
	should.So(t, encodeBinary(0), should.Equal, "00000000")
	should.So(t, encodeBinary(1), should.Equal, "00000001")
	should.So(t, encodeBinary(255), should.Equal, "11111111")
}
