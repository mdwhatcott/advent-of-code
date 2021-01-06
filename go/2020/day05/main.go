package advent

import (
	"sort"
	"strconv"
	"strings"

	"advent/lib/util"
)

func Part1() int {
	var max int
	for _, line := range util.InputLines() {
		id := parseSeatID(line)
		if id > max {
			max = id
		}
	}
	return max
}

func Part2() interface{} {
	var all []int
	for _, line := range util.InputLines() {
		all = append(all, parseSeatID(line))
	}

	sort.Ints(all)

	for x := 1; x < len(all)-1; x++ {
		last := all[x-1]
		this := all[x]
		next := all[x+1]
		if last+1 != this {
			return last + 1
		}
		if this+1 != next {
			return this + 1
		}
	}
	panic("NOT FOUND")
}

func parseSeatID(raw string) int {
	raw = strings.ReplaceAll(raw, "F", "0")
	raw = strings.ReplaceAll(raw, "B", "1")
	raw = strings.ReplaceAll(raw, "L", "0")
	raw = strings.ReplaceAll(raw, "R", "1")
	parsed, _ := strconv.ParseInt(raw, 2, 16)
	return int(parsed)
}
