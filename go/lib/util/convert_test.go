package util

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"
)

func TestBinaryHammingWeight(t *testing.T) {
	assert.So(t, BinaryHammingWeight(0), should.Equal, 0)
	assert.So(t, BinaryHammingWeight(1), should.Equal, 1)
	assert.So(t, BinaryHammingWeight(2), should.Equal, 1)
	assert.So(t, BinaryHammingWeight(3), should.Equal, 2)
	assert.So(t, BinaryHammingWeight(255), should.Equal, 8)
}

func TestEncodeBinary(t *testing.T) {
	assert.So(t, EncodeBinary(0), should.Equal, "00000000")
	assert.So(t, EncodeBinary(1), should.Equal, "00000001")
	assert.So(t, EncodeBinary(255), should.Equal, "11111111")
}
