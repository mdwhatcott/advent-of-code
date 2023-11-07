package starter

import (
	"fmt"
	"testing"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/grid"
	_ "github.com/mdwhatcott/must/must"
	"github.com/mdwhatcott/testing/should"
)

const TODO = -1

var (
	inputLines  = inputs.Read(2018, 15).Lines()
	sampleLines = []string{
		fmt.Sprint(TODO),
	}
)

func Test(t *testing.T) {
	should.So(t, Part1(sampleLines), should.Equal, TODO)
	should.So(t, Part1(inputLines), should.Equal, TODO)

	should.So(t, Part2(sampleLines), should.Equal, TODO)
	should.So(t, Part2(inputLines), should.Equal, TODO)
}

func Part1(lines []string) any {
	return TODO
}

func Part2(lines []string) any {
	return TODO
}

////////////////////////////////////////////////////////////////////////////////

func TestParseWorld(t *testing.T) {
	lines := []string{
		"#####",
		"#...#",
		"#...#",
		"#...#",
		"#####",
	}
	world := ParseWorld(lines)
	should.So(t, world, should.Equal, set.Of(
		grid.NewPoint(1, 1),
		grid.NewPoint(1, 2),
		grid.NewPoint(1, 3),

		grid.NewPoint(2, 1),
		grid.NewPoint(2, 2),
		grid.NewPoint(2, 3),

		grid.NewPoint(3, 1),
		grid.NewPoint(3, 2),
		grid.NewPoint(3, 3),
	))
}
func TestParseCharacters(t *testing.T) {
	lines := []string{
		"#####",
		"#.G.#",
		"#E..#",
		"#..G#",
		"#####",
	}
	characters := ParseCharacters(lines)
	should.So(t, characters, should.Equal, []*Character{
		NewCharacter('G', 2, 1),
		NewCharacter('E', 1, 2),
		NewCharacter('G', 3, 3),
	})
}
func TestAssociateCharacters(t *testing.T) {
	g1 := NewCharacter('G', 1, 1)
	g2 := NewCharacter('G', 2, 2)
	e1 := NewCharacter('E', 3, 3)
	e2 := NewCharacter('E', 4, 4)

	AssociateCharacters(g1, g2, e1, e2)

	should.So(t, g1.targets, should.Equal, []*Character{e1, e2})
	should.So(t, g2.targets, should.Equal, []*Character{e1, e2})
	should.So(t, e1.targets, should.Equal, []*Character{g1, g2})
	should.So(t, e2.targets, should.Equal, []*Character{g1, g2})
}
