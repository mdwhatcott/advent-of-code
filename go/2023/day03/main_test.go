package day03

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/must/jsonmust"
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
func (this *Suite) TestPart2A() {
	this.So(this.Part2(sampleLines), should.Equal, 467835)
}
func (this *Suite) TestPart2() {
	this.So(this.Part2(inputLines), should.Equal, 80694070)
}
func (this *Suite) Part1(lines []string) (result int) {
	_, parts := this.ParseSymbolsAndParts(lines)
	for _, part := range parts {
		result += part.Value
	}
	return result
}

func (this *Suite) ParseSymbolsAndParts(lines []string) (map[Point]rune, []PartNumber) {
	symbols := make(map[Point]rune)
	var numbers []PartNumber
	for y, line := range lines {
		line = strings.ReplaceAll(line, ".", " ")
		number := PartNumber{}
		for x, char := range line {
			point := NewPoint(x, y)
			if char != ' ' && !unicode.IsDigit(char) {
				symbols[point] = char
			}
			if unicode.IsDigit(char) {
				digit := int(char - '0')
				number.Value *= 10
				number.Value += digit
				number.Points = append(number.Points, point)
				continue
			}
			if len(number.Points) > 0 {
				numbers = append(numbers, number)
				number = PartNumber{}
				continue
			}
		}
		if len(number.Points) > 0 {
			numbers = append(numbers, number)
			number = PartNumber{}
		}
	}
	partSet := set.Of[string]()
	for _, number := range numbers {
		for _, point := range number.Points {
			for _, neighbor := range point.Neighbors8() {
				if _, contains := symbols[neighbor]; contains {
					partSet.Add(string(jsonmust.Marshal(number)))
				}
			}
		}
	}
	return symbols, funcy.Map(Unmarshal[PartNumber], partSet.Slice())
}
func (this *Suite) Part2(lines []string) (result int) { // so slow
	symbols, parts := this.ParseSymbolsAndParts(lines)
	for at, symbol := range symbols {
		if symbol == '*' {
			var matches []PartNumber
			for _, part := range parts {
				if part.AdjacentTo(at) {
					matches = append(matches, part)
				}
			}
			if len(matches) == 2 {
				result += matches[0].Value * matches[1].Value
			}
		}
	}
	return result
}

type PartNumber struct {
	Value  int
	Points []Point
}

func (this PartNumber) AdjacentTo(p Point) bool {
	return set.Of(funcy.MapCat(Point.Neighbors8, this.Points)...).Contains(p)
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
func (this Point) Neighbors8() (neighbors []Point) {
	for _, offset := range Neighbors8() {
		neighbors = append(neighbors, this.Offset(offset.dx, offset.dy))
	}
	return neighbors
}
func Neighbors8() []Direction {
	return []Direction{
		NewDirection(1, 0),
		NewDirection(-1, 0),
		NewDirection(0, 1),
		NewDirection(0, -1),
		NewDirection(1, 1),
		NewDirection(-1, 1),
		NewDirection(1, -1),
		NewDirection(-1, -1),
	}
}

type Direction struct{ dx, dy int }

func NewDirection(dx, dy int) Direction {
	return Direction{dx: dx, dy: dy}
}

func Unmarshal[T any](data string) (result T) {
	jsonmust.Unmarshal([]byte(data), &result)
	return result
}
