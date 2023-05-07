package main

import (
	"fmt"
	"sort"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
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

func sum(values ...int) (sum int) {
	for _, value := range values {
		sum += value
	}
	return sum
}

func product(items ...int) (result int) {
	result = items[0]
	for _, item := range items[1:] {
		result *= item
	}
	return result
}

func combinations(iterable []int, count int) chan []int {
	stream := make(chan []int)
	go combinationGenerator(iterable, count, stream)
	return stream
}

// Credit: http://rosettacode.org/wiki/Combinations#Go
func combinationGenerator(source []int, count int, out chan []int) {
	working := make([]int, count)
	last := count - 1
	var generate func(int, int)
	generate = func(i, next int) {
		for j := next; j < len(source); j++ {
			working[i] = source[j]
			if i == last {
				combination := make([]int, len(working))
				copy(combination, working)
				out <- combination
			} else {
				generate(i+1, j+1)
			}
		}
	}
	generate(0, 0)
	close(out)
}
