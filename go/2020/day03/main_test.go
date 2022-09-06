package advent

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestSlopeFixture(t *testing.T) {
	should.Run(&SlopeFixture{T: should.New(t)}, should.Options.UnitTests())
}

type SlopeFixture struct {
	*should.T
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
