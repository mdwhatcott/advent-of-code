package day05

import (
	"bytes"
	"strings"
	"unicode"
)

func react(s string) string {
	b := bytes.NewBufferString(s[:1])
	for x := 1; x < len(s); x++ {
		if b.Len() == 0 {
			b.WriteByte(s[x])
			x++
		}
		this, last := s[x], b.Bytes()[b.Len()-1]
		if this != last && strings.ToUpper(string(this)) == strings.ToUpper(string(last)) {
			b = bytes.NewBuffer(b.Bytes()[:b.Len()-1])
		} else {
			b.WriteByte(this)
		}
	}
	return b.String()
}
func reactAggressive(s string) string {
	minLength := len(s)
	minPolymer := s
	for _, c := range "abcdefghijklmnopqrstuvwxyz" {
		ss := s
		ss = strings.Replace(ss, string(c), "", -1)
		ss = strings.Replace(ss, string(unicode.ToUpper(c)), "", -1)
		polymer := react(ss)
		length := len(polymer)
		if length < minLength {
			minLength = length
			minPolymer = polymer
		}
	}
	return minPolymer
}
