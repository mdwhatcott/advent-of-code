package day15

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"testing"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/testing/should"
)

const TODO = -1

func Test(t *testing.T) {
	inputLines := inputs.Read(2018, 15).Lines()
	should.So(t, Part1(inputLines), should.Equal, TODO)
	should.So(t, Part2(inputLines), should.Equal, TODO)
}
func Part1(lines []string) any {
	return TODO
}
func Part2(lines []string) any {
	return TODO
}

type tWriter struct{ *testing.T }

func (this tWriter) Write(p []byte) (n int, err error) {
	this.T.Log(strings.TrimSpace(string(p)))
	return len(p), nil
}

func TestBFS(t *testing.T) {
	cave := []string{
		"#########",
		"#.......#",
		"#.......#",
		"#....321#",
		"#......0#",
		"#.......#",
		"#.......#",
		"#.......#",
		"#########",
	}
	_, walls := ParseCave(cave)
	path := BFS(XY(7, 4), XY(5, 3), walls)
	should.So(t, path, should.Equal, []Point{XY(7, 3), XY(6, 3), XY(5, 3)})
}

func TestSort(t *testing.T) {
	points := []*Unit{
		NewUnit(9, 2, ""),
		NewUnit(0, 3, ""),
		NewUnit(4, 9, ""),
		NewUnit(0, 4, ""),
		NewUnit(7, 3, ""),
		NewUnit(8, 8, ""),
		NewUnit(8, 8, ""),
		NewUnit(7, 9, ""),
		NewUnit(6, 0, ""),
		NewUnit(9, 4, ""),
	}

	points = Sort(points)

	should.So(t, points, should.Equal, []*Unit{
		NewUnit(6, 0, ""),
		NewUnit(9, 2, ""),
		NewUnit(0, 3, ""),
		NewUnit(7, 3, ""),
		NewUnit(0, 4, ""),
		NewUnit(9, 4, ""),
		NewUnit(8, 8, ""),
		NewUnit(8, 8, ""),
		NewUnit(4, 9, ""),
		NewUnit(7, 9, ""),
	})
}
func TestParseCave(t *testing.T) {
	lines := []string{
		"#######",
		"#E..G.#",
		"#...#.#",
		"#.G.#G#",
		"#######",
	}
	units, obstacles := ParseCave(lines)
	should.So(t, units, should.Equal, []*Unit{
		NewUnit(1, 1, "E"),
		NewUnit(4, 1, "G"),
		NewUnit(2, 3, "G"),
		NewUnit(5, 3, "G"),
	})
	should.So(t, obstacles, should.Equal, set.Of(
		// top side walls
		XY(0, 0),
		XY(1, 0),
		XY(2, 0),
		XY(3, 0),
		XY(4, 0),
		XY(5, 0),
		XY(6, 0),

		// left side walls
		XY(0, 1),
		XY(0, 2),
		XY(0, 3),

		// right side walls
		XY(6, 1),
		XY(6, 2),
		XY(6, 3),

		// bottom side walls
		XY(0, 4),
		XY(1, 4),
		XY(2, 4),
		XY(3, 4),
		XY(4, 4),
		XY(5, 4),
		XY(6, 4),

		// middle walls
		XY(4, 2),
		XY(4, 3),
	))
}
func TestRenderCave(t *testing.T) {
	lines := []string{
		"#######",
		"#E..G.#",
		"#...#.#",
		"#.G.#G#",
		"#######",
	}
	cave := RenderCave(ParseCave(lines))
	should.So(t, strings.Split(cave, "\n"), should.Equal, lines)
}
func TestMoveUnit_ExampleA(t *testing.T) {
	/*
		Targets:      In range:     Reachable:    Nearest:      Chosen:
		#######       #######       #######       #######       #######
		#E..G.#       #E.?G?#       #E.@G.#       #E.!G.#       #E.+G.#
		#...#.#  -->  #.?.#?#  -->  #.@.#.#  -->  #.!.#.#  -->  #...#.#
		#.G.#G#       #?G?#G#       #@G@#G#       #!G.#G#       #.G.#G#
		#######       #######       #######       #######       #######
	*/
	units, walls := ParseCave([]string{
		"#######",
		"#E..G.#",
		"#...#.#",
		"#.G.#G#",
		"#######",
	})
	MoveUnit(units[0], units, walls)

	should.So(t, units[0].Point, should.Equal, XY(2, 1))
}
func TestMoveUnit_ExampleB(t *testing.T) {
	/*
		In range:     Nearest:      Chosen:       Distance:     Step:
		#######       #######       #######       #######       #######
		#.E...#       #.E...#       #.E...#       #4E212#       #..E..#
		#...?.#  -->  #...!.#  -->  #...+.#  -->  #32101#  -->  #.....#
		#..?G?#       #..!G.#       #...G.#       #432G2#       #...G.#
		#######       #######       #######       #######       #######
	*/
	units, walls := ParseCave([]string{
		"#######",
		"#.E...#",
		"#.....#",
		"#...G.#",
		"#######",
	})
	MoveUnit(units[0], units, walls)

	should.So(t, units[0].Point, should.Equal, XY(3, 1))
}
func TestMoveAll_SeveralRounds(t *testing.T) {
	rounds := [][]string{
		{
			"#########",
			"#G..G..G#",
			"#.......#",
			"#.......#",
			"#G..E..G#",
			"#.......#",
			"#.......#",
			"#G..G..G#",
			"#########",
		},
		{
			"#########",
			"#.G...G.#",
			"#...G...#",
			"#...E..G#",
			"#.G.....#",
			"#.......#",
			"#G..G..G#",
			"#.......#",
			"#########",
		},
		{
			"#########",
			"#..G.G..#",
			"#...G...#",
			"#.G.E.G.#",
			"#.......#",
			"#G..G..G#",
			"#.......#",
			"#.......#",
			"#########",
		},
		{
			"#########",
			"#.......#",
			"#..GGG..#",
			"#..GEG..#",
			"#G..G...#",
			"#......G#",
			"#.......#",
			"#.......#",
			"#########",
		},
	}

	log.SetFlags(0)
	log.SetOutput(tWriter{t})
	units, walls := ParseCave(rounds[0])
	for x := 1; x < len(rounds); x++ {
		log.Println("hi")
		units = MoveAll(units, walls)
		t.Run(fmt.Sprintf("After round %d", x), func(t *testing.T) {
			rendered := RenderCave(units, walls)
			should.So(t, strings.Split(rendered, "\n"), should.Equal, rounds[x])
			t.Log("Expected:\n" + strings.Join(rounds[x], "\n"))
			t.Log("Actual:\n" + rendered)
		})
	}
}

