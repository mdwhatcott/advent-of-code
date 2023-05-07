package advent

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func TestStuffFixture(t *testing.T) {
	should.Run(&StuffFixture{T: should.New(t)}, should.Options.UnitTests())
}

type StuffFixture struct {
	*should.T
}

const sampleInputA = "123456789012" // 3x2

func (this *StuffFixture) TestOperations_SampleImage() {
	image := ParseImage(sampleInputA, 3, 2)
	this.So(image.LayerCount(), should.Equal, 2)
	this.So(image.PixelCount(), should.Equal, 3*2)
	this.So(image.LayerSegment(0), should.Equal, "123456")
	this.So(image.LayerSegment(1), should.Equal, "789012")
}

const sampleInputB = "0222112222120000" // 2x2

func (this *StuffFixture) TestPixelColorResolution() {
	image := ParseImage(sampleInputB, 2, 2)
	this.So(image.RenderPixel(0, 0), should.Equal, "0")
	this.So(image.RenderPixel(1, 0), should.Equal, "1")
	this.So(image.RenderPixel(0, 1), should.Equal, "1")
	this.So(image.RenderPixel(1, 1), should.Equal, "0")
}

func (this *StuffFixture) TestFullRender_RealImage() {
	this.So(ParseImage(util.InputString(), 25, 6).RenderFull(), should.Equal, expectedImageRendering)
}

var expectedImageRendering = strings.TrimSpace(`
*   * **  ***  *   ***** 
*   **  * *  * *   *   * 
 * * *    *  *  * *   *  
  *  * ** ***    *   *   
  *  *  * * *    *  *    
  *   *** *  *   *  **** 
`)
