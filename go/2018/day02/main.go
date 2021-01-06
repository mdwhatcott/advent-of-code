package day02

import "advent/lib/util"

func Part1() int {
	twos := 0
	threes := 0
	for _, line := range util.InputLines() {
		all := make(map[rune]int)
		for _, c := range line {
			all[c]++
		}
		two, three := false, false
		for _, r := range all {
			if r == 2 && !two {
				twos++
				two = true
			}
			if r == 3 && !three {
				threes++
				three = true
			}
		}
	}
	return twos * threes
}

func Part2() string {
	all := util.InputLines()
	for x := 0; x < len(all); x++ {
		for y := 0; y < len(all); y++ {
			compare := same(all[x], all[y])
			if len(compare) == len(all[x])-1 {
				return compare
			}
		}
	}
	panic("Nada")
}

func same(a, b string) (same string) {
	for x := 0; x < len(a); x++ {
		if a[x] == b[x] {
			same += string(a[x])
		}
	}
	return same
}
