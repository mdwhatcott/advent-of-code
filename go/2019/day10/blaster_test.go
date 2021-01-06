package advent

import (
	"fmt"
	"strings"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestBlasterFixture(t *testing.T) {
	gunit.Run(new(BlasterFixture), t)
}

type BlasterFixture struct {
	*gunit.Fixture
	field   AsteroidField
	blaster *Blaster
}

var part2Example1 = strings.Split(strings.TrimSpace(""+
	`
.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....X...###..
..#.#.....#....##

`), "\n")

func (this *BlasterFixture) TestBlastCircle1() {
	this.field = removeOrigin(offsetField(scanField(part2Example1), -8, -3))
	this.blaster = NewBlaster(this.field)

	this.assertTargetOrder(
		"(0,-2)",
		"(1,-3)",
		"(1,-2)",
		"(2,-3)",
		"(1,-1)",
		"(3,-2)",
		"(4,-2)",
		"(3,-1)",
		"(7,-2)",

		"(4,-1)",
		"(5,-1)",
		"(6,-1)",
		"(7,-1)",
		"(4,0)",
		"(8,1)",
		"(7,1)",
		"(2,1)",
		"(-4,1)",

		"(-6,1)",
		"(-6,0)",
		"(-8,-1)",
		"(-7,-1)",
		"(-8,-2)",
		"(-7,-2)",
		"(-3,-1)",
		"(-7,-3)",
		"(-3,-2)",

		"(-2,-2)",
		"(-2,-3)",
		"(-1,-3)",
		"(0,-3)",
		"(2,-2)",
		"(6,-3)",
		"(8,-2)",
		"(5,0)",
		"(6,0)",
	)
	this.So(this.blaster.Field(), should.BeEmpty)
}
func (this *BlasterFixture) TestBlastCircle2() {
	this.field = removeOrigin(offsetField(scanField(exampleMap5), -11, -13))
	this.blaster = NewBlaster(this.field)

	this.assertSelectedTargets(map[string]bool{
		"The 1 asteroid to be vaporized is at 11,12.":   true,
		"The 2 asteroid to be vaporized is at 12,1.":    true,
		"The 3 asteroid to be vaporized is at 12,2.":    true,
		"The 10 asteroid to be vaporized is at 12,8.":   true,
		"The 20 asteroid to be vaporized is at 16,0.":   true,
		"The 50 asteroid to be vaporized is at 16,9.":   true,
		"The 100 asteroid to be vaporized is at 10,16.": true,
		"The 199 asteroid to be vaporized is at 9,6.":   true,
		"The 200 asteroid to be vaporized is at 8,2.":   true,
		"The 201 asteroid to be vaporized is at 10,9.":  true,
		"The 299 asteroid to be vaporized is at 11,1.":  true,
	})
	this.So(this.blaster.Field(), should.BeEmpty)
}

func (this *BlasterFixture) assertTargetOrder(expected ...string) {
	for _, e := range expected {
		aim := this.blaster.Aim()
		asteroid := this.blaster.Field()[aim]
		this.So(asteroid.String(), should.Equal, e)
		this.blaster.Fire(aim)
	}
}

func (this *BlasterFixture) assertSelectedTargets(expected map[string]bool) {
	for x := 1; len(this.blaster.Field()) > 0; x++ {
		aim := this.blaster.Aim()
		asteroid := this.blaster.Field()[aim]
		asteroid.X += 11
		asteroid.Y += 13
		targeted := asteroid.String()
		targeted = strings.TrimPrefix(targeted, "(")
		targeted = strings.TrimSuffix(targeted, ")")
		formatted := fmt.Sprintf("The %d asteroid to be vaporized is at %s.", x, targeted)
		if expected[formatted] {
			delete(expected, formatted)
		}
		this.blaster.Fire(aim)
	}
	this.So(expected, should.BeEmpty)
}
