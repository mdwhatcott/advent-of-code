package advent

import "github.com/mdwhatcott/advent-of-code/go/lib/util"

func Part1() int {
	return Forest(util.InputLines()).CountTreesHitOnDescent(1, 3)
}

func Part2() int {
	forest := Forest(util.InputLines())
	return forest.CountTreesHitOnDescent(1, 1) *
		forest.CountTreesHitOnDescent(1, 3) *
		forest.CountTreesHitOnDescent(1, 5) *
		forest.CountTreesHitOnDescent(1, 7) *
		forest.CountTreesHitOnDescent(2, 1)
}

type Forest []string

func (this Forest) CountTreesHitOnDescent(rise, run int) (hits int) {
	row, col := 0, 0
	for row < len(this) {
		if this[row][col] == tree {
			hits++
		}
		row += rise
		col = (col + run) % len(this[0])
	}
	return hits
}

const tree = '#'
