package day15

import (
	"fmt"
	"log"
	"slices"
	"sort"
	"strings"

	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/go-set/v2/set"
)

func ParseCave(lines []string) (units []*Unit, walls set.Set[Point]) {
	walls = set.Of[Point]()
	maxLineLength := len(lines[0])
	for y, line := range lines {
		for x, char := range line {
			if x >= maxLineLength {
				break
			}
			switch char {
			case '#':
				walls.Add(XY(x, y))
			case 'G':
				units = append(units, NewUnit(x, y, string(char)))
			case 'E':
				units = append(units, NewUnit(x, y, string(char)))
			}
		}
	}
	return units, walls
}
func RenderCave(units []*Unit, walls set.Set[Point]) string {
	var all []point
	for _, g := range units {
		all = append(all, g)
	}
	for p := range walls {
		all = append(all, p)
	}
	all = Sort(all)
	minWall, maxWall := funcy.First(all), funcy.Last(all)

	index := make(map[Point]point)
	for _, pin := range all {
		index[XY(pin.x(), pin.y())] = pin
	}

	var builder strings.Builder
	for y := minWall.y(); y <= maxWall.y(); y++ {
		for x := minWall.x(); x <= maxWall.x(); x++ {
			item, ok := index[XY(x, y)]
			if ok {
				if _, ok := item.(Point); ok {
					builder.WriteString("#")
				} else if unit, ok := item.(*Unit); ok {
					builder.WriteString(unit.Team)
				}
			} else {
				builder.WriteString(".")
			}
		}
		builder.WriteString("\n")
	}
	return strings.TrimSpace(builder.String())
}
func AnnotateUnits(rendering string, units []*Unit) string {
	var builder strings.Builder
	index := make(map[Point]*Unit)
	for _, unit := range units {
		index[unit.Point] = unit
	}
	lines := strings.Split(rendering, "\n")
	for y, line := range lines {
		var lineUnits []string
		for x := range line {
			unit, ok := index[XY(x, y)]
			if ok {
				lineUnits = append(lineUnits, unit.Status())
			}
		}
		builder.WriteString(line)
		if len(lineUnits) > 0 {
			builder.WriteString("   ")
		}
		builder.WriteString(strings.Join(lineUnits, ", "))
		builder.WriteString("\n")
	}
	return strings.TrimSpace(builder.String())
}

func Points[T point](things []T) (result []Point) {
	for _, thing := range things {
		result = append(result, XY(thing.x(), thing.y()))
	}
	return result
}

type Point struct{ X, Y int }

func XY(x, y int) Point { return Point{X: x, Y: y} }

func (this Point) x() int { return this.X }
func (this Point) y() int { return this.Y }

func (this Point) String() string { return fmt.Sprintf("(%d,%d)", this.X, this.Y) }

type point interface {
	fmt.Stringer
	x() int
	y() int
}

func Neighbors(p point) []Point {
	x, y := p.x(), p.y()
	return []Point{
		XY(x+0, y-1),
		XY(x-1, y+0),
		XY(x+1, y+0),
		XY(x+0, y+1),
	}
}

type Unit struct {
	Point
	Team   string
	Health int
	Attack int
}

func NewUnit(x, y int, team string) *Unit {
	return &Unit{
		Point:  XY(x, y),
		Team:   team,
		Health: 200,
		Attack: 3,
	}
}

func (this *Unit) Status() string {
	return fmt.Sprintf("%s(%d)", this.Team, this.Health)
}

func FilterTeam(units []*Unit, team string) (results []*Unit) {
	for _, unit := range units {
		if unit.Team == team {
			results = append(results, unit)
		}
	}
	return results
}

func Sort[T point](points []T) []T {
	sort.SliceStable(points, func(i, j int) bool {
		if points[i].y() == points[j].y() {
			return points[i].x() < points[j].x()
		}
		return points[i].y() < points[j].y()
	})
	return points
}

func BFS(origin, target Point, obstacles set.Set[Point]) (result []Point) {
	frontier := []Point{origin}
	crumbs := make(map[Point]Point)
	crumbs[origin] = Point{}
	for len(frontier) > 0 {
		at := frontier[0]
		frontier = frontier[1:]
		if at == target {
			break
		}
		for _, neighbor := range Neighbors(at) {
			if obstacles.Contains(neighbor) {
				continue
			}
			_, ok := crumbs[neighbor]
			if ok {
				continue
			}
			frontier = append(frontier, neighbor)
			crumbs[neighbor] = at
		}
	}
	var ok bool
	current := target
	for current != origin {
		result = append(result, current)
		current, ok = crumbs[current]
		if !ok {
			return nil
		}
	}
	slices.Reverse(result)
	return result
}

