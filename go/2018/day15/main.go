package starter

import (
	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/grid"
)

/*
for rounds := 0; ; rounds++
	sort characters in reading order
	for each character in reading order:
		if no targets:
			return rounds * sum(map(hit-points, living characters))

		if not adjacent to a target and there is any pathway to a target:
			move()

		if there is no pathway to any target:
			continue

		if adjacent to a target:
			attack()

*/

func ParseWorld(lines []string) (result set.Set[grid.Point[int]]) {
	result = set.Make[grid.Point[int]](0)
	for y, line := range lines {
		for x, char := range line {
			if char != '#' {
				result.Add(grid.NewPoint(x, y))
			}
		}
	}
	return result
}
func ParseCharacters(lines []string) (result []*Character) {
	for y, line := range lines {
		for x, char := range line {
			if char == 'G' || char == 'E' {
				result = append(result, NewCharacter(char, x, y))
			}
		}
	}
	return result
}

type Character struct {
	spec rune
	loc  grid.Point[int]
}

func NewCharacter(spec rune, x, y int) *Character {
	return &Character{spec: spec, loc: grid.NewPoint(x, y)}
}
