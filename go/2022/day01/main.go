package day01

import (
	"sort"
	"strings"

	"advent/lib/util"
)

func sums() (part1, part2 int) {
	var sums []int
	for _, chunk := range strings.Split(util.InputString(), "\n\n") {
		sums = append(sums, util.Sum(util.ParseInts(strings.Split(chunk, "\n"))...))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	return sums[0], util.Sum(sums[:3]...)
}
