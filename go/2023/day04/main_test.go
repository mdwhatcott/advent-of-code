package day04

import (
	"math"
	"strings"
	"testing"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/go-set/v2/set"
	_ "github.com/mdwhatcott/go-set/v2/set"
	_ "github.com/mdwhatcott/must/must"
	"github.com/mdwhatcott/must/strconvmust"
	"github.com/mdwhatcott/testing/should"
)

const TODO = -1

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

type Suite struct {
	*should.T
}

func (this *Suite) Setup() {
}

func (this *Suite) TestPart1A() {
	this.So(this.Part1(sampleLines), should.Equal, 13)
}
func (this *Suite) TestPart1() {
	this.So(this.Part1(inputLines), should.Equal, 26426)
}
func (this *Suite) TestPart2A() {
	this.So(this.Part2(sampleLines), should.Equal, 30)
}
func (this *Suite) TestPart2() {
	this.So(this.Part2(inputLines), should.Equal, TODO)
}
func (this *Suite) Part1(lines []string) (result int) {
	return funcy.Sum(funcy.Map(this.Score, lines))
}
func (this *Suite) Part2(lines []string) (result int) {
	cards := []int{-1}
	var queue []int
	for l, line := range lines {
		line = line[9:]
		before, after, _ := strings.Cut(line, "|")
		winning := set.Of[int](funcy.Map(strconvmust.Atoi, strings.Fields(before))...)
		numbers := set.Of[int](funcy.Map(strconvmust.Atoi, strings.Fields(after))...)
		copies := numbers.Intersection(winning)
		cards = append(cards, copies.Len())
		queue = append(queue, l+1)
		this.Println("load:", l+1)
	}
	for len(queue) > 0 {
		result++
		p := queue[0]
		queue = queue[1:]
		this.Println("pop:", p, "queue:", queue, "result:", result)
		for x := 1; x <= cards[p]; x++ {
			next := p + x
			this.Println("load:", next)
			queue = append(queue, next)
		}
	}
	return result
}

func (this *Suite) Score(line string) (result int) {
	line = line[9:]
	before, after, _ := strings.Cut(line, "|")
	winning := set.Of[int](funcy.Map(strconvmust.Atoi, strings.Fields(before))...)
	numbers := set.Of[int](funcy.Map(strconvmust.Atoi, strings.Fields(after))...)
	winners := numbers.Intersection(winning)
	if len(winners) < 2 {
		return len(winners)
	}
	return int(math.Pow(2, float64(len(winners)-1)))
}
