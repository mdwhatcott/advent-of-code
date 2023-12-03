package day03

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/go-set/v2/set"
	_ "github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/must/jsonmust"
	_ "github.com/mdwhatcott/must/must"
	"github.com/mdwhatcott/testing/should"
)

var (
	inputLines  = inputs.Read(2023, 3).Lines()
	sampleLines = []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
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
	this.So(this.Part1(sampleLines), should.Equal, 4361)
}
func (this *Suite) TestPart1() {
	this.So(this.Part1(inputLines), should.Equal, 521601)
}
func (this *Suite) SkipTestPart2A() {
	this.So(this.Part2(sampleLines), should.Equal, -1)
}
func (this *Suite) SkipTestPart2() {
	this.So(this.Part2(inputLines), should.Equal, -1)
}
func (this *Suite) Part1(lines []string) (result int) {
	symbols := set.Of[Point]()
	var numbers []Number
	for y, line := range lines {
		line = strings.ReplaceAll(line, ".", " ")
		number := Number{}
		for x, char := range line {
			if char == ' ' && len(number.Points) > 0 {
				numbers = append(numbers, number)
				number = Number{}
				continue
			}
			point := NewPoint(x, y)
			if char != ' ' && !unicode.IsDigit(char) {
				symbols.Add(point)
				if len(number.Points) > 0 {
					numbers = append(numbers, number)
					number = Number{}
				}
				continue
			}
			if unicode.IsDigit(char) {
				digit := int(char - '0')
				number.Value *= 10
				number.Value += digit
				number.Points = append(number.Points, point)
			}
		}
		if len(number.Points) > 0 {
			numbers = append(numbers, number)
			number = Number{}
		}
	}
	parts := set.Of[string]()
	for _, number := range numbers {
		for _, point := range number.Points {
			for _, neighbor := range point.Neighbors8() {
				if symbols.Contains(neighbor) {
					parts.Add(string(jsonmust.Marshal(number)))
				}
			}
		}
	}
	for part := range parts {
		var p Number
		jsonmust.Unmarshal([]byte(part), &p)
		result += p.Value
	}
	return result
}
func (this *Suite) Part2(lines []string) any {
	return -1
}

type Number struct {
	Value  int
	Points []Point
}

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func (this Point) Offset(x, y int) Point {
	return NewPoint(this.X+x, this.Y+y)
}
func (this Point) String() string {
	return fmt.Sprintf("(%v, %v)", this.X, this.Y)
}
func (this Point) Neighbors4() (neighbors []Point) {
	for _, offset := range Neighbors4() {
		neighbors = append(neighbors, this.Offset(offset.dx, offset.dy))
	}
	return neighbors
}
func (this Point) Neighbors8() (neighbors []Point) {
	for _, offset := range Neighbors8() {
		neighbors = append(neighbors, this.Offset(offset.dx, offset.dy))
	}
	return neighbors
}
func Neighbors4() []Direction {
	return []Direction{
		NewDirection(1, 0),
		NewDirection(-1, 0),
		NewDirection(0, 1),
		NewDirection(0, -1),
	}
}
func Neighbors8() []Direction {
	return append(Neighbors4(),
		NewDirection(1, 1),
		NewDirection(-1, 1),
		NewDirection(1, -1),
		NewDirection(-1, -1),
	)
}

type Direction struct{ dx, dy int }

func NewDirection(dx, dy int) Direction {
	return Direction{dx: dx, dy: dy}
}
