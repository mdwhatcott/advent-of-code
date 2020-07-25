package advent

import (
	"fmt"

	"advent/lib/util"
)

// example 1
var (
	e1a = []int{-1, 0, 2}
	e1b = []int{2, -10, -7}
	e1c = []int{4, -8, 8}
	e1d = []int{3, 5, -1}
)

// example 2
var (
	e2a = []int{-8, -10, 0}
	e2b = []int{5, 5, 10}
	e2c = []int{2, -7, 3}
	e2d = []int{9, -8, -3}
)

// part 1
var (
	part1a = []int{-9, -1, -1}
	part1b = []int{2, 9, 5}
	part1c = []int{10, 18, -12}
	part1d = []int{-6, 15, -7}
)

func CalculateCombinedEnergy(steps int, a, b, c, d []int) interface{} {
	var (
		A = NewMoon(a...)
		B = NewMoon(b...)
		C = NewMoon(c...)
		D = NewMoon(d...)
	)
	A.Pair(B, C, D)
	B.Pair(C, D, A)
	C.Pair(D, A, B)
	D.Pair(A, B, C)

	for x := 0; x < steps; x++ {
		aax, aay, aaz := A.CalculateUpcomingVelocityChanges()
		bbx, bby, bbz := B.CalculateUpcomingVelocityChanges()
		ccx, ccy, ccz := C.CalculateUpcomingVelocityChanges()
		ddx, ddy, ddz := D.CalculateUpcomingVelocityChanges()

		A.Move(aax, aay, aaz)
		B.Move(bbx, bby, bbz)
		C.Move(ccx, ccy, ccz)
		D.Move(ddx, ddy, ddz)

		//fmt.Println(A.String())
		//fmt.Println(B.String())
		//fmt.Println(C.String())
		//fmt.Println(D.String())
		//fmt.Println()
	}

	return A.TotalEnergy() +
		B.TotalEnergy() +
		C.TotalEnergy() +
		D.TotalEnergy()
}

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

func (this *Moon) Pair(those ...*Moon) {
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

func comp(a, b int) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

func (this *Moon) String() string {
	return fmt.Sprintf(
		"(%d, %d, %d) [%d, %d, %d] = %d",
		this.x, this.y, this.z,
		this.dx, this.dy, this.dz,
		this.TotalEnergy(),
	)
}