func MoveUnit(unit *Unit, units []*Unit, obstacles set.Set[Point]) {
	enemies := FilterTeam(units, EnemyOf[unit.Team])
	occupied := set.Of(Points(units)...).Difference(set.Of(unit.Point))
	obstacles = obstacles.Union(occupied)

	minDistance := 0xFFFF
	targets := make(map[Point][]Point)
	for _, enemy := range enemies {
		for _, target := range Neighbors(enemy.Point) {
			if unit.Point == target { // already in range of a target
				return
			}
			path := BFS(unit.Point, target, obstacles)
			if len(path) == 0 { // invalid target
				continue
			}
			if len(path) < minDistance {
				minDistance = len(path)
			}
			targets[target] = path
		}
	}
	if len(targets) == 0 {
		return
	}
	for target, path := range targets {
		if len(path) > minDistance {
			delete(targets, target)
		}
	}
	keys := Sort(funcy.MapKeys(targets))
	path := targets[keys[0]]
	step := path[0]
	unit.Point = step
}

func MoveAll(units []*Unit, walls set.Set[Point]) []*Unit {
	for _, unit := range Sort(units) {
		MoveUnit(unit, units, walls)
	}
	return units
}

var EnemyOf = map[string]string{
	"E": "G",
	"G": "E",
}

func SimulateRound(units []*Unit, walls set.Set[Point]) (gameOver bool, results []*Unit) {
	for _, unit := range Sort(units) {
		if unit.Health <= 0 {
			continue
		}
		units = Living(units)
		enemies := FilterTeam(units, EnemyOf[unit.Team])
		if len(enemies) == 0 {
			gameOver = true
			break
		}
		MoveUnit(unit, units, walls)
		AttackWith(unit, units)
	}
	return gameOver, Living(units)
}

func Living(units []*Unit) (results []*Unit) {
	for _, unit := range units {
		if unit.Health > 0 {
			results = append(results, unit)
		}
	}
	return results
}

func AttackWith(attacker *Unit, units []*Unit) {
	enemies := FilterTeam(units, EnemyOf[attacker.Team])
	targets := make(map[Point][]*Unit)
	minHealth := 200
	for _, enemy := range enemies {
		for _, target := range Neighbors(enemy.Point) {
			if attacker.Point == target {
				targets[target] = append(targets[target], enemy)
				if enemy.Health < minHealth {
					minHealth = enemy.Health
				}
			}
		}
	}
	if len(targets) == 0 {
		return
	}
	minEnemies := make(map[Point][]*Unit)
	for target, enemies := range targets {
		for _, enemy := range enemies {
			if enemy.Health == minHealth {
				minEnemies[target] = append(minEnemies[target], enemy)
			}
		}
	}
	keys := Sort(funcy.MapKeys(minEnemies))
	target := Sort(minEnemies[keys[0]])[0]
	target.Health -= attacker.Attack
}

func TotalHealth(units []*Unit) (result int) {
	for _, unit := range Living(units) {
		result += unit.Health
	}
	return result
}

func BeverageBanditsBattle(units []*Unit, walls set.Set[Point]) (rounds, health int, steps []string) {
	steps = append(steps, FullRendering(units, walls))
	for gameOver := false; ; {
		gameOver, units = SimulateRound(units, walls)
		if gameOver {
			steps = append(steps, FullRendering(units, walls))
			break
		}
		steps = append(steps, FullRendering(units, walls))
		rounds++
	}
	return rounds, TotalHealth(units), steps
}

func FullRendering(units []*Unit, walls set.Set[Point]) string {
	return AnnotateUnits(RenderCave(units, walls), units)
}

func BoostedBeverageBanditsBattle(input []string) (rounds, health int, rendering string) {
	for attackBonus := 1; ; attackBonus++ {
		log.Println("Elf attack strength:", 3+attackBonus)
		units, walls := ParseCave(input)
		starters := FilterTeam(units, "E")
		BoostAll(attackBonus, starters)
		rounds, health, steps := BeverageBanditsBattle(units, walls)
		survivors := FilterTeam(Living(units), "E")
		if len(survivors) == len(starters) {
			return rounds, health, funcy.Last(steps)
		}
	}
}

func BoostAll(bonus int, units []*Unit) {
	for _, unit := range units {
		unit.Attack += bonus
	}
}
