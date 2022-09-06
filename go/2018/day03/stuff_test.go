package day03

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestStuffFixture(t *testing.T) {
	should.Run(&StuffFixture{T: should.New(t)}, should.Options.UnitTests())
}

type StuffFixture struct {
	*should.T
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
	this.So(undisputed, should.Equal, []int{3})
}
