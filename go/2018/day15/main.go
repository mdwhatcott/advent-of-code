package starter

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/grid"
)

/*
for rounds := 0; ; rounds++
	sort units in reading order
	for each unit in reading order:
		if no targets:
			return rounds * sum(map(hit-points, living units))

		if not adjacent to a target and there is any pathway to a target:
			move()

		if there is no pathway to any target:
			continue

		if adjacent to a target:
			attack()
*/

type World struct {
	height int
	width  int
	cave   set.Set[grid.Point[int]]
	units  map[grid.Point[int]]*Unit
}

func (this *World) String() string {
	var result strings.Builder
	result.WriteString("\n")
	for y := 0; y < this.height; y++ {
		var units []*Unit
		for x := 0; x < this.width; x++ {
			at := grid.NewPoint(x, y)
			if unit, ok := this.units[at]; ok {
				units = append(units, unit)
				result.WriteRune(unit.species)
			} else if this.cave.Contains(at) {
				result.WriteRune('.')
			} else {
				result.WriteRune('#')
			}
		}
		if len(units) == 0 {
			result.WriteString("\n")
			continue
		}
		result.WriteString("   ")
		for u, unit := range units {
			_, _ = fmt.Fprintf(&result, "%c(%d)", unit.species, 200) // TODO: HP
			if u < len(units)-1 {
				result.WriteString(", ")
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}

func ParseWorld(lines []string) *World {
	units := ParseUnits(lines)
	index := make(map[grid.Point[int]]*Unit)
	for _, unit := range units {
		index[unit.location] = unit
	}
	world := &World{
		height: len(lines),
		width:  len(lines[0]),
		cave:   ParseCaveMap(lines),
		units:  index, // TODO: associate
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
			if char == 'G' || char == 'E' {
				result = append(result, NewUnit(char, x, y))
			}
		}
	}
	return result
}

type Unit struct {
	species  rune
	location grid.Point[int]
	targets  []*Unit
}

func NewUnit(species rune, x, y int) *Unit {
	return &Unit{species: species, location: grid.NewPoint(x, y)}
}

func AssociateEnemyUnits(all ...*Unit) {
	for c, c1 := range all {
		for _, c2 := range all[c+1:] {
			if c1.species == c2.species {
				continue
			}
			c1.targets = append(c1.targets, c2)
			c2.targets = append(c2.targets, c1)
		}
	}
}
