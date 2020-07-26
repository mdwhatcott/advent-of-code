package advent

import (
	"fmt"

	"advent/lib/util"
)

type Moon struct {
	x int
	y int
	z int

	dx int
	dy int
	dz int

	others []*Moon
}

func NewMoon(xyz ...int) *Moon {
	return &Moon{
		x: xyz[0],
		y: xyz[1],
		z: xyz[2],
	}
}

func (this *Moon) PairWith(those ...*Moon) {
	for _, that := range those {
		this.others = append(this.others, that)
	}
}

func (this *Moon) CalculateUpcomingVelocityChanges() (ddx, ddy, ddz int) {
	for _, other := range this.others {
		ddx += comp(other.x, this.x)
		ddy += comp(other.y, this.y)
		ddz += comp(other.z, this.z)
	}
	return ddx, ddy, ddz
}

func (this *Moon) Move(ddx, ddy, ddz int) {
	this.updateVelocity(ddx, ddy, ddz)
	this.updatePosition()
}

func (this *Moon) updateVelocity(ddx, ddy, ddz int) {
	this.dx += ddx
	this.dy += ddy
	this.dz += ddz
}
func (this *Moon) updatePosition() {
	this.x += this.dx
	this.y += this.dy
	this.z += this.dz
}

func (this *Moon) TotalEnergy() int {
	potential := util.Abs(this.x) + util.Abs(this.y) + util.Abs(this.z)
	kinetic := util.Abs(this.dx) + util.Abs(this.dy) + util.Abs(this.dz)
	return potential * kinetic
}

func (this *Moon) String() string {
	return fmt.Sprintf(
		"(%d, %d, %d) [%d, %d, %d] = %d",
		this.x, this.y, this.z,
		this.dx, this.dy, this.dz,
		this.TotalEnergy(),
	)
}

func comp(a, b int) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}
