package day24

import "advent/lib/util"

func Part1() int {
	// NOT: 1350
	return MaxBridgeStrength(buildGraph(util.InputLines()).Traverse())
}

// TODO: Try this: https://github.com/resurtm/adventofcode-2017/blob/master/problem24/main.go#L72-L92

func Part2() int {
	return 0
}
