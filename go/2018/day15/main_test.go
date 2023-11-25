package day15

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/testing/should"
)

func Test(t *testing.T) {
	if testing.Short() {
		t.Skip("long-running test (duration of last run on 2023-11-25: 1m10s)")
	}
	inputLines := inputs.Read(2018, 15).Lines()
	should.So(t, Part1(t, inputLines), should.Equal, 261855)
	should.So(t, Part2(t, inputLines), should.Equal, 59568)
}
func Part1(t *testing.T, lines []string) any {
	t.Log("Before:\n" + strings.Join(lines, "\n"))
	units, walls := ParseCave(lines)
	rounds, health, steps := BeverageBanditsBattle(units, walls)
	t.Logf("Rounds: %d; Health: %d; Result: %d\n%s", rounds, health, rounds*health, funcy.Last(steps))
	return rounds * health
}
func Part2(t *testing.T, lines []string) any {
	t.Log("Before:\n" + strings.Join(lines, "\n"))
	rounds, health, final := BoostedBeverageBanditsBattle(lines)
	t.Logf("Rounds: %d; Health: %d; Result: %d\n%s", rounds, health, rounds*health, final)
	return rounds * health
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
		units = MoveAll(units, walls)
		t.Run(fmt.Sprintf("After round %d", x), func(t *testing.T) {
			rendered := RenderCave(units, walls)
			should.So(t, strings.Split(rendered, "\n"), should.Equal, rounds[x])
			t.Log("Expected:\n" + strings.Join(rounds[x], "\n"))
			t.Log("Actual:\n" + rendered)
		})
	}
}
func TestFullSimulations(t *testing.T) {
	tests := []struct {
		name  string
		input []string

		part1WantRounds    int
		part1WantHealth    int
		part1WantRendering string

		part2WantRounds    int
		part2WantHealth    int
		part2WantRendering string
	}{
		{
			name: "A",
			input: []string{
				"#######",
				"#.G...#   G(200)",
				"#...EG#   E(200), G(200)",
				"#.#.#G#   G(200)",
				"#..G#E#   G(200), E(200)",
				"#.....#",
				"#######",
			},
			part1WantRounds: 47,
			part1WantHealth: 200 + 131 + 59 + 200,
			part1WantRendering: strings.Join([]string{
				"#######",
				"#G....#   G(200)",
				"#.G...#   G(131)",
				"#.#.#G#   G(59)",
				"#...#.#",
				"#....G#   G(200)",
				"#######",
			}, "\n"),
			part2WantRounds: 29,
			part2WantHealth: 172,
			part2WantRendering: strings.Join([]string{
				"#######",
				"#..E..#   E(158)",
				"#...E.#   E(14)",
				"#.#.#.#",
				"#...#.#",
				"#.....#",
				"#######",
			}, "\n"),
		},
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
			part1WantRounds: 37,
			part1WantHealth: 982,
			part1WantRendering: strings.Join([]string{
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
			part1WantRounds: 46,
			part1WantHealth: 859,
			part1WantRendering: strings.Join([]string{
				"#######",
				"#.E.E.#   E(164), E(197)",
				"#.#E..#   E(200)",
				"#E.##.#   E(98)",
				"#.E.#.#   E(200)",
				"#...#.#",
				"#######",
			}, "\n"),
			part2WantRounds: 33,
			part2WantHealth: 948,
			part2WantRendering: strings.Join([]string{
				"#######",
				"#.E.E.#   E(200), E(23)",
				"#.#E..#   E(200)",
				"#E.##E#   E(125), E(200)",
				"#.E.#.#   E(200)",
				"#...#.#",
				"#######",
			}, "\n"),
		},
		{
			name: "D",
			input: []string{
				"#######",
				"#E.G#.#",
				"#.#G..#",
				"#G.#.G#",
				"#G..#.#",
				"#...E.#",
				"#######",
			},
			part1WantRounds: 35,
			part1WantHealth: 793,
			part1WantRendering: strings.Join([]string{
				"#######",
				"#G.G#.#   G(200), G(98)",
				"#.#G..#   G(200)",
				"#..#..#",
				"#...#G#   G(95)",
				"#...G.#   G(200)",
				"#######",
			}, "\n"),
			part2WantRounds: 37,
			part2WantHealth: 94,
			part2WantRendering: strings.Join([]string{
				"#######",
				"#.E.#.#   E(8)",
				"#.#E..#   E(86)",
				"#..#..#",
				"#...#.#",
				"#.....#",
				"#######",
			}, "\n"),
		},
		{
			name: "E",
			input: []string{
				"#######",
				"#.E...#",
				"#.#..G#",
				"#.###.#",
				"#E#G#G#",
				"#...#G#",
				"#######",
			},
			part1WantRounds: 54,
			part1WantHealth: 536,
			part1WantRendering: strings.Join([]string{
				"#######",
				"#.....#",
				"#.#G..#   G(200)",
				"#.###.#",
				"#.#.#.#",
				"#G.G#G#   G(98), G(38), G(200)",
				"#######",
			}, "\n"),
			part2WantRounds: 39,
			part2WantHealth: 166,
			part2WantRendering: strings.Join([]string{
				"#######",
				"#...E.#   E(14)",
				"#.#..E#   E(152)",
				"#.###.#",
				"#.#.#.#",
				"#...#.#",
				"#######",
			}, "\n"),
		},
		{
			name: "F",
			input: []string{
				"#########",
				"#G......#",
				"#.E.#...#",
				"#..##..G#",
				"#...##..#",
				"#...#...#",
				"#.G...G.#",
				"#.....G.#",
				"#########",
			},
			part1WantRounds: 20,
			part1WantHealth: 937,
			part1WantRendering: strings.Join([]string{
				"#########",
				"#.G.....#   G(137)",
				"#G.G#...#   G(200), G(200)",
				"#.G##...#   G(200)",
				"#...##..#",
				"#.G.#...#   G(200)",
				"#.......#",
				"#.......#",
				"#########",
			}, "\n"),
			part2WantRounds: 30,
			part2WantHealth: 38,
			part2WantRendering: strings.Join([]string{
				"#########",
				"#.......#",
				"#.E.#...#   E(38)",
				"#..##...#",
				"#...##..#",
				"#...#...#",
				"#.......#",
				"#.......#",
				"#########",
			}, "\n"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Run("part1", func(t *testing.T) {
				log.SetOutput(tWriter{t})
				units, walls := ParseCave(test.input)
				rounds, health, steps := BeverageBanditsBattle(units, walls)
				should.So(t, rounds, should.Equal, test.part1WantRounds)
				should.So(t, health, should.Equal, test.part1WantHealth)
				should.So(t, funcy.Last(steps), should.Equal, test.part1WantRendering)
				for s, step := range steps {
					t.Run(fmt.Sprint(s), func(t *testing.T) { t.Log("\n" + step) })
				}
				t.Logf("Expected after round %d:\n%s", test.part1WantRounds, test.part1WantRendering)
				t.Logf("Actual   after round %d:\n%s", rounds, steps[rounds])
			})
			if test.part2WantRounds == 0 {
				return
			}
			t.Run("part2", func(t *testing.T) {
				log.SetOutput(tWriter{t})
				rounds, health, final := BoostedBeverageBanditsBattle(test.input)
				should.So(t, rounds, should.Equal, test.part2WantRounds)
				should.So(t, health, should.Equal, test.part2WantHealth)
				should.So(t, final, should.Equal, test.part2WantRendering)
				t.Logf("Expected after round %d:\n%s", test.part2WantRounds, test.part2WantRendering)
				t.Logf("Actual   after round %d:\n%s", rounds, final)
			})
		})
	}
}
