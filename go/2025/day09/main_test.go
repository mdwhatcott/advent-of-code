package day09

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/mdw-go/funcy/ranger/op"
	"github.com/mdw-go/testing/v2/better"
	"github.com/mdw-go/testing/v2/should"
	"github.com/mdw-go/testing/v2/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(&Suite{T: suite.New(t)}, suite.Options.UnitTests())
}

type Suite struct {
	*suite.T
}

func (this *Suite) TestArea() {
	this.So(RectangleArea(Point{2, 5}, Point{11, 1}), should.Equal, 50)
}
func (this *Suite) TestPart1Sample() {
	this.So(this.maxArea(this.gatherPoints("sample-input.txt")), should.Equal, 50)
}
func (this *Suite) TestPart1() {
	this.So(this.maxArea(this.gatherPoints("input.txt")), should.Equal, 4729332959)
}
func (this *Suite) gatherPoints(filename string) (results []Point) {
	file, err := os.Open(filename)
	this.So(err, better.BeNil)
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rawX, rawY, _ := strings.Cut(line, ",")
		var point Point
		point.X, err = strconv.Atoi(rawX)
		this.So(err, better.BeNil)
		point.Y, err = strconv.Atoi(rawY)
		this.So(err, better.BeNil)
		results = append(results, point)
	}
	this.So(scanner.Err(), better.BeNil)
	return results
}
func (this *Suite) maxArea(coords []Point) (maxArea int) {
	for p1 := range coords {
		for p2 := p1 + 1; p2 < len(coords); p2++ {
			area := RectangleArea(coords[p1], coords[p2])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

type Point struct {
	X, Y int
}

func RectangleArea(p1, p2 Point) int {
	return (op.Abs(p1.X-p2.X) + 1) * (op.Abs(p1.Y-p2.Y) + 1)
}
