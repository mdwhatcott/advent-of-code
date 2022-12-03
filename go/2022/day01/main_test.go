package day01

import (
	"sort"
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

var sampleInput = strings.TrimSpace(`
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)

func TestDay01(t *testing.T) {
	sample1, sample2 := sums(sampleInput)
	should.So(t, sample1, should.Equal, 24000)
	should.So(t, sample2, should.Equal, 45000)

	part1, part2 := sums(util.InputString())
	should.So(t, part1, should.Equal, 71924)
	should.So(t, part2, should.Equal, 210406)
}

func sums(input string) (topElf, top3Elves int) {
	var sums []int
	for _, chunk := range strings.Split(input, "\n\n") {
		sums = append(sums, util.Sum(util.ParseInts(strings.Split(chunk, "\n"))...))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	return sums[0], util.Sum(sums[:3]...)
}
