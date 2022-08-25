package day19

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestDay19Suite(t *testing.T) {
	suite.Run(&Day19Suite{T: suite.New(t)}, suite.Options.UnitTests())
}

type Day19Suite struct {
	*suite.T
}

func (this *Day19Suite) TestParseScannerReports() {
	beacons := ParseScannerReports(exampleBeaconReport)
	this.So(beacons, should.Equal, [][]Point{
		{
			NewPoint(0, 2, 3),
			NewPoint(4, 1, 5),
			NewPoint(3, 3, 7),
		},
		{
			NewPoint(-1, -1, 3),
			NewPoint(-5, 0, 5),
			NewPoint(-2, 1, 7),
		},
	})
}
func (this *Day19Suite) TestRotations() {
	beacons := ParseScannerReports(exampleRotations)
	this.So(beacons[0], should.Equal, []Point{
		NewPoint(-1, -1, 1),
		NewPoint(-2, -2, 2),
		NewPoint(-3, -3, 3),
		NewPoint(-2, -3, 1),
		NewPoint(5, 6, -4),
		NewPoint(8, 0, 7),
	})
}

var exampleBeaconReport = strings.TrimSpace(`
--- scanner 0 ---
0,2,3
4,1,5
3,3,7

--- scanner 1 ---
-1,-1,3
-5,0,5
-2,1,7
`)
var exampleRotations = strings.TrimSpace(`
--- scanner 0 ---
-1,-1,1
-2,-2,2
-3,-3,3
-2,-3,1
5,6,-4
8,0,7

--- scanner 0 ---
1,-1,1
2,-2,2
3,-3,3
2,-1,3
-5,4,-6
-8,-7,0

--- scanner 0 ---
-1,-1,-1
-2,-2,-2
-3,-3,-3
-1,-3,-2
4,6,5
-7,0,8

--- scanner 0 ---
1,1,-1
2,2,-2
3,3,-3
1,3,-2
-4,-6,5
7,0,8

--- scanner 0 ---
1,1,1
2,2,2
3,3,3
3,1,2
-6,-4,-5
0,7,-8
`)
