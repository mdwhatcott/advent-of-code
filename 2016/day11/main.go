package main

import "fmt"

func main() {
	fmt.Println("Part 1 - Solution calculated with temporary shortcut:", shortestNumberOfSteps(4, 5, 1, 0))
	fmt.Println("Part 2 - Solution calculated with temporary shortcut:", shortestNumberOfSteps(8, 5, 1, 0))
}

func shortestNumberOfSteps(itemsOnFloors ...int) (steps int) {
	for x := 1; x < 4; x++ {
		steps += 2*sum(itemsOnFloors[:x]) - 3
	}
	return steps

	// Python shortcut: (https://www.reddit.com/r/adventofcode/comments/5hoia9/2016_day_11_solutions/db27z3h/?context=3)
	// print sum(2 * sum([4, 5, 1, 0][:x]) - 3 for x in range(1,4))
}

func sum(all []int) (sum int) {
	for _, item := range all {
		sum += item
	}
	return sum
}
