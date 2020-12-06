package advent

import (
	"strings"

	"advent/lib/util"
)

func Part1() interface{} {
	total := 0
	for _, block := range strings.Split(util.InputString(), "\n\n") {
		unique := make(map[rune]struct{})
		for _, c := range block {
			if c == ' ' {
				continue
			}
			if c == '\n' {
				continue
			}
			unique[c] = struct{}{}
		}
		total += len(unique)
	}
	return total
}

func Part2() interface{} {
	total := 0
	for _, block := range strings.Split(util.InputString(), "\n\n") {
		unique := make(map[rune]int)
		people := strings.Split(block, "\n")
		for _, person := range people {
			for _, c := range person {
				if c == ' ' {
					continue
				}
				if c == '\n' {
					continue
				}
				unique[c]++
			}
		}
		for _, n := range unique {
			if n == len(people) {
				total++
			}
		}
	}
	return total
}
