package starter

import (
	"container/list"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/grid"
)

type Point grid.Point[int]

func NewPoint(x, y int) Point             { return Point(grid.NewPoint(x, y)) }
func (this Point) Point() Point           { return this }
func (this Point) X() int                 { return this.point().X() }
func (this Point) Y() int                 { return this.point().Y() }
func (this Point) point() grid.Point[int] { return grid.Point[int](this) }
func (this Point) String() string         { return this.point().String() }

type Unit struct {
	point Point
	class rune
}

func NewUnit(x, y int, class rune) *Unit {
	return &Unit{
		point: NewPoint(x, y),
		class: class,
	}
}
func (this *Unit) String() string { return string(this.class) }
func (this *Unit) Point() Point   { return this.point }
func (this *Unit) GoTo(p Point)   { this.point = p }

type (
	Path []Point

	Pin interface {
		fmt.Stringer
		Point() Point
	}
	Mobile interface {
		GoTo(Point)
	}
)

func ParseCave(lines []string) (elves, goblins []*Unit, walls set.Set[Point]) {
	walls = set.Of[Point]()
	for y, line := range lines {
		for x, char := range line {
			switch char {
			case '#':
				walls.Add(NewPoint(x, y))
			case 'G':
				goblins = append(goblins, NewUnit(x, y, char))
			case 'E':
				elves = append(elves, NewUnit(x, y, char))
			}
		}
	}
	return elves, goblins, walls
}
func RenderCave(elves, goblins []*Unit, walls set.Set[Point]) string {
	var all []Pin
	for _, e := range elves {
		all = append(all, e)
	}
	for _, g := range goblins {
		all = append(all, g)
	}
	for p := range walls {
		all = append(all, p)
	}
	SortReadingOrder(all)
	minWall, maxWall := funcy.First(all).Point(), funcy.Last(all).Point()

	index := make(map[Point]Pin)
	for _, pin := range all {
		index[pin.Point()] = pin
	}

	var builder strings.Builder
	for y := minWall.Y(); y <= maxWall.Y(); y++ {
		for x := minWall.X(); x <= maxWall.X(); x++ {
			item, ok := index[NewPoint(x, y)]
			if ok {
				if _, ok := item.(Point); ok {
					builder.WriteString("#")
				} else {
					builder.WriteString(item.String())
				}
			} else {
				builder.WriteString(".")
			}
		}
		builder.WriteString("\n")
	}
	return strings.TrimSpace(builder.String())
}
func SortReadingOrder[T Pin](points []T) {
	sort.SliceStable(points, func(i, j int) bool {
		I := points[i].Point()
		J := points[j].Point()
		if I.Y() == J.Y() {
			return I.X() < J.X()
		}
		return I.Y() < J.Y()
	})
}
func Points[T Pin](units []T) (result []Point) {
	return funcy.Map(T.Point, units)
}
func MoveUnit[T Pin](actor T, targets []T, obstacles set.Set[Point]) {
	log.Println("actor", actor, "at", actor.Point(), "with targets", len(targets))
	// Calculate all paths to all targets:
	var all []Path
	for _, target := range targets {
		log.Println("aiming for", target.Point())
		newPaths := findShortestPaths(actor.Point(), target.Point(), obstacles)
		for _, p := range newPaths {
			log.Println("path found:", p)
		}
		all = append(all, newPaths...)
	}
	if len(all) == 0 {
		return
	}

	// Determine length of shortest path(s):
	sort.SliceStable(all, func(i, j int) bool { return len(all[i]) < len(all[j]) })
	shortestPathLength := len(all[0])
	if shortestPathLength == 0 {
		return
	}

	// Filter the shortest paths and index them by their final step:
	shortestByLast := make(map[Point][]Path)
	for _, path := range all {
		if len(path) > shortestPathLength {
			break
		}
		shortestByLast[path[len(path)-1]] = append(shortestByLast[path[len(path)-1]], path)
	}

	// Sort (in reading order) the shortest paths by final step:
	keys := funcy.MapKeys(shortestByLast)
	SortReadingOrder(keys)

	// Filter the shortest paths and index them by their first step:
	candidates := shortestByLast[keys[0]]
	for _, candidate := range candidates {
		log.Println("actor", actor, "at", actor.Point(), "candidate path", candidate)
	}
	byFirst := make(map[Point]Path)
	for _, path := range candidates {
		byFirst[path[0]] = path
	}

	// Sort (in reading order) the shortest paths to the previously selected last step by their first step:
	keys = funcy.MapKeys(byFirst)
	SortReadingOrder(keys)

	// Actor takes first step along best path:
	best := byFirst[keys[0]]
	log.Println("actor", actor, "at", actor.Point(), "will move on path", best)
	any(actor).(Mobile).GoTo(best[0])
}
func findShortestPaths(start, end Point, obstacles set.Set[Point]) (results []Path) {
	queue := list.New()
	queue.PushBack([]Point{start})
	visited := set.Of(start)
	for queue.Len() > 0 {
		path := queue.Remove(queue.Front()).([]Point)
		current := funcy.Last(path).point()

		if grid.CityBlockDistance(current, end.point()) == 1 && len(path) > 1 {
			results = append(results, path[1:])
			continue
		}
		for _, next := range current.Neighbors4() {
			point := Point(next)
			if !obstacles.Contains(point) && !visited.Contains(point) {
				visited.Add(point)
				newPath := append([]Point{}, path...)
				newPath = append(newPath, point)
				queue.PushBack(newPath)
			}
		}
	}
	return results
}
func MoveAll(elves, goblins []*Unit, walls set.Set[Point]) (outElves, outGoblins []*Unit) {
	all := append(elves, goblins...)
	SortReadingOrder(all)
	for _, unit := range all {
		var others = goblins
		if unit.class == 'G' {
			others = elves
		}
		log.Println("-----")
		log.Println("actor", unit, "starts at", unit.Point())
		MoveUnit(unit, others, walls.Union(set.Of(Points(all)...)))
		log.Println("actor", unit, "moved to ", unit.Point())
	}
	return elves, goblins
}