func TestSomething(t *testing.T) {
	cave := []string{
		"#########",
		"#.......#",
		"#...G...#",
		"#...E...#",
		"#.......#",
		"#.......#",
		"#.......#",
		"#.......#",
		"#########",
	}
	units, walls := ParseCave(cave)

	goblin := FilterTeam(units, "G")[0]
	start := goblin.Point
	MoveUnit(goblin, units, walls)
	should.So(t, goblin.Point, should.Equal, start)
	t.Log("\n" + RenderCave(units, walls))
}

func TestFullSimulationA1(t *testing.T) {
	caves := map[int][]string{
		0: {
			"#######",
			"#.G...#   G(200)",
			"#...EG#   E(200), G(200)",
			"#.#.#G#   G(200)",
			"#..G#E#   G(200), E(200)",
			"#.....#",
			"#######",
		},
		1: {
			"#######",
			"#..G..#   G(200)",
			"#...EG#   E(197), G(197)",
			"#.#G#G#   G(200), G(197)",
			"#...#E#   E(197)",
			"#.....#",
			"#######",
		},
		2: {
			"#######",
			"#...G.#   G(200)",
			"#..GEG#   G(200), E(188), G(194)",
			"#.#.#G#   G(194)",
			"#...#E#   E(194)",
			"#.....#",
			"#######",
		},
		23: {
			"#######",
			"#...G.#   G(200)",
			"#..G.G#   G(200), G(131)",
			"#.#.#G#   G(131)",
			"#...#E#   E(131)",
			"#.....#",
			"#######",
		},
		28: {
			"#######",
			"#G....#   G(200)",
			"#.G...#   G(131)",
			"#.#.#G#   G(116)",
			"#...#E#   E(113)",
			"#....G#   G(200)",
			"#######",
		},
		47: {
			"#######",
			"#G....#   G(200)",
			"#.G...#   G(131)",
			"#.#.#G#   G(59)",
			"#...#.#",
			"#....G#   G(200)",
			"#######",
		},
	}
	rounds, health, steps := BeverageBanditsBattle(caves[0])
	should.So(t, rounds, should.Equal, 47)
	should.So(t, health, should.Equal, 200+131+59+200)

	keys := funcy.MapKeys(caves)
	sort.Ints(keys)

	for _, checkpoint := range keys {
		t.Run(fmt.Sprintf("After round %02d", checkpoint), func(t *testing.T) {
			should.So(t, strings.Split(steps[checkpoint], "\n"), should.Equal, caves[checkpoint])
			t.Logf("Expected after round %d:\n%s", checkpoint, strings.Join(caves[checkpoint], "\n"))
			t.Logf("Actual   after round %d:\n%s", checkpoint, steps[checkpoint])
		})
	}
}

func TestFullSimulations(t *testing.T) {
	tests := []struct {
		name          string
		input         []string
		wantRounds    int
		wantHealth    int
		wantRendering string
	}{
		{
			name: "B",
			input: []string{
				"#######",
				"#G..#E#",
				"#E#E.E#",
				"#G.##.#",
				"#...#E#",
				"#...E.#",
				"#######",
			},
			wantRounds: 37,
			wantHealth: 982,
			wantRendering: strings.Join([]string{
				"#######",
				"#...#E#   E(200)",
				"#E#...#   E(197)",
				"#.E##.#   E(185)",
				"#E..#E#   E(200), E(200)",
				"#.....#",
				"#######",
			}, "\n"),
		},
		{
			name: "C",
			input: []string{
				"#######",
				"#E..EG#",
				"#.#G.E#",
				"#E.##E#",
				"#G..#.#",
				"#..E#.#",
				"#######",
			},
			wantRounds: 46,
			wantHealth: 859,
			wantRendering: strings.Join([]string{
				"#######",
				"#.E.E.#   E(164), E(197)",
				"#.#E..#   E(200)",
				"#E.##.#   E(98)",
				"#.E.#.#   E(200)",
				"#...#.#",
				"#######",
			}, "\n"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rounds, health, steps := BeverageBanditsBattle(test.input)
			should.So(t, rounds, should.Equal, test.wantRounds)
			should.So(t, health, should.Equal, test.wantHealth)
			should.So(t, steps[rounds], should.Equal, test.wantRendering)
			for s, step := range steps {
				t.Run(fmt.Sprint(s), func(t *testing.T) {
					t.Log("\n" + step)
				})
			}
			t.Logf("Expected after round %d:\n%s", test.wantRounds, test.wantRendering)
			t.Logf("Actual   after round %d:\n%s", rounds, steps[rounds])
		})
	}
}
