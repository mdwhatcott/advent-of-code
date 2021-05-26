package advent

import (
	"strings"

	"advent/lib/util"
)

func Part1() interface{} {
	world := ParseInitialWorld(util.InputString())
	world.Boot()
	return len(world)
}

func Part2() interface{} {
	return nil
}

type P struct {
	X int
	Y int
	Z int
}

func (p P) Neighbors2d() (all []P) {
	return append(all,
		p.North(),
		p.South(),
		p.East(),
		p.West(),
		p.NorthWest(),
		p.NorthEast(),
		p.SouthWest(),
		p.SouthEast(),
	)
}
func (p P) Neighbors3d() (all []P) {
	all = append(all, p.Neighbors2d()...)
	all = append(all, p.Above())
	all = append(all, p.Above().Neighbors2d()...)
	all = append(all, p.Below())
	all = append(all, p.Below().Neighbors2d()...)
	return all
}
func (p P) Above() P     { return Point(p.X, p.Y, p.Z+1) }
func (p P) Below() P     { return Point(p.X, p.Y, p.Z-1) }
func (p P) North() P     { return Point(p.X, p.Y+1, p.Z) }
func (p P) South() P     { return Point(p.X, p.Y-1, p.Z) }
func (p P) East() P      { return Point(p.X-1, p.Y, p.Z) }
func (p P) West() P      { return Point(p.X+1, p.Y, p.Z) }
func (p P) NorthWest() P { return Point(p.X-1, p.Y+1, p.Z) }
func (p P) NorthEast() P { return Point(p.X+1, p.Y+1, p.Z) }
func (p P) SouthWest() P { return Point(p.X-1, p.Y-1, p.Z) }
func (p P) SouthEast() P { return Point(p.X+1, p.Y-1, p.Z) }

func Point(x, y, z int) P {
	return P{
		X: x,
		Y: y,
		Z: z,
	}
}

type World map[P]struct{}

func ParseInitialWorld(s string) World {
	world := make(World)
	for y, line := range strings.Fields(s) {
		for x, char := range line {
			if char == '#' {
				world.Set(Point(x, y, 0), true)
			}
		}
	}
	return world
}

func (this World) Set(p P, activate bool) {
	if activate {
		this[p] = struct{}{}
	} else {
		delete(this, p)
	}
}
func (this World) IsActive(p P) int {
	_, found := this[p]
	if found {
		return 1
	}
	return 0
}

func (this World) Boot() {
	for x := 0; x < 6; x++ {
		this.Cycle()
	}
}

func (this World) Cycle() {
	// Establish upper/lower bounds of active world state.
	var xMin, yMin, zMin int
	var xMax, yMax, zMax int
	for p := range this {
		if p.X < xMin {
			xMin = p.X
		} else if p.X > xMax {
			xMax = p.X
		}
		if p.Y < yMin {
			yMin = p.Y
		} else if p.Y > yMax {
			yMax = p.Y
		}
		if p.Z < zMin {
			zMin = p.Z
		} else if p.Z > zMax {
			zMax = p.Z
		}
	}

	// Determine upcoming state of all points within bounds.
	// The inspection must be extended to just beyond the bounds.
	upcoming := make(map[P]bool)
	for x := xMin - 1; x <= xMax+1; x++ {
		for y := yMin - 1; y <= yMax+1; y++ {
			for z := zMin - 1; z <= zMax+1; z++ {
				p := Point(x, y, z)
				active := this.countActiveNeighbors(p)
				if this.IsActive(p) > 0 {
					upcoming[p] = active == 2 || active == 3
				} else {
					upcoming[p] = active == 3
				}
			}
		}
	}

	// Deactivate and activate cells according to calculations.
	for p, state := range upcoming {
		this.Set(p, state)
	}
}

func (this World) countActiveNeighbors(p P) (active int) {
	for _, neighbor := range p.Neighbors3d() {
		active += this.IsActive(neighbor)
	}
	return active
}
