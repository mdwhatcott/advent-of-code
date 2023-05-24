package advent

import (
	"github.com/mdwhatcott/advent-of-code-go-lib/maths"
)

var ( // example 1
	e1a = []int{-1, 0, 2}
	e1b = []int{2, -10, -7}
	e1c = []int{4, -8, 8}
	e1d = []int{3, 5, -1}
)

var ( // example 2
	e2a = []int{-8, -10, 0}
	e2b = []int{5, 5, 10}
	e2c = []int{2, -7, 3}
	e2d = []int{9, -8, -3}
)

var ( // part 1
	inputA = []int{-9, -1, -1}
	inputB = []int{2, 9, 5}
	inputC = []int{10, 18, -12}
	inputD = []int{-6, 15, -7}
)

func CalculateCombinedEnergy(steps int, a, b, c, d []int) interface{} {
	A := NewMoon(a...)
	B := NewMoon(b...)
	C := NewMoon(c...)
	D := NewMoon(d...)

	A.PairWith(B, C, D)
	B.PairWith(C, D, A)
	C.PairWith(D, A, B)
	D.PairWith(A, B, C)

	for x := 0; x < steps; x++ {
		aax, aay, aaz := A.CalculateUpcomingVelocityChanges()
		bbx, bby, bbz := B.CalculateUpcomingVelocityChanges()
		ccx, ccy, ccz := C.CalculateUpcomingVelocityChanges()
		ddx, ddy, ddz := D.CalculateUpcomingVelocityChanges()

		A.Move(aax, aay, aaz)
		B.Move(bbx, bby, bbz)
		C.Move(ccx, ccy, ccz)
		D.Move(ddx, ddy, ddz)
	}

	return A.TotalEnergy() +
		B.TotalEnergy() +
		C.TotalEnergy() +
		D.TotalEnergy()
}

func CalculatePeriods(a, b, c, d []int) (px, py, pz int) {
	A := NewMoon(a...)
	B := NewMoon(b...)
	C := NewMoon(c...)
	D := NewMoon(d...)

	A.PairWith(B, C, D)
	B.PairWith(C, D, A)
	C.PairWith(D, A, B)
	D.PairWith(A, B, C)

	ix := []int{a[0], b[0], c[0], d[0]}
	iy := []int{a[1], b[1], c[1], d[1]}
	iz := []int{a[2], b[2], c[2], d[2]}

	for step := 0; ; step++ {
		if step > 0 && px == 0 && A.x == ix[0] && B.x == ix[1] && C.x == ix[2] && D.x == ix[3] {
			px = step - px + 1
		}

		if step > 0 && py == 0 && A.y == iy[0] && B.y == iy[1] && C.y == iy[2] && D.y == iy[3] {
			py = step - py + 1
		}

		if step > 0 && pz == 0 && A.z == iz[0] && B.z == iz[1] && C.z == iz[2] && D.z == iz[3] {
			pz = step - pz + 1
		}

		if px > 0 && py > 0 && pz > 0 {
			return px, py, pz
		}

		aax, aay, aaz := A.CalculateUpcomingVelocityChanges()
		bbx, bby, bbz := B.CalculateUpcomingVelocityChanges()
		ccx, ccy, ccz := C.CalculateUpcomingVelocityChanges()
		ddx, ddy, ddz := D.CalculateUpcomingVelocityChanges()

		A.Move(aax, aay, aaz)
		B.Move(bbx, bby, bbz)
		C.Move(ccx, ccy, ccz)
		D.Move(ddx, ddy, ddz)
	}
}

func CalculatePeriodIntersection(x int, y int, z int) int {
	max := maths.Max(x, y, z)
	for q := max; q < (x * y * z); q += max {
		if q%x == 0 && q%y == 0 && q%z == 0 {
			return q
		}
	}
	panic("nope")
}
