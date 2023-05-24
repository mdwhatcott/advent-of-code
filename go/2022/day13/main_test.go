package day13

import (
	"encoding/json"
	"sort"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code-go-lib/maths"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
	"github.com/mdwhatcott/advent-of-code/go/lib/must"
)

var (
	inputLines  = util.InputLines()
	sampleLines = []string{
		"[1,1,3,1,1]", "[1,1,5,1,1]", "",
		"[[1],[2,3,4]]", "[[1],4]", "",
		"[9]", "[[8,7,6]]", "",
		"[[4,4],4,4]", "[[4,4],4,4,4]", "",
		"[7,7,7,7]", "[7,7,7]", "",
		"[]", "[3]", "",
		"[[[]]]", "[[]]", "",
		"[1,[2,[3,[4,[5,6,7]]]],8,9]", "[1,[2,[3,[4,[5,6,0]]]],8,9]", "",
	}
)

func TestDay13(t *testing.T) {
	should.So(t, Part1(sampleLines), should.Equal, 13)
	should.So(t, Part1(inputLines), should.Equal, 5623)

	should.So(t, Part2(sampleLines), should.Equal, 140)
	should.So(t, Part2(inputLines), should.Equal, 20570)
}

func Part1(lines []string) (result int) {
	index := 1
	for x := 0; x < len(lines); x += 3 {
		a := lines[x+0]
		b := lines[x+1]
		if compare(parse(a), parse(b)) < 0 {
			result += index
		}
		index++
	}
	return result
}
func Part2(lines []string) (result int) {
	packets := []string{"[[2]]", "[[6]]"}
	for _, line := range lines {
		if len(line) > 0 {
			packets = append(packets, line)
		}
	}
	sort.Slice(packets, func(i, j int) bool {
		return compare(parse(packets[i]), parse(packets[j])) == -1
	})
	for p, packet := range packets {
		if packet == "[[2]]" {
			result = p + 1
		} else if packet == "[[6]]" {
			result *= p + 1
		}
	}
	return result
}
func compare(A, B any) int {
	// This is a Go translation of the python found here:
	// https://github.com/fogleman/AdventOfCode2022/blob/main/13.py#L8
	// Admittedly this problem, by nature, is somewhat clunky in a statically typed language.
	aa, aaOK := A.([]any)
	bb, bbOK := B.([]any)
	if aaOK && bbOK {
		for x := 0; x < maths.Min(len(aa), len(bb)); x++ {
			if comparison := compare(aa[x], bb[x]); comparison != 0 {
				return comparison
			}
		}
		return compare(float64(len(aa)), float64(len(bb)))
	} else if aaOK {
		return compare(A, []any{B})
	} else if bbOK {
		return compare([]any{A}, B)
	}
	an := A.(float64) // JSON parses all numbers to float64s
	bn := B.(float64)
	if an < bn {
		return -1
	} else if an > bn {
		return 1
	} else {
		return 0
	}
}
func parse(s string) (result any) {
	must.Nada(json.Unmarshal([]byte(s), &result))
	return result
}
