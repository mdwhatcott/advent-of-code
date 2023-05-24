package advent

import "github.com/mdwhatcott/advent-of-code-go-lib/util"

func Part1() interface{} {
	for _, a := range util.InputInts("\n") {
		for _, b := range util.InputInts("\n") {
			if a+b == 2020 {
				return a * b
			}
		}
	}
	return nil
}

func Part2() interface{} {
	for _, a := range util.InputInts("\n") {
		for _, b := range util.InputInts("\n") {
			for _, c := range util.InputInts("\n") {
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}
	return nil
}
