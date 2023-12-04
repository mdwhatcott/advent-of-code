package starter

import (
	"fmt"
	"testing"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	_ "github.com/mdwhatcott/go-set/v2/set"
	_ "github.com/mdwhatcott/must/must"
	"github.com/mdwhatcott/testing/should"
)

const TODO = -1

var (
	inputLines  = inputs.Read(TODO, TODO).Lines()
	sampleLines = []string{
		fmt.Sprint(TODO),
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
	this.So(this.Part1(sampleLines), should.Equal, TODO)
}
func (this *Suite) TestPart1() {
	this.So(this.Part1(inputLines), should.Equal, TODO)
}
func (this *Suite) TestPart2A() {
	this.So(this.Part2(sampleLines), should.Equal, TODO)
}
func (this *Suite) TestPart2() {
	this.So(this.Part2(inputLines), should.Equal, TODO)
}
func (this *Suite) Part1(lines []string) any {
	return TODO
}

func (this *Suite) Part2(lines []string) any {
	return TODO
}
