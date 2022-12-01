package main

import (
	"fmt"
	"sort"
	"strings"

	"advent/lib/util"
)

func main() {
	topElf, top3Elves := sums()
	fmt.Println("Part 1:", topElf)
	fmt.Println("Part 2:", top3Elves)
}
func sums() (int, int) {
	var sums []int
	for _, chunk := range strings.Split(util.InputString(), "\n\n") {
		sums = append(sums, util.Sum(util.ParseInts(strings.Split(chunk, "\n"))...))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	return sums[0], util.Sum(sums[:3]...)
}
