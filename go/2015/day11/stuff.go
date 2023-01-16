package day11

import (
	"bytes"
	"fmt"
)

func Increment(current string) string {
	password := Password(current)
	password.Increment()
	return password.String()
}

type password struct {
	original string
	scratch  []byte
	i        int
}

func Password(s string) *password {
	return &password{original: s, scratch: []byte(s)}
}

func (this *password) Increment() {
	for {
		this.i = len(this.original) - 1
		this.increment()

		if this.shouldRollOver() {
			this.rollover()
		}

		if IsValid(this.scratch) {
			return
		}
	}

	panic(fmt.Sprintf("Could not increment from %s. Current value: %s", this.original, string(this.scratch)))
}

func (this *password) shouldRollOver() bool {
	return this.scratch[this.i] > 'z'
}
func (this *password) rollover() {
	for this.scratch[this.i] == 'z'+1 {
		this.scratch[this.i] = 'a'
		if this.i > 0 {
			this.i--
			this.increment()
		}
	}
	this.i++
}

func (this *password) String() string {
	return string(this.scratch)
}

func (this *password) increment() {
	this.scratch[this.i]++
	if this.scratch[this.i] == 'i' || this.scratch[this.i] == 'l' || this.scratch[this.i] == 'o' {
		this.scratch[this.i]++
	}
}

func IsValid(scratch []byte) bool {
	if bytes.ContainsAny(scratch, "ilo") {
		return false
	}
	if len(scratch) < 8 {
		return true
	}
	if !contains3Straight(scratch) {
		return false
	}
	if !contains2Pairs(scratch) {
		return false
	}
	return true
}

func contains3Straight(scratch []byte) bool {
	for x := range scratch[:len(scratch)-2] {
		s1 := scratch[x+0]
		s2 := scratch[x+1]
		s3 := scratch[x+2]
		if s1+1 == s2 && s2+1 == s3 {
			return true
		}
	}
	return false
}

func contains2Pairs(scratch []byte) bool {
	pairs := 0
	for x := 0; x <= len(scratch)-2; x++ {
		s1 := scratch[x+0]
		s2 := scratch[x+1]
		if s1 == s2 {
			pairs++
			x++
		}
	}
	return pairs >= 2
}
