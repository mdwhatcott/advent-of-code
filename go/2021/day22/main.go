package day22

import (
	"fmt"

	"github.com/mdwhatcott/advent-of-code-go-lib/maths"
)

// inspiration:
// - https://github.com/tginsberg/advent-2021-kotlin/blob/master/src/main/kotlin/com/ginsberg/advent2021/Day22.kt
// - https://github.com/lukechampine/advent/blob/master/2021/day22.go
var part1Space = Cuboid{0,
	Range{-50, 50},
	Range{-50, 50},
	Range{-50, 50},
}

func solve(lines []string) (part1, part2 int64) {
	var results []Cuboid
	for _, line := range lines {
		current := parseCuboid(line)
		for _, already := range results {
			if current.intersects(already) {
				results = append(results, current.intersection(already))
			}
		}
		if current.sign == 1 {
			results = append(results, current)
		}
	}
	for _, c := range results {
		volume := c.volume()
		if c.intersects(part1Space) {
			part1 += volume
		}
		part2 += volume
	}
	return part1, part2
}

type Cuboid struct {
	sign    int64
	x, y, z Range
}

func (a Cuboid) intersection(b Cuboid) Cuboid {
	return Cuboid{sign: -b.sign,
		x: Range{maths.Max(a.x.lo, b.x.lo), maths.Min(a.x.hi, b.x.hi)},
		y: Range{maths.Max(a.y.lo, b.y.lo), maths.Min(a.y.hi, b.y.hi)},
		z: Range{maths.Max(a.z.lo, b.z.lo), maths.Min(a.z.hi, b.z.hi)},
	}
}
func (a Cuboid) intersects(b Cuboid) bool {
	return a.x.intersects(b.x) &&
		a.y.intersects(b.y) &&
		a.z.intersects(b.z)
}
func (c Cuboid) volume() int64 {
	return c.sign *
		(c.x.hi - c.x.lo + 1) *
		(c.y.hi - c.y.lo + 1) *
		(c.z.hi - c.z.lo + 1)
}

type Range struct {
	lo, hi int64
}

func (a Range) intersects(b Range) bool {
	return a.lo <= b.hi && a.hi >= b.lo
}

func parseCuboid(line string) (c Cuboid) {
	c.sign = int64(1)
	var onOff string
	_, _ = fmt.Sscanf(line,
		"%s x=%d..%d,y=%d..%d,z=%d..%d",
		&onOff,
		&c.x.lo, &c.x.hi,
		&c.y.lo, &c.y.hi,
		&c.z.lo, &c.z.hi,
	)
	if onOff == "off" {
		c.sign = -1
	}
	return c
}
