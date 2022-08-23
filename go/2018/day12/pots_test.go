package day12

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestStuffFixture(t *testing.T) {
	suite.Run(&StuffFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type StuffFixture struct {
	*suite.T
}

func (this *StuffFixture) TestSingleRule() {
	rule := ParseRule("..... => #")
	this.So(rule.IsSatisfiedBy("....."), should.BeTrue)
	this.So(rule.IsSatisfiedBy("#...."), should.BeFalse)
	this.So(rule.String(), should.Equal, "#")
}

func (this *StuffFixture) TestRowSum() {
	//                    0123456789012345678901234
	pots := NewRowOfPots("#..#.#..##......###...###")
	this.So(pots.Sum(), should.Equal, 145)
	this.So(pots.Render(), should.Equal, "#..#.#..##......###...###")
}

func (this *StuffFixture) TestRowTransformations() {
	rules := []Rule{
		ParseRule("...## => #"),
		ParseRule("..#.. => #"),
		ParseRule(".#... => #"),
		ParseRule(".#.#. => #"),
		ParseRule(".#.## => #"),
		ParseRule(".##.. => #"),
		ParseRule(".#### => #"),
		ParseRule("#.#.# => #"),
		ParseRule("#.### => #"),
		ParseRule("##.#. => #"),
		ParseRule("##.## => #"),
		ParseRule("###.. => #"),
		ParseRule("###.# => #"),
		ParseRule("####. => #"),
	}
	pots := NewRowOfPots("#..#.#..##......###...###", rules...)

	for _, text := range ExampleGenerations {
		pots.Update(pots.Scan())
		actual := strings.Trim(pots.Render(), ".")
		expected := this.extractExpectedPotGrowth(text)
		this.So(actual, should.Equal, expected)
	}
	this.So(pots.Sum(), should.Equal, 325)
}

func (this *StuffFixture) extractExpectedPotGrowth(text string) string {
	line := strings.TrimSpace(text)
	parts := strings.Fields(line)
	return strings.Trim(parts[1], ".")
}

var ExampleGenerations = []string{
	" 1: ...#...#....#.....#..#..#..#...........",
	" 2: ...##..##...##....#..#..#..##..........",
	" 3: ..#.#...#..#.#....#..#..#...#..........",
	" 4: ...#.#..#...#.#...#..#..##..##.........",
	" 5: ....#...##...#.#..#..#...#...#.........",
	" 6: ....##.#.#....#...#..##..##..##........",
	" 7: ...#..###.#...##..#...#...#...#........",
	" 8: ...#....##.#.#.#..##..##..##..##.......",
	" 9: ...##..#..#####....#...#...#...#.......",
	"10: ..#.#..#...#.##....##..##..##..##......",
	"11: ...#...##...#.#...#.#...#...#...#......",
	"12: ...##.#.#....#.#...#.#..##..##..##.....",
	"13: ..#..###.#....#.#...#....#...#...#.....",
	"14: ..#....##.#....#.#..##...##..##..##....",
	"15: ..##..#..#.#....#....#..#.#...#...#....",
	"16: .#.#..#...#.#...##...#...#.#..##..##...",
	"17: ..#...##...#.#.#.#...##...#....#...#...",
	"18: ..##.#.#....#####.#.#.#...##...##..##..",
	"19: .#..###.#..#.#.#######.#.#.#..#.#...#..",
	"20: .#....##....#####...#######....#.#..##.",
}
