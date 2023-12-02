package day17

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mdwhatcott/go-set/v2/set"
)

func Part1() interface{} {
	return nil
}

func Part2() interface{} {
	return nil
}

type Point struct{ X, Y int }

func XY(x, y int) Point { return Point{X: x, Y: y} }

func ParseInput(lines []string) set.Set[Point] {
	result := set.Of[Point]()
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			break
		}
		line = strings.Replace(line, ",", "", 1)
		fields := strings.Fields(line)
		sort.Strings(fields)
		rawX := fields[0][2:]
		rawY := fields[1][2:]

		if strings.Contains(rawX, "..") {
			xs := strings.Split(rawX, "..")
			x1, _ := strconv.Atoi(xs[0])
			x2, _ := strconv.Atoi(xs[1])
			y, _ := strconv.Atoi(rawY)
			for x := x1; x <= x2; x++ {
				result.Add(XY(x, y))
			}
		} else {
			ys := strings.Split(rawY, "..")
			y1, _ := strconv.Atoi(ys[0])
			y2, _ := strconv.Atoi(ys[1])
			x, _ := strconv.Atoi(rawX)
			for y := y1; y <= y2; y++ {
				result.Add(XY(x, y))
			}
		}
	}
	return result
}
func Display(clay, water set.Set[Point]) string {
	mixture := clay.Intersection(water)
	if len(mixture) > 0 {
		return "WATER AND CLAY MIXTURE: " + fmt.Sprint(mixture)
	}
	all := clay.Slice()
	var small, large = all[0], all[0]
	for _, c := range all[1:] {
		if c.X < small.X {
			small.X = c.X
		}
		if c.X > large.X {
			large.X = c.X
		}
		if c.Y < small.Y {
			small.Y = c.Y
		}
		if c.Y > large.Y {
			large.Y = c.Y
		}
	}
	var builder strings.Builder
	for y := 0; y < large.Y+1; y++ {
		for x := small.X; x < large.X+1; x++ {
			if clay.Contains(XY(x, y)) {
				builder.WriteString("#")
			} else if water.Contains(XY(x, y)) {
				builder.WriteString("~")
			} else {
				builder.WriteString(".")
			}
		}
		builder.WriteString("\n")
	}
	return strings.TrimSpace(builder.String())
}

func DumpWater(clay set.Set[Point]) set.Set[Point] {
	limit := MaxY(clay)
	water := set.Of[Point]()
	flow := []Point{XY(500, 1)}

	for len(flow) > 0 {
		drop := flow[0]
		flow = flow[1:]

		if clay.Contains(drop) {
			continue
		}
		if water.Contains(drop) {
			continue
		}
		if drop.Y < 1 {
			continue
		}

		up := XY(drop.X, drop.Y-1)
		down := XY(drop.X, drop.Y+1)
		left := XY(drop.X-1, drop.Y-1)
		right := XY(drop.X+1, drop.Y-1)
		if clay.Contains(down) || water.Contains(down) {
			water.Add(drop)
			flow = append(flow, left, right, up)
			continue
		}
		if down.Y > limit {
			continue
		}
		flow = append(flow, down)
	}
	return water
}

func MaxY(cave set.Set[Point]) (result int) {
	for c := range cave {
		if c.Y > result {
			result = c.Y
		}
	}
	return result
}
