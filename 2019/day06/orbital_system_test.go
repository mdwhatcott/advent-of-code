package advent

import (
	"strings"
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

func (this *StuffFixture) TestSystemWithNoSatellites_ChecksumZero() {
	system := &System{Label: "COM", Magnitude: 0}
	this.So(system.OrbitalChecksum(), should.Equal, 0)
}
func (this *StuffFixture) TestSystemWithSatellites_ChecksumFortyTwo() {
	system := assembleOrbitalSystem(exampleLines)
	this.So(system.OrbitalChecksum(), should.Equal, 42)
}
func (this *StuffFixture) TestPath() {
	system := assembleOrbitalSystem(exampleLines)
	this.So(system.TracePath("F"), should.Equal, "/COM/B/C/D/E/F")
	this.So(system.TracePath("L"), should.Equal, "/COM/B/C/D/E/J/K/L")
}
func (this *StuffFixture) TestDistance() {
	system := assembleOrbitalSystem(exampleLines2)
	this.So(system.OrbitalDistance("YOU", "SAN"), should.Equal, 4)
}

var exampleLines = strings.Split(strings.TrimSpace(`
J)K
B)C
G)H
C)D
D)E
E)F
B)G
COM)B
D)I
E)J
K)L
`), "\n")

var exampleLines2 = strings.Split(strings.TrimSpace(`
J)K
B)C
G)H
C)D
D)E
E)F
B)G
COM)B
D)I
E)J
K)L
K)YOU
I)SAN
`), "\n")
