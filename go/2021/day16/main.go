package day16

import (
	"container/list"
	"fmt"
)

func Part1(message string) (versionSum int) {
	bits := hex2bit(message)
	//fmt.Println("BITS:", bits)
	cooked := Parse(NewBitReader(bits))
	queue := list.New()
	queue.PushFront(cooked)
	for queue.Len() > 0 {
		element := queue.Front()
		packet := element.Value.(Packet)
		//fmt.Println(packet)
		versionSum += packet.Version
		queue.Remove(element)
		for _, sub := range packet.SubPackets {
			queue.PushFront(sub)
		}
	}
	return versionSum
}

///////////////////////////////////////////////////////////////

type BitReader struct {
	source string
	cursor int
}

func NewBitReader(source string) *BitReader {
	return &BitReader{source: source}
}

func (this *BitReader) ReadInt(length int) int {
	return bit2int(this.Read(length))
}
func (this *BitReader) Read(length int) (result string) {
	result = this.source[this.cursor : this.cursor+length]
	this.cursor += length
	return result
}
func (this *BitReader) Unread() int   { return len(this.source) - this.cursor }
func (this *BitReader) Seek(to int)   { this.cursor = to }
func (this *BitReader) Position() int { return this.cursor }

///////////////////////////////////////////////////////////////

type Packet struct {
	Version    int
	Type       int
	Value      int
	SubPackets []Packet

	From int
	To   int
}

func (this Packet) String() string {
	return fmt.Sprintf(
		"Version: [%d] Type: [%d] Value: [%d] Subpackets: [%d]",
		this.Version, this.Type, this.Value, len(this.SubPackets),
	)
}

func Parse(bits *BitReader) Packet {
	packet := Packet{
		From:    bits.Position(),
		Version: bits.ReadInt(3),
		Type:    bits.ReadInt(3),
	}
	defer func() {
		packet.To = bits.Position()
	}()

	if packet.Type == 4 {
		packet.Value = parseLiteralValue(bits)
		return packet
	}

	lengthType := bit2int(bits.Read(1))
	if lengthType == 0 {
		packet.SubPackets = parseSubpacketBits(bits, bits.ReadInt(15))
	} else {
		packet.SubPackets = parseManySubPackets(bits, bits.ReadInt(11))
	}
	return packet
}

func parseLiteralValue(bits *BitReader) (value int) {
	groups := ""
	for {
		flag := bits.Read(1)
		groups += bits.Read(4)
		if flag == "0" {
			break
		}
	}
	return bit2int(groups)
}

func parseSubpacketBits(bits *BitReader, length int) (packets []Packet) {
	defer func() { _ = recover() }() // HACK! (for dealing with trailing zero padding)
	from := bits.Position()
	to := from + length
	for from <= to {
		packets = append(packets, Parse(bits))
	}
	return packets
}

func parseManySubPackets(bits *BitReader, count int) (packets []Packet) {
	defer func() { _ = recover() }() // HACK! (for dealing with trailing zero padding)
	for ; count > 0; count-- {
		packets = append(packets, Parse(bits))
	}
	return packets
}

/////////////////////////////////////////////////////////////////

func bit2int(bits string) (result int) {
	_, _ = fmt.Sscanf(bits, "%b", &result)
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
