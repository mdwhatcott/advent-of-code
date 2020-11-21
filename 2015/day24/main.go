package main

import (
	"fmt"
	"sort"

	"advent/lib/util"
)

func main() {
	arrangePresentsOnSleigh(3) // 10439961859
	arrangePresentsOnSleigh(4) // 72050269
}

func arrangePresentsOnSleigh(compartments int) {
	quantumEntanglement := 0
	inputs := util.InputInts("\n")
	target := sum(inputs...) / compartments

	// Use heaviest packages first to achieve target weight with fewer packages.
	sort.Sort(sort.Reverse(sort.IntSlice(inputs)))

	// Range of 5..7 packages to get best quantum entanglement for
	// 3 or 4 compartments was arrived at by trial and error.
	for x := 5; x < 7; x++ {
		for combo := range combinations(inputs, x) {
			if sum(combo...) == target {
				if quantumEntanglement == 0 {
					quantumEntanglement = product(combo...)
				} else {
					proposed := product(combo...)
					if proposed < quantumEntanglement {
						quantumEntanglement = proposed
					}
				}
			}
		}
	}

	fmt.Println(quantumEntanglement)
}
