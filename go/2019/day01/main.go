package advent

import "github.com/mdwhatcott/advent-of-code/go/lib/util"

func Part1() (sum int) {
	for _, input := range util.InputInts("\n") {
		fuel := int(float64(input)/3.0) - 2
		sum = sum + fuel
	}
	return sum
}

func Part2() (sum int) {
	for _, input := range util.InputInts("\n") {
		module := 0
		fuel := int(float64(input)/3.0) - 2
		module = module + fuel

		for {
			fuel = int(float64(fuel)/3.0) - 2
			if fuel < 0 {
				break
			}
			module = module + fuel
		}
		sum = sum + module
	}
	return sum
}
