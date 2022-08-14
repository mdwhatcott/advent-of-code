package day16

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

const (
	exampleA   = "D2FE28"
	exampleAA  = "38006F45291200"
	exampleAAA = "EE00D40C823060"
	exampleB   = "8A004A801A8002F478"
	exampleC   = "620080001611562C8802118E34"
	exampleD   = "C0015000016115A2E0802F182340"
	exampleE   = "A0016C880162017C3686B18A3D4780"
)

func TestParse(t *testing.T) {
	T := assert.Error(t)

	packet := Parse(NewBitReader(hex2bit(exampleA)))
	T.So(packet, should.Equal, Packet{
		Version: 6,
		Type:    4,
		Value:   2021,
	})
}
func TestPart1(t *testing.T) {
	T := assert.Error(t)
	T.So(Part1(exampleA), should.Equal, 6)
	T.So(Part1(exampleAA), should.Equal, 9)
	T.So(Part1(exampleAAA), should.Equal, 14)
	T.So(Part1(exampleB), should.Equal, 16)
	T.So(Part1(exampleC), should.Equal, 12)
	T.So(Part1(exampleD), should.Equal, 23)
	T.So(Part1(exampleE), should.Equal, 31)
	T.So(Part1(util.InputString()), should.Equal, 927)
}
