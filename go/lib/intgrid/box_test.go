package intgrid

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestBoxSuite(t *testing.T) {
	should.Run(&BoxSuite{T: should.New(t)}, should.Options.UnitTests())
}

type BoxSuite struct {
	*should.T
}

func (this *BoxSuite) Test() {
	box := NewBoundingBox(NewPoint(0, 0), NewPoint(2, 2))
	for x := 0; x <= 2; x++ {
		for y := 0; y <= 2; y++ {
			this.So(box.Contains(NewPoint(x, y)), should.BeTrue)
		}
	}
	this.So(box.Contains(NewPoint(-1, 0)), should.BeFalse)
	this.So(box.Contains(NewPoint(0, -1)), should.BeFalse)
	this.So(box.Contains(NewPoint(-1, -1)), should.BeFalse)
	this.So(box.Contains(NewPoint(2, 3)), should.BeFalse)
	this.So(box.Contains(NewPoint(3, 2)), should.BeFalse)
	this.So(box.Contains(NewPoint(3, 3)), should.BeFalse)
}
