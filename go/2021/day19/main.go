package day19

import (
	"fmt"
	"strings"

	"advent/lib/util"
)

type XYZ struct{ x, y, z int }

func NewXYZ(x, y, z int) XYZ { return XYZ{x: x, y: y, z: z} }
func (a XYZ) Add(b XYZ) XYZ  { return NewXYZ(a.x+b.x, a.y+b.y, a.z+b.z) }
func (a XYZ) Sub(b XYZ) XYZ  { return NewXYZ(a.x-b.x, a.y-b.y, a.z-b.z) }
func (a XYZ) Rot(i int) XYZ  { return NewXYZ(spins[i%4](faces[i/4](a.x, a.y, a.z))) }
func (a XYZ) Dist(b XYZ) int { d := a.Sub(b); return util.Abs(d.x) + util.Abs(d.y) + util.Abs(d.z) }
func (a XYZ) String() string { return fmt.Sprintf("(%v, %v, %v)", a.x, a.y, a.z) }

var faces = []func(x, y, z int) (X, Y, Z int){
	func(x, y, z int) (X, Y, Z int) { return x, y, z },
	func(x, y, z int) (X, Y, Z int) { return x, -y, -z },
	func(x, y, z int) (X, Y, Z int) { return x, -z, y },
	func(x, y, z int) (X, Y, Z int) { return -y, -z, x },
	func(x, y, z int) (X, Y, Z int) { return y, -z, -x },
	func(x, y, z int) (X, Y, Z int) { return -x, -z, -y },
}
var spins = []func(x, y, z int) (X, Y, Z int){
	func(x, y, z int) (X, Y, Z int) { return x, y, z },
	func(x, y, z int) (X, Y, Z int) { return -y, x, z },
	func(x, y, z int) (X, Y, Z int) { return -x, -y, z },
	func(x, y, z int) (X, Y, Z int) { return y, -x, z },
}

func ParseScannerReports(reports string) (results [][]XYZ) {
	rawGroups := strings.Split(reports, "\n\n")
	for _, rawGroup := range rawGroups {
		var group []XYZ
		for _, line := range strings.Split(rawGroup, "\n")[1:] {
			fields := strings.Split(line, ",")
			x := util.ParseInt(fields[0])
			y := util.ParseInt(fields[1])
			z := util.ParseInt(fields[2])
			group = append(group, NewXYZ(x, y, z))
		}
		results = append(results, group)
	}
	return results
}
