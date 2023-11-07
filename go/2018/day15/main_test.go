package starter

import (
	"fmt"
	"strings"
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
	t.Log(ParseWorld(inputLines).String())
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

func TestParseMap(t *testing.T) {
	lines := []string{
		"#####",
		"#...#",
		"#...#",
		"#...#",
		"#####",
	}
	cave := ParseCaveMap(lines)
	should.So(t, cave, should.Equal, set.Of(
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
func TestParseUnits(t *testing.T) {
	lines := []string{
		"#####",
		"#.G.#",
		"#E..#",
		"#..G#",
		"#####",
	}
	units := ParseUnits(lines)
	should.So(t, units, should.Equal, []*Unit{
		NewUnit('G', 2, 1),
		NewUnit('E', 1, 2),
		NewUnit('G', 3, 3),
	})
}
func TestAssociateEnemyCharacters(t *testing.T) {
	g1 := NewUnit('G', 1, 1)
	g2 := NewUnit('G', 2, 2)
	e1 := NewUnit('E', 3, 3)
	e2 := NewUnit('E', 4, 4)

	AssociateEnemyUnits(g1, g2, e1, e2)

	should.So(t, g1.targets, should.Equal, []*Unit{e1, e2})
	should.So(t, g2.targets, should.Equal, []*Unit{e1, e2})
	should.So(t, e1.targets, should.Equal, []*Unit{g1, g2})
	should.So(t, e2.targets, should.Equal, []*Unit{g1, g2})
}
func TestParseAndPrintWorld1(t *testing.T) {
	lines := []string{
		"#######",
		"#.G...#   G(200)",
		"#...EG#   E(200), G(200)",
		"#.#.#G#   G(200)",
		"#..G#E#   G(200), E(200)",
		"#.....#",
		"#######",
	}
	world := ParseWorld(lines)
	actual := world.String()
	expected := "\n" + strings.Join(lines, "\n") + "\n"
	should.So(t, actual, should.Equal, expected)
	t.Log(actual)
}
