package util

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestBinaryHammingWeight(t *testing.T) {
	should.So(t, BinaryHammingWeight(0), should.Equal, 0)
	should.So(t, BinaryHammingWeight(1), should.Equal, 1)
	should.So(t, BinaryHammingWeight(2), should.Equal, 1)
	should.So(t, BinaryHammingWeight(3), should.Equal, 2)
	should.So(t, BinaryHammingWeight(255), should.Equal, 8)
}

func TestEncodeBinary(t *testing.T) {
	should.So(t, EncodeBinary(0), should.Equal, "00000000")
	should.So(t, EncodeBinary(1), should.Equal, "00000001")
	should.So(t, EncodeBinary(255), should.Equal, "11111111")
}
