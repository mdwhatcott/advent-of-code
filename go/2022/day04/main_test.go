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
	should.So(t, Part1(sampleLines), should.Equal, 2)
	should.So(t, Part1(inputLines), should.Equal, 584)

	should.So(t, Part2(sampleLines), should.Equal, 4)
	should.So(t, Part2(inputLines), should.Equal, 933)
}

var pattern = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

func Part1(lines []string) (result int) {
	for _, line := range lines {
		matches := util.ParseInts(pattern.FindAllStringSubmatch(line, 4)[0][1:])
		a, b, c, d := matches[0], matches[1], matches[2], matches[3]
		A := set.New[int](0)
		B := set.New[int](0)
		for x := a; x <= b; x++ {
			A.Add(x)
		}
		for x := c; x <= d; x++ {
			B.Add(x)
		}
		if A.IsSubset(B) || B.IsSubset(A) {
			result++
		}
	}
	return result
}
func Part2(lines []string) (result int) {
	for _, line := range lines {
		matches := util.ParseInts(pattern.FindAllStringSubmatch(line, 4)[0][1:])
		a, b, c, d := matches[0], matches[1], matches[2], matches[3]
		A := set.New[int](0)
		B := set.New[int](0)
		for x := a; x <= b; x++ {
			A.Add(x)
		}
		for x := c; x <= d; x++ {
			B.Add(x)
		}
		if A.Intersection(B).Len() > 0 {
			result++
		}
	}
	return result
}
