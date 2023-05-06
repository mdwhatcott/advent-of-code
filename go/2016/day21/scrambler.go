package main

import (
	"bytes"
	"strings"

	"advent/lib/parse"
)

type Scrambler struct {
	working      string
	instructions []func(string) string
}

func NewUnscrambler(instructions []string) *Scrambler {
	scrambler := &Scrambler{}
	for x := len(instructions) - 1; x >= 0; x-- {
		scrambler.insertInstruction(instructions[x], strings.Fields(instructions[x]), true)
	}
	return scrambler
}

func NewScrambler(instructions []string) *Scrambler {
	scrambler := &Scrambler{}
	for _, line := range instructions {
		scrambler.insertInstruction(line, strings.Fields(line), false)
	}
	return scrambler
}

func (this *Scrambler) insertInstruction(line string, fields []string, invert bool) {
	switch {

	case strings.HasPrefix(line, "move"):
		this.insertMove(fields, invert)

	case strings.HasPrefix(line, "swap position"):
		this.insertSwapPosition(fields, invert)

	case strings.HasPrefix(line, "swap letter"):
		this.insertSwapLetters(fields, invert)

	case strings.HasPrefix(line, "reverse"):
		this.insertReverse(fields, invert)

	case strings.HasPrefix(line, "rotate based"):
		this.insertRotateByChar(fields, invert)

	case strings.HasPrefix(line, "rotate left"):
		this.insertRotateLeft(fields, invert)

	case strings.HasPrefix(line, "rotate right"):
		this.insertRotateRight(fields, invert)
	}
}

func (this *Scrambler) insertMove(fields []string, invert bool) {
	this.insert(func(s string) string {
		x, y := parse.Int(fields[2]), parse.Int(fields[5])
		if invert {
			return Move(s, y, x)
		} else {
			return Move(s, x, y)
		}
	})
}
func (this *Scrambler) insertSwapPosition(fields []string, invert bool) {
	this.insert(func(s string) string {
		return SwapPositions(s, parse.Int(fields[2]), parse.Int(fields[5]))
	})
}
func (this *Scrambler) insertSwapLetters(fields []string, invert bool) {
	this.insert(func(s string) string {
		return SwapLetters(s, fields[2], fields[5])
	})
}
func (this *Scrambler) insertReverse(fields []string, invert bool) {
	this.insert(func(s string) string {
		return ReverseRange(s, parse.Int(fields[2]), parse.Int(fields[4]))
	})
}
func (this *Scrambler) insertRotateByChar(fields []string, invert bool) {
	this.insert(func(s string) string {
		char := fields[6]
		index := strings.Index(s, char)

		if invert { // See test cases for the interpretation of these instructions.
			switch index { // TODO: there's probably a more formulaic way to do this...
			case 0:
				return Rotate(s, 7)
			case 1:
				return Rotate(s, -1)
			case 2:
				return Rotate(s, 2)
			case 3:
				return Rotate(s, -2)
			case 4:
				return Rotate(s, 1)
			case 5:
				return Rotate(s, -3)
			case 6:
				return Rotate(s, 0) // return s
			case 7:
				return Rotate(s, 4)
			}
		}

		return RotateByChar(s, fields[6])
	})
}
func (this *Scrambler) insertRotateLeft(fields []string, invert bool) {
	this.insert(func(s string) string {
		x := parse.Int(fields[2])
		if invert {
			return Rotate(s, x)
		} else {
			return Rotate(s, -x)
		}
	})
}
func (this *Scrambler) insertRotateRight(fields []string, invert bool) {
	this.insert(func(s string) string {
		x := parse.Int(fields[2])
		if invert {
			return Rotate(s, -x)
		} else {
			return Rotate(s, x)
		}
	})
}

func (this *Scrambler) insert(instruction func(string) string) {
	this.instructions = append(this.instructions, instruction)
}

func (this *Scrambler) Process(password string) string {
	this.working = password
	for _, i := range this.instructions {
		this.working = i(this.working)
	}
	return this.working
}

func SwapLetters(subject string, x, y string) string {
	var b bytes.Buffer
	for i := range subject {
		if c := subject[i : i+1]; c == x {
			b.WriteString(y)
		} else if c == y {
			b.WriteString(x)
		} else {
			b.WriteString(c)
		}
	}
	return b.String()
}

func SwapPositions(subject string, x, y int) string {
	var b bytes.Buffer
	for i := range subject {
		if i == x {
			b.WriteByte(subject[y])
		} else if i == y {
			b.WriteByte(subject[x])
		} else {
			b.WriteByte(subject[i])
		}
	}
	return b.String()
}

func ReverseRange(subject string, start, stop int) string {
	var b bytes.Buffer
	if start > 0 {
		b.WriteString(subject[:start])
	}
	for x := stop; x >= start; x-- {
		b.WriteByte(subject[x])
	}
	if stop < len(subject)-1 {
		b.WriteString(subject[stop+1:])
	}
	return b.String()
}

func Rotate(subject string, rotation int) string {
	if rotation%len(subject) == 0 {
		return subject
	}
	var b bytes.Buffer
	for x := 0; x < len(subject); x++ {
		i := x - rotation
		if i < 0 {
			i += len(subject)
		} else if i >= len(subject) {
			i -= len(subject)
		}
		b.WriteByte(subject[i])
	}
	return b.String()
}

func Move(subject string, source, destination int) string {
	c := subject[source : source+1]
	without := subject[:source] + subject[source+1:]
	return without[:destination] + c + without[destination:]
}

func RotateByChar(subject string, char string) string {
	index := strings.Index(subject, char)
	subject = Rotate(subject, 1)
	subject = Rotate(subject, index)
	if index >= 4 {
		subject = Rotate(subject, 1)
	}
	return subject
}
