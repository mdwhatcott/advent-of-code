package starter

import (
	"container/list"
	"fmt"
	"sort"
	"strings"

	"github.com/mdwhatcott/funcy"
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
			_, _ = fmt.Fprintf(&result, "%c(%d)", unit.species, unit.HP())
			if u < len(units)-1 {
				result.WriteString(", ")
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}

func (this *World) readingOrder(points []grid.Point[int]) []grid.Point[int] {
	return funcy.SortAscending(this.pointID, points)
}
func (this *World) pointID(p grid.Point[int]) int {
	return p.Y()*this.width + p.X()
}
func (this *World) PlayRound() bool {
	// TODO: filter out dead units
	for _, key := range this.readingOrder(funcy.MapKeys(this.units)) {
		// TODO: check for no targets and return false
		unit := this.units[key]
		if this.weakestTargetInRangeOf(unit) == nil {
			step, ok := this.firstStepTowardsClosestTarget(unit)
			if ok {
				unit.location = step
				delete(this.units, key)
				this.units[step] = unit
			}
		}
		if target := this.weakestTargetInRangeOf(unit); target != nil {
			target.damage += 3
		}
	}
	return true
}

func (this *World) firstStepTowardsClosestTarget(mover *Unit) (result grid.Point[int], ok bool) {
	cave := set.Of[grid.Point[int]](this.cave.Slice()...)
	for location, target := range this.units {
		if target.HP() > 0 {
			cave.Add(location)
		}
	}
	var paths [][]grid.Point[int]
	for _, target := range mover.targets {
		paths = append(paths, findShortestPaths(cave, mover.location, target.location)...)
	}
	if len(paths) == 0 {
		return result, false
	}
	sort.SliceStable(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	candidates := set.Of[grid.Point[int]]()
	shortestLength := len(paths[0])
	for x := 0; x < len(paths); x++ {
		if len(paths[x]) > shortestLength {
			break
		}
		candidates.Add(paths[x][0])
	}
	return this.readingOrder(candidates.Slice())[0], true
}

func (this *World) weakestTargetInRangeOf(attacker *Unit) *Unit {
	minHP := 200
	byHP := make(map[int][]*Unit)
	for _, target := range attacker.targets {
		hp := target.HP()
		if hp > 0 && grid.CityBlockDistance(attacker.location, target.location) == 1 {
			byHP[hp] = append(byHP[hp], target)
			if hp < minHP {
				minHP = hp
			}
		}
	}
	if len(byHP) == 0 {
		return nil
	}
	weakestInRange := byHP[minHP]
	byLocations := make(map[grid.Point[int]]*Unit)
	for _, target := range weakestInRange {
		byLocations[target.location] = target
	}
	order := this.readingOrder(funcy.MapKeys(byLocations))
	return byLocations[order[0]]
}

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

type Unit struct {
	species  rune
	location grid.Point[int]
	targets  []*Unit
	damage   int
}

func NewUnit(species rune, x, y int) *Unit {
	return &Unit{species: species, location: grid.NewPoint(x, y)}
}
func (this *Unit) HP() int {
	return max(0, 200-this.damage)
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

func findShortestPaths(world set.Set[grid.Point[int]], start, end grid.Point[int]) (results [][]grid.Point[int]) {
	queue := list.New()
	queue.PushBack([]grid.Point[int]{start})
	visited := set.Of(start)
	for queue.Len() > 0 {
		path := queue.Remove(queue.Front()).([]grid.Point[int])
		current := path[len(path)-1]

		if grid.CityBlockDistance(current, end) == 1 && len(path) > 1 {
			results = append(results, path[1:])
			continue
		}
		for _, next := range current.Neighbors4() {
			if world.Contains(next) && !visited.Contains(next) {
				visited.Add(next)
				newPath := append([]grid.Point[int]{}, path...)
				newPath = append(newPath, next)
				queue.PushBack(newPath)
			}
		}
	}
	return results
}
