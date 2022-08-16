package day16

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

const (
	example1A   = "D2FE28"
	example1AA  = "38006F45291200"
	example1AAA = "EE00D40C823060"
	example1B   = "8A004A801A8002F478"
	example1C   = "620080001611562C8802118E34"
	example1D   = "C0015000016115A2E0802F182340"
	example1E   = "A0016C880162017C3686B18A3D4780"

	example2a = "C200B40A82"
	example2b = "04005AC33890"
	example2c = "880086C3E88112"
	example2d = "CE00C43D881120"
	example2e = "D8005AC2A8F0"
	example2f = "F600BC2D8F"
	example2g = "9C005AC2F8F0"
	example2h = "9C0141080250320F1802104A08"
)

func TestParse(t *testing.T) {
	assert.Error(t).So(
		Parse(strings.NewReader(hex2bit(example1A))),
		should.Equal,
		Packet{Version: 6, Type: 4, Value: 2021})
}
func TestPart1(t *testing.T) {
	T := assert.Error(t)
	T.So(Part1(example1A), should.Equal, 6)
	T.So(Part1(example1AA), should.Equal, 9)
	T.So(Part1(example1AAA), should.Equal, 14)
	T.So(Part1(example1B), should.Equal, 16)
	T.So(Part1(example1C), should.Equal, 12)
	T.So(Part1(example1D), should.Equal, 23)
	T.So(Part1(example1E), should.Equal, 31)
	T.So(Part1(util.InputString()), should.Equal, 927)
}
func TestPart2(t *testing.T) {
	T := assert.Error(t)
	T.So(Part2(example2a), should.Equal, 3)
	T.So(Part2(example2b), should.Equal, 54)
	T.So(Part2(example2c), should.Equal, 7)
	T.So(Part2(example2d), should.Equal, 9)
	T.So(Part2(example2e), should.Equal, 1)
	T.So(Part2(example2f), should.Equal, 0)
	T.So(Part2(example2g), should.Equal, 0)
	T.So(Part2(example2h), should.Equal, 1)
	T.So(Part2(util.InputString()), should.Equal, 1725277876501)
}
