package day18

import (
	"github.com/mdwhatcott/advent-of-code/go/lib/intgrid"
	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func Part1() interface{} {
	field := parseField()
	for x := 0; x < 10; x++ {
		tick(field)
	}
	return field.resourceValue()
}

func Part2() interface{} {
	// Let's detect a cycle and its period to project the 1,000,000,000th value...
	field := parseField()
	index := make(map[int][]int)
	for x := 1; x < 1000; x++ { // hopefully we arrive at cyclical values before 1000...
		tick(field)
		value := field.resourceValue()
		hits := index[value]
		index[value] = append(hits, x)
	}

	period := -1
	for _, hits := range index {
		// hopefully 5 or more hits on a value indicates a cycle...
		if len(hits) > 5 {
			// the difference between the last two hits is the period of the cycle...
			period = hits[1] - hits[0]
		}
	}

	// walk back from the target (in period-length steps) to the range of the index
	target := 1_000_000_000
	for target > 500 {
		target -= period
	}

	// search for the value associated with the target
	// (which will correspond with the 1_000_000_000th minute because of the cycle)
	for value, hits := range index {
		for _, hit := range hits {
			if hit == target {
				return value
			}
		}
	}

	panic("FAIL")
}

type Field map[intgrid.Point]string

func parseField() Field {
	field := make(Field)
	for y, line := range util.InputLines() {
		for x, char := range line {
			field[intgrid.NewPoint(x, y)] = string(char)
		}
	}
	return field
}

func (this Field) resourceValue() int {
	trees := 0
	lumbers := 0
	for _, acre := range this {
		switch acre {
		case "|":
			trees++
		case "#":
			lumbers++
		}
	}
	return trees * lumbers
}

func (this Field) String() (result string) {
	for _, acre := range this {
		if len(result)%50 == 0 {
			result += "\n"
		}
		result += acre
	}
	return result
}

func tick(field map[intgrid.Point]string) {
	updates := make(map[intgrid.Point]string, len(field))
	for key := range updates {
		delete(updates, key)
	}
	for point, acre := range field {
		treeNeighbors := countNeighbors(field, point, "|")
		lumberNeighbors := countNeighbors(field, point, "#")

		switch acre {
		case ".":
			if treeNeighbors >= 3 {
				updates[point] = "|"
			}
		case "|":
			if lumberNeighbors >= 3 {
				updates[point] = "#"
			}
		case "#":
			if lumberNeighbors == 0 || treeNeighbors == 0 {
				updates[point] = "."
			}
		}
	}
	for point, update := range updates {
		field[point] = update
	}
}

func countNeighbors(field map[intgrid.Point]string, point intgrid.Point, char string) (result int) {
	for _, neighbor := range point.Neighbors8() {
		if field[neighbor] == char {
			result++
		}
	}
	return result
}
