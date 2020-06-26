package advent

import "advent/lib/util"

func Part1() interface{} {
	program := util.InputInts(",")
	max := 0
	for _, combo := range phaseCombinations {
		answer := NewIO(combo...).Run(program...)
		if answer > max {
			max = answer
		}
	}
	return max
}

func Part2() interface{} {
	return nil
}
