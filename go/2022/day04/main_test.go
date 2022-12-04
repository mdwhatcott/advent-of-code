package day03

import (
	"regexp"
	"testing"

	"github.com/mdwhatcott/go-collections/set"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

var (
	inputLines  = util.InputLines()
	sampleLines = []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}
)

func TestDay04(t *testing.T) {
	part1, part2 := Analyze(sampleLines)
	should.So(t, part1, should.Equal, 2)
	should.So(t, part2, should.Equal, 4)

	part1, part2 = Analyze(inputLines)
	should.So(t, part1, should.Equal, 584)
	should.So(t, part2, should.Equal, 933)
}

var pattern = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

func Range(a, b int) (result []int) {
	for ; a < b; a++ {
		result = append(result, a)
	}
	return result
}
func Analyze(lines []string) (subsets, intersections int) {
	for _, line := range lines {
		matches := util.ParseInts(pattern.FindStringSubmatch(line)[1:])
		a, b, c, d := matches[0], matches[1], matches[2], matches[3]
		A := set.From(Range(a, b+1)...)
		B := set.From(Range(c, d+1)...)
		if A.IsSubset(B) || B.IsSubset(A) {
			subsets++
		}
		if A.Intersection(B).Len() > 0 {
			intersections++
		}
	}
	return subsets, intersections
}
