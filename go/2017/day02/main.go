package day02

import (
	"strings"

	"advent/lib/maths"
	"advent/lib/parse"
	"advent/lib/util"
)

func Part1() int {
	return part1Checksum(util.InputLines())
}

func Part2() int {
	return part2Checksum(util.InputLines())
}

func part1Checksum(lines []string) int {
	checksum := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		ints := parse.Ints(fields)
		min, max := maths.MinMax(ints...)
		diff := max - min
		checksum += diff
	}
	return checksum
}

func part2Checksum(lines []string) int {
	checksum := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		ints := parse.Ints(fields)
		checksum += divider(ints)
	}
	return checksum
}

func divider(nums []int) int {
	for _, a := range nums {
		for _, b := range nums {
			if a == b {
				continue
			}
			if a%b == 0 && a/b > 0 {
				return a / b
			}
		}
	}
	panic("BAD ROW")
}
