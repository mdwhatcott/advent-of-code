package advent

import "advent/lib/util"

func Part1() interface{} {
	field := scanField(util.InputLines())
	best := BestPlaceWithCount(field)
	return best.Count
}

func Part2() interface{} {
	return nil
}
