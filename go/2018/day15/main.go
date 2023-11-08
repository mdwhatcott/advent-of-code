package starter

import (
	"sort"

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

func Battle(lines []string, debug bool) (roundCount int, hitPoints int, history []string) {
	world := ParseWorld(lines)
	if debug {
		history = append(history, RenderWorld(world))
	}
	for world.PlayRound() && len(history) < 100 {
		roundCount++
		if debug {
			history = append(history, RenderWorld(world))
		}
	}
	for _, unit := range world.units {
		hitPoints += unit.HP()
	}
	return roundCount, hitPoints, history
}

type World struct {
	height int
	width  int
	cave   set.Set[grid.Point[int]]
	units  map[grid.Point[int]]*Unit
}

func (this *World) PlayRound() bool {
	for _, key := range this.readingOrder(funcy.MapKeys(this.units)) {
		unit := this.units[key]
		if unit == nil {
			continue
		}
		targetCount := 0
		for _, target := range this.units {
			if target != nil && target.species != unit.species {
				if target.HP() > 0 {
					targetCount++
				}
			}
		}
		if targetCount == 0 {
			return false
		}
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
			if target.HP() == 0 {
				delete(this.units, target.location)
			}
		}
	}
	return true
}
func (this *World) readingOrder(points []grid.Point[int]) []grid.Point[int] {
	return funcy.SortAscending(this.pointID, points)
}
func (this *World) pointID(p grid.Point[int]) int {
	return p.Y()*this.width + p.X()
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
