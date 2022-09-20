package day19

import (
	"fmt"
	"log"
	"strings"

	"github.com/mdwhatcott/go-collections/queue"
	"github.com/mdwhatcott/go-collections/set"

	"advent/lib/util"
)

const ROTs = 24
const minNeighbors = 12

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

type Cube struct {
	scanner   XYZ
	beacons   set.Set[XYZ]
	rotations [ROTs]set.Set[XYZ]
}

func PrepareCubes(beaconGroups [][]XYZ) (results []*Cube) {
	for _, group := range beaconGroups {
		results = append(results, &Cube{
			scanner:   XYZ{},
			beacons:   set.From[XYZ](group...),
			rotations: prepareRotations(group...),
		})
	}
	return results
}
func prepareRotations(group ...XYZ) (results [ROTs]set.Set[XYZ]) {
	for r := 0; r < ROTs; r++ {
		results[r] = set.From[XYZ](RotateMany(r, group...)...)
	}
	return results
}
func RotateMany(r int, group ...XYZ) (results []XYZ) {
	for _, item := range group {
		results = append(results, item.Rot(r))
	}
	return results
}
func ShiftMany(diff XYZ, group ...XYZ) (results []XYZ) {
	for _, item := range group {
		results = append(results, item.Add(diff))
	}
	return results
}
func TryAlign(a, b *Cube) (c *Cube, ok bool) {
	for A := range a.beacons {
		for _, rotation := range b.rotations {
			for B := range rotation {
				diff := A.Sub(B)
				BB := set.From[XYZ](ShiftMany(diff, rotation.Slice()...)...)
				shared := a.beacons.Intersection(BB)
				if shared.Len() >= minNeighbors {
					return &Cube{scanner: b.scanner.Add(diff), beacons: BB}, true
				}
			}
		}
	}
	return nil, false
}
func AlignAll(cubes []*Cube) (results []*Cube) {
	aligned := set.From[*Cube](cubes[0])
	misaligned := queue.New[*Cube](0)
	for _, c := range cubes[1:] {
		misaligned.Enqueue(c)
	}
	for !misaligned.Empty() {
		for a := range aligned {
			b := misaligned.Dequeue()
			c, ok := TryAlign(a, b)
			if ok {
				aligned.Add(c)
				log.Print("Progress:", aligned.Len(), "/", len(cubes))
				break
			} else {
				misaligned.Enqueue(b)
			}
		}
	}
	return aligned.Slice()
}
func GatherBeacons(cubes ...*Cube) set.Set[XYZ] {
	result := set.New[XYZ](0)
	for _, cube := range cubes {
		for beacon := range cube.beacons {
			result.Add(beacon)
		}
	}
	return result
}
func GatherScanners(cubes ...*Cube) (result []XYZ) {
	for _, cube := range cubes {
		result = append(result, cube.scanner)
	}
	return result
}
func MaxDistanceBetween(items []XYZ) (max int) {
	for _, a := range items {
		for _, b := range items {
			distance := a.Dist(b)
			if distance > max {
				max = distance
			}
		}
	}
	return max
}
