package knot

import "encoding/hex"

func HashString(input string) string {
	return hex.EncodeToString(KnotHash([]byte(input)))
}

func KnotHash(input []byte) []byte {
	var salt = []byte{17, 31, 73, 47, 23}
	lengths := append(input, salt...)
	loop := NewLoop()
	for x := 0; x < 64; x++ {
		loop.TwistAll(lengths)
	}
	return loop.Digest()
}

type Loop struct {
	list []byte
	i    int
	skip int
}

func NewLoop() *Loop {
	var ring []byte
	for x := 0; x < 256; x++ {
		ring = append(ring, byte(x))
	}
	return &Loop{list: ring}
}

func (this *Loop) FirstTwo() (byte, byte) {
	return this.list[0], this.list[1]
}

func (this *Loop) TwistAll(lengths []byte) {
	for _, length := range lengths {
		this.Twist(int(length))
	}
}

func (this *Loop) Twist(length int) {
	this.Replace(this.ReverseSlice(length))
	this.Advance(length)
}

func (this *Loop) ReverseSlice(length int) (slice []byte) {
	for x := length - 1; x >= 0; x-- {
		slice = append(slice, this.list[this.offset(x)])
	}
	return slice
}

func (this *Loop) Replace(slice []byte) {
	for x := 0; x < len(slice); x++ {
		i := this.offset(x)
		this.list[i] = slice[x]
	}
}

func (this *Loop) Advance(length int) {
	this.i = this.offset(length + this.skip)
	this.skip++
}

func (this *Loop) safe(i int) int {
	return i % len(this.list)
}
func (this *Loop) offset(i int) int {
	return this.safe(this.i + i)
}

func (this *Loop) Digest() []byte {
	digest := make([]byte, 16)
	for i, x := range this.list {
		digest[i/16] ^= x
	}
	return digest
}
