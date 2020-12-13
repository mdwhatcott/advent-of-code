package advent

import (
	"strings"

	"advent/lib/util"
)

func Part1() interface{} {
	lines := util.InputLines()
	timestamp := util.ParseInt(lines[0])
	busses := util.ParseInts(strings.Split(lines[1], ","))
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

func Part2() interface{} {
	return nil
}
