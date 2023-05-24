package advent

import (
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/parse"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func Part1() interface{} {
	lines := util.InputLines()
	timestamp := parse.Int(lines[0])
	busses := parse.Ints(strings.Split(lines[1], ","))
	min := 1_000_000
	minBus := -1
	for _, bus := range busses {
		if bus == 0 {
			continue
		}
		for x := 2; ; x++ {
			b := bus * x
			if b > timestamp {
				if b-timestamp < min {
					min = b - timestamp
					minBus = bus
				}
				break
			}
		}
	}
	return minBus * min
}

// See https://github.com/fogleman/AdventOfCode2020/blob/main/13.py
// Ugh.
func Part2() interface{} {
	rules := loadRules(util.InputLines()[1])
	timestamp := 0
	increment := 1

	for i := 1; i < len(rules)+1; i++ {
		marker := 0

		for {
			if checkRules(timestamp, rules[:i]) {

				if marker > 0 {
					increment = timestamp - marker
					timestamp = marker
					break
				}

				marker = timestamp
			}

			timestamp += increment
		}
	}

	return timestamp
}

func checkRules(timestamp int, rules []Rule) bool {
	for _, rule := range rules {
		if (timestamp+rule.Index)%rule.Cycle != 0 {
			return false
		}
	}
	return true
}

func loadRules(line string) (rules []Rule) {
	for i, c := range strings.Split(line, ",") {
		if c != "x" {
			rules = append(rules, Rule{Index: i, Cycle: parse.Int(c)})
		}
	}
	return rules
}

type Rule struct {
	Index int
	Cycle int
}
