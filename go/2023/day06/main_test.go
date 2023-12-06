package day06

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/must/strconvmust"
	"github.com/mdwhatcott/testing/should"
)

var (
	inputLines  = inputs.Read(2023, 6).Lines()
	sampleLines = []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
)

func TestSuite(t *testing.T) {
	should.Run(&Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Suite struct{ *should.T }

func (this *Suite) TestPart1A() {
	this.So(this.Part1(sampleLines), should.Equal, 288)
}
func (this *Suite) TestPart1() {
	this.So(this.Part1(inputLines), should.Equal, 440000)
}
func (this *Suite) TestPart2A() {
	this.So(this.Part2(sampleLines), should.Equal, 71503)
}
func (this *Suite) TestPart2() {
	this.So(this.Part2(inputLines), should.Equal, 26187338)
}
func (this *Suite) Part1(lines []string) any {
	times := funcy.Map(strconvmust.Atoi, strings.Fields(strings.TrimSpace(lines[0][len("Time:"):])))
	distances := funcy.Map(strconvmust.Atoi, strings.Fields(strings.TrimSpace(lines[1][len("Distance:"):])))

	winningVelocities := make(map[int]int)
	for race := 0; race < len(times); race++ {
		for velocity := 1; velocity < times[race]; velocity++ {
			remaining := times[race] - velocity
			if remaining*velocity > distances[race] {
				winningVelocities[race]++
			}
		}
	}
	return funcy.Product(funcy.MapValues(winningVelocities))
}

func (this *Suite) Part2(lines []string) (result int) {
	time := strconvmust.Atoi(strings.ReplaceAll(strings.TrimSpace(lines[0][len("Time:"):]), " ", ""))
	distance := strconvmust.Atoi(strings.ReplaceAll(strings.TrimSpace(lines[1][len("Distance:"):]), " ", ""))

	for lower, velocity := 0, 1; ; velocity++ {
		travelled := (time - velocity) * velocity
		if lower == 0 && travelled > distance {
			lower = velocity
		} else if lower > 0 && travelled <= distance {
			return velocity - lower
		}
	}
}
