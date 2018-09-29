package day04

import (
	"strings"

	"advent/lib/util"
)

func Part1() int {
	v, _ := parts()
	return v
}

func Part2() int {
	_, v := parts()
	return v
}

func parts() (valid1, valid2 int) {
	input := util.InputScanner()

	for input.Scan() {
		line := input.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if Valid1(line) {
			valid1++
		}
		if Valid(line) {
			valid2++
		}
	}

	return valid1, valid2
}
