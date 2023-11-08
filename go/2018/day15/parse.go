package starter

import (
	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/grid"
)

func ParseWorld(lines []string) *World {
	units := ParseUnits(lines)
	AssociateEnemyUnits(units...)
	index := make(map[grid.Point[int]]*Unit)
	for _, unit := range units {
		index[unit.location] = unit
	}
	world := &World{
		height: len(lines),
		width:  len(lines[0]),
		cave:   ParseCaveMap(lines),
		units:  index,
	}
	return world
}
func ParseCaveMap(lines []string) (result set.Set[grid.Point[int]]) {
	result = set.Make[grid.Point[int]](0)
	for y, line := range lines {
		for x, char := range line {
			if x >= len(lines[0]) {
				break
			}
			if char != '#' {
				result.Add(grid.NewPoint(x, y))
			}
		}
	}
	return result
}
func ParseUnits(lines []string) (result []*Unit) {
	for y, line := range lines {
		for x, char := range line {
			if x >= len(lines[0]) {
				break
			}
			if char == 'G' || char == 'E' {
				result = append(result, NewUnit(char, x, y))
			}
		}
	}
	return result
}
