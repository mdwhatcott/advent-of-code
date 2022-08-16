package day16

import (
	"fmt"
	"io"
	"strings"
)

func Part1(message string) (versionSum int) {
	return Parse(strings.NewReader(hex2bit(message))).VersionSum()
}

func Part2(message string) int {
	return Parse(strings.NewReader(hex2bit(message))).Eval()
}

type Packet struct {
	Version    int
	Type       int
	Value      int
	SubPackets []Packet
}

func (this Packet) VersionSum() (result int) {
	for _, sub := range this.SubPackets {
		result += sub.VersionSum()
	}
	return result + this.Version
}

func (this Packet) Eval() int {
	switch this.Type {
	case 0:
		return this.sum()
	case 1:
		return this.product()
	case 2:
		return this.minimum()
	case 3:
		return this.maximum()
	case 4:
		return this.Value
	case 5:
		return this.greaterThan()
	case 6:
		return this.lessThan()
	case 7:
		return this.equalTo()
	}
	panic("NOPE")
}
func (this Packet) sum() (result int) {
	for _, value := range this.subValues() {
		result += value
	}
	return result
}
func (this Packet) product() (result int) {
	result = 1
	for _, value := range this.subValues() {
		result *= value
	}
	return result
}
func (this Packet) minimum() (result int) {
	result = 0xFFFFFFFF
	for _, value := range this.subValues() {
		if value < result {
			result = value
		}
	}
	return result
}
func (this Packet) maximum() (result int) {
	result = -0xFFFFFFFF
	for _, value := range this.subValues() {
		if value > result {
			result = value
		}
	}
	return result
}
func (this Packet) subValues() (results []int) {
	for _, sub := range this.SubPackets {
		results = append(results, sub.Eval())
	}
	return results
}
func (this Packet) greaterThan() int {
	return bool2int(this.SubPackets[0].Eval() > this.SubPackets[1].Eval())
}
func (this Packet) lessThan() int {
	return bool2int(this.SubPackets[0].Eval() < this.SubPackets[1].Eval())
}
func (this Packet) equalTo() int {
	return bool2int(this.SubPackets[0].Eval() == this.SubPackets[1].Eval())
}

func Parse(bits io.Reader) Packet {
	packet := Packet{
		Version: read(bits, 3),
		Type:    read(bits, 3),
	}

	if packet.Type == 4 {
		packet.Value = parseLiteralValue(bits)
		return packet
	}

	lengthType := read(bits, 1)
	if lengthType == 0 {
		packet.SubPackets = parseSubpacketBits(chunk(bits, 15))
	} else {
		packet.SubPackets = parseManySubPackets(bits, read(bits, 11))
	}
	return packet
}

func parseLiteralValue(bits io.Reader) (value int) {
	for {
		flag := read(bits, 1)
		value = value<<4 | read(bits, 4)
		if flag == 0 {
			break
		}
	}
	return value
}

func parseSubpacketBits(bits io.Reader) (packets []Packet) {
	defer func() { _ = recover() }() // HACK! (for dealing with trailing zero padding)
	for {
		packets = append(packets, Parse(bits))
	}
}

func parseManySubPackets(bits io.Reader, count int) (packets []Packet) {
	for ; count > 0; count-- {
		packets = append(packets, Parse(bits))
	}
	return packets
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func chunk(bits io.Reader, length int) io.Reader {
	b := make([]byte, read(bits, length))
	_, _ = bits.Read(b)
	return strings.NewReader(string(b))
}

func read(reader io.Reader, n int) (result int) {
	b := make([]byte, n)
	n, err := reader.Read(b)
	if n == 0 && err == io.EOF {
		panic(err) // HACK! (for dealing with trailing zero padding)
	}
	_, _ = fmt.Sscanf(string(b), "%b", &result)
	return result
}

func hex2bit(message string) (result string) {
	for _, m := range message {
		result += hexes[m]
	}
	return result
}

// Use this map instead of the encoding/hex package because
// it ensures that leading zeros in decoded binary strings
// are never truncated.
var hexes = map[rune]string{
	'0': "0000", '1': "0001", '2': "0010", '3': "0011",
	'4': "0100", '5': "0101", '6': "0110", '7': "0111",
	'8': "1000", '9': "1001", 'A': "1010", 'B': "1011",
	'C': "1100", 'D': "1101", 'E': "1110", 'F': "1111",
}
