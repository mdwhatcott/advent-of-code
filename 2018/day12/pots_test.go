package day12

import (
	"bufio"
	"strings"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuffFixture(t *testing.T) {
	gunit.Run(new(StuffFixture), t)
}

type StuffFixture struct {
	*gunit.Fixture
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
}

func (this *StuffFixture) TestFinalRowSum() {
	//                    -- ++++++++++++++++++++++++
	//                    210123456789012345678901234
	pots := NewRowOfPots("#....##....#####...#######....#.#..##")
	pots.min = -2
	pots.max = len(pots.state) + pots.min
	this.So(pots.Sum(), should.Equal, 325)
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
	scanner := bufio.NewScanner(strings.NewReader(ExampleGenerations))
	for scanner.Scan() {
		pots.Scan()
		pots.Update()
		continue
		this.So(pots.Render(), should.Equal, this.extractExpectedPotGrowth(scanner.Text()))
	}
}

func (this *StuffFixture) extractExpectedPotGrowth(text string) string {
	line := strings.TrimSpace(text)
	parts := strings.Fields(line)
	return strings.Trim(parts[1], ".")
}

const ExampleGenerations = ` 1: ...#...#....#.....#..#..#..#...........` /*
 2: ...##..##...##....#..#..#..##..........
 3: ..#.#...#..#.#....#..#..#...#..........
 4: ...#.#..#...#.#...#..#..##..##.........
 5: ....#...##...#.#..#..#...#...#.........
 6: ....##.#.#....#...#..##..##..##........
 7: ...#..###.#...##..#...#...#...#........
 8: ...#....##.#.#.#..##..##..##..##.......
 9: ...##..#..#####....#...#...#...#.......
10: ..#.#..#...#.##....##..##..##..##......
11: ...#...##...#.#...#.#...#...#...#......
12: ...##.#.#....#.#...#.#..##..##..##.....
13: ..#..###.#....#.#...#....#...#...#.....
14: ..#....##.#....#.#..##...##..##..##....
15: ..##..#..#.#....#....#..#.#...#...#....
16: .#.#..#...#.#...##...#...#.#..##..##...
17: ..#...##...#.#.#.#...##...#....#...#...
18: ..##.#.#....#####.#.#.#...##...##..##..
19: .#..###.#..#.#.#######.#.#.#..#.#...#..
20: .#....##....#####...#######....#.#..##.`
*/
