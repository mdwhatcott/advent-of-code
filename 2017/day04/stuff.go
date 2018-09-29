package day04

import (
	"strings"

	"advent/lib/util"
)

func Valid(s string) bool {
	return Valid1(s) && Valid2(s)
}

func Valid1(s string) bool {
	fields := strings.Fields(s)
	length := len(fields)
	words := make(map[string]struct{})
	for _, word := range fields {
		words[word] = struct{}{}
	}
	return length == len(words)
}

func Valid2(s string) bool {
	fields := strings.Fields(s)
	for a, x := range fields {
		for b, y := range fields {
			if a == b {
				continue
			}
			if util.Anagram(x, y) {
				return false
			}
		}
	}
	return true
}
