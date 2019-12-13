package advent

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuffFixture(t *testing.T) {
	gunit.Run(new(StuffFixture), t)
}

type StuffFixture struct {
	*gunit.Fixture
}

const sampleInput = "123456789012"

func (this *StuffFixture) TestOperations_SampleImage() {
	image := ParseImage(sampleInput, 3, 2)
	this.So(image.LayerCount(), should.Equal, 2)
	this.So(image.PixelCount(), should.Equal, 3*2)
	this.So(image.LayerSegment(0), should.Equal, "123456")
	this.So(image.LayerSegment(1), should.Equal, "789012")
}
