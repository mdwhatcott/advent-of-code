package main

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func main() {
	lines := util.InputLines()
	var discs []Disc
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		discs = append(discs, ParseDisc(line))
	}

	for x := 1; ; x++ {
		if allReady(discs, x) {
			fmt.Println("Part 1: All discs ready at time ==", x)
			break
		}
	}

	discs = append(discs, Disc{
		Positions: 11,
		Start:     0,
		Delay:     len(lines) + 1,
	})

	for x := 1; ; x++ {
		if allReady(discs, x) {
			fmt.Println("Part 2: All discs ready at time ==", x)
			break
		}
	}
}

func allReady(discs []Disc, t int) bool {
	for _, disc := range discs {
		if disc.PositionAtTime(t) != 0 {
			return false
		}
	}
	return true
}
