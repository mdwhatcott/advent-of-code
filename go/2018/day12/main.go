package day12

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func Part1() interface{} {
	lines := util.InputLines()
	initial := strings.Fields(lines[0])[2]
	var rules []Rule
	for _, line := range lines[2:] {
		rules = append(rules, ParseRule(line))
	}
	return NewRowOfPots(initial, rules...).Grow(20)
}

func Part2() interface{} {
	lines := util.InputLines()
	initial := strings.Fields(lines[0])[2]
	var rules []Rule
	for _, line := range lines[2:] {
		rules = append(rules, ParseRule(line))
	}

	oldDiff := 0
	diffRepeats := 0
	x := 0
	old := 0

	row := NewRowOfPots(initial, rules...)

	for {
		x++
		row.Update(row.Scan())
		sum := row.Sum()

		// Wait for the diff between sums of subsequent generations to stabilize, then calculate the end result.
		diff := sum - old
		if diff == oldDiff {
			diffRepeats++
		} else {
			oldDiff = diff
			diffRepeats = 0
		}
		old = sum

		if diffRepeats > 10 {
			fmt.Println(x, sum)
			return sum + (diff * (50000000000 - x))
		}
	}
}
