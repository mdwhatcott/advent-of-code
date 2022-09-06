package day19

import (
	"testing"

	"github.com/mdwhatcott/go-collections/set"
	"github.com/mdwhatcott/testing/should"
)

func TestGeometrySuite(t *testing.T) {
	should.Run(&GeometrySuite{T: should.New(t)}, should.Options.UnitTests())
}

type GeometrySuite struct {
	*should.T
}

func (this *GeometrySuite) TestRotations() {
	var all []Point
	point := NewPoint(5, 6, -4)
	for x := 0; x < 24; x++ {
		rotated := Rotate(x, point)
		this.Log(rotated)
		all = append(all, rotated)
	}
	this.So(len(set.From[Point](all...)), should.Equal, 24)
}
func (this *GeometrySuite) TestRotateMany() {
	var rotationGroups []set.Set[Point]
	for x := 0; x < 24; x++ {
		rotated := RotateAll(x, NewPoint(5, 6, -4), NewPoint(8, 0, 7))
		this.Log("rotated:", rotated)
		rotationGroups = append(rotationGroups, set.From[Point](rotated...))
	}

	expected := ParseScannerReports(exampleRotations)
	success := 0
	for _, scanner := range expected {
		this.Log("expected:", scanner)
		beacons := set.From[Point](scanner[4:]...)
		for _, group := range rotationGroups {
			if group.Equal(beacons) {
				success++
			}
		}
	}
	this.So(success, should.Equal, len(expected))
}
