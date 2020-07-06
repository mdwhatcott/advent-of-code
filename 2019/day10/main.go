package advent

import "advent/lib/util"

func Part1() interface{} {
	field := scanField(util.InputLines())
	best := BestPlace(field)
	return CountVisible(field, best)
}

func Part2() interface{} {
	return nil
}
