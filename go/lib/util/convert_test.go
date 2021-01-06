package util

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestBinaryHammingWeight(t *testing.T) {
	assert := assertions.New(t)
	assert.So(BinaryHammingWeight(0), should.Equal, 0)
	assert.So(BinaryHammingWeight(1), should.Equal, 1)
	assert.So(BinaryHammingWeight(255), should.Equal, 8)
}

func TestEncodeBinary(t *testing.T) {
	assert := assertions.New(t)
	assert.So(EncodeBinary(0), should.Equal, "00000000")
	assert.So(EncodeBinary(1), should.Equal, "00000001")
	assert.So(EncodeBinary(255), should.Equal, "11111111")
}
