package day02

import (
	"slices"
	"strings"
	"testing"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/must/strconvmust"
	"github.com/mdwhatcott/testing/should"
)

var (
	inputLines  = inputs.Read(2023, 2).Lines()
	sampleLines = []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
)

func TestSuite(t *testing.T) {
	should.Run(&Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Suite struct{ *should.T }

func (this *Suite) TestPart1Sample() { this.So(this.Part1(sampleLines), should.Equal, 8) }
func (this *Suite) TestPart1()       { this.So(this.Part1(inputLines), should.Equal, 2449) }
func (this *Suite) TestPart2Sample() { this.So(this.Part2(sampleLines), should.Equal, 2286) }
func (this *Suite) TestPart2()       { this.So(this.Part2(inputLines), should.Equal, 63981) }

func (this *Suite) Part1(lines []string) (total int) {
	for _, line := range lines {
		if game := ParseGame(line); game.IsPossible() {
			total += game["id"][0]
		}
	}
	return total
}
func (this *Suite) Part2(lines []string) (result int) {
	for _, line := range lines {
		result += ParseGame(line).Power()
	}
	return result
}

type Game map[string][]int

func ParseGame(line string) (result Game) {
	result = make(map[string][]int)
	words := strings.Fields(cleaner.Replace(line))
	result["id"] = append(result["id"], strconvmust.Atoi(words[0]))
	for words = words[1:]; len(words) > 0; words = words[2:] {
		num, color := strconvmust.Atoi(words[0]), words[1]
		result[color] = append(result[color], num)
	}
	return result
}
func (this Game) IsPossible() bool {
	return this.power("red") <= 12 && this.power("green") <= 13 && this.power("blue") <= 14
}
func (this Game) Power() int {
	return this.power("red") * this.power("green") * this.power("blue")
}
func (this Game) power(color string) int {
	return slices.Max(this[color])
}

var cleaner = strings.NewReplacer("Game", "", ":", "", ";", "", ",", "")
