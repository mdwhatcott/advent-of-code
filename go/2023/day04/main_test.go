package day04

import (
	"math"
	"strings"
	"testing"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/must/strconvmust"
	"github.com/mdwhatcott/testing/should"
)

var (
	inputLines  = inputs.Read(2023, 4).Lines()
	sampleLines = []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}
)

func TestSuite(t *testing.T) {
	should.Run(&Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Suite struct{ *should.T }

func (this *Suite) TestSampleData() {
	part1, part2 := this.Solve(sampleLines)
	this.So(part1, should.Equal, 13)
	this.So(part2, should.Equal, 30)
}
func (this *Suite) TestRealData() {
	part1, part2 := this.Solve(inputLines)
	this.So(part1, should.Equal, 26426)
	this.So(part2, should.Equal, 6227972)
}
func (this *Suite) Solve(lines []string) (part1, part2 int) {
	counts := make(map[int]int)
	for l, line := range lines {
		winners, inHand, _ := strings.Cut(line[9:], "|")
		copies := numberSet(inHand).Intersection(numberSet(winners)).Len()
		part1 += int(math.Pow(2, float64(copies-1)))
		part2++
		card := l + 1
		counts[card]++
		count := counts[card]
		for x := 1; x <= copies; x++ {
			counts[card+x] += count
			part2 += count
		}
		delete(counts, card)
	}
	return part1, part2
}
func numberSet(raw string) set.Set[int] {
	return set.Of[int](funcy.Map(strconvmust.Atoi, strings.Fields(raw))...)
}
