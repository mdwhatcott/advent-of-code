package advent

import (
	"fmt"
	"strings"
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"

	"advent/lib/util"
)

//lines := []string{"939", "7,13,x,x,59,x,31,19"}
// 29,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,37,x,x,x,x,x,653,x,x,x,x,x,x,x,x,x,x,x,x,13,x,x,x,17,x,x,x,x,x,23,x,x,x,x,x,x,x,823,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,19

func TestA(t *testing.T) {
	s := 1
	busses := util.InputLines()[1]
	t.Log(busses)
	for x, num := range strings.Split(busses, ",") {
		if num == "x" {
			continue
		}
		s *= util.ParseInt(num)
		fmt.Println(x, num, s)
	}
}

//100000000000000
//2283338533368659

func TestMaxBusIndex(t *testing.T) {
	assert := assertions.New(t)
	assert.So(MaxIndex(loadBusses()...), should.Equal, 60)
}

func TestBusCheck(t *testing.T) {
	assert := assertions.New(t)
	assert.So(check([]int{67, 7, 59, 61}, 0, 754018), should.BeTrue)
	assert.So(check([]int{67, 0, 7, 59, 61}, 0, 779210), should.BeTrue)
	assert.So(check([]int{67, 7, 0, 59, 61}, 0, 1261476), should.BeTrue)
	assert.So(check([]int{1789, 37, 47, 1889}, 0, 1202161486), should.BeTrue)

	assert.So(check([]int{67, 7, 59, 61}, 4, 754022), should.BeTrue)     // using non-zero index and offset timestamp
	assert.So(check([]int{68, 0, 7, 59, 61}, 0, 779210), should.BeFalse) // tampered with
	assert.So(check(loadBusses(), 0, 110727917695340), should.BeTrue)
}
