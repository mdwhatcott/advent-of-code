package util

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"
)

func TestBinaryHammingWeight(t *testing.T) {
	a := assert.Error(t)
	a.So(BinaryHammingWeight(0), should.Equal, 0)
	a.So(BinaryHammingWeight(1), should.Equal, 1)
	a.So(BinaryHammingWeight(255), should.Equal, 8)
}

func TestEncodeBinary(t *testing.T) {
	a := assert.Error(t)
	a.So(EncodeBinary(0), should.Equal, "00000000")
	a.So(EncodeBinary(1), should.Equal, "00000001")
	a.So(EncodeBinary(255), should.Equal, "11111111")
}
