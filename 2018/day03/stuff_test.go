package day03

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

const toy = `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`

func (this *StuffFixture) TestToy() {
	fabric := plotClaimsOnFabric(parseClaims(toy))

	conflict := 0
	for point := range fabric {
		if len(fabric[point].claims) > 1 {
			conflict++
		}
	}
	this.So(conflict, should.Equal, 4)
}

func (this *StuffFixture) TestFindNoConflicts() {
	var undisputed []int
	claims := parseClaims(toy)
	fabric := plotClaimsOnFabric(claims)
	for id, claim := range claims {
		if fabric.IsUndisputed(claim) {
			undisputed = append(undisputed, id+1)
		}
	}
	this.So(undisputed, should.Resemble, []int{3})
}

