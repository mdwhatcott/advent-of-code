package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func main() {
	blocked := make([]bool, 4294967295+1)
	for _, line := range util.InputLines() {
		if len(line) == 0 {
			continue
		}
		values := strings.Split(line, "-")
		min, _ := strconv.Atoi(values[0])
		max, _ := strconv.Atoi(values[1])
		for x := min; x <= max; x++ {
			blocked[x] = true
		}
	}

	min := -1
	allowed := 0
	for x := 0; x < len(blocked); x++ {
		if !blocked[x] {
			if min < 0 {
				fmt.Println("Part 1 - Lowest unblocked IP:", x)
				min = x
			}
			allowed++
		}
	}
	fmt.Println("Part 2 - How many IPs allowed?", allowed)
}
