package advent

import (
	"strings"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestSlopeFixture(t *testing.T) {
	gunit.Run(new(SlopeFixture), t)
}

type SlopeFixture struct {
	*gunit.Fixture
}

func (this *SlopeFixture) Test() {
	forest := Forest(strings.Split(exampleInput, "\n"))
	this.So(forest.CountTreesHitOnDescent(1, 1), should.Equal, 2)
	this.So(forest.CountTreesHitOnDescent(1, 3), should.Equal, 7)
	this.So(forest.CountTreesHitOnDescent(1, 5), should.Equal, 3)
	this.So(forest.CountTreesHitOnDescent(1, 7), should.Equal, 4)
	this.So(forest.CountTreesHitOnDescent(2, 1), should.Equal, 2)
}

var exampleInput = strings.TrimSpace(`
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`)
