package day20

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code-go-lib/intgrid"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func TestImageEnhancementSuite(t *testing.T) {
	should.Run(&ImageEnhancementSuite{T: should.New(t)}, should.Options.UnitTests())
}

type ImageEnhancementSuite struct{ *should.T }

func (this *ImageEnhancementSuite) TestParseGrid() {
	images := parseImage(sampleImage)
	this.So(len(images), should.Equal, 25)
	this.So(images[intgrid.NewPoint(1, 1)], should.BeFalse)
	this.So(images[intgrid.NewPoint(-1, -1)], should.BeFalse)

	this.So(images[intgrid.NewPoint(0, 0)], should.BeTrue)
	this.So(images[intgrid.NewPoint(0, 1)], should.BeTrue)
	this.So(images[intgrid.NewPoint(0, 2)], should.BeTrue)
	this.So(images[intgrid.NewPoint(1, 2)], should.BeTrue)
	this.So(images[intgrid.NewPoint(2, 3)], should.BeTrue)
	this.So(images[intgrid.NewPoint(2, 4)], should.BeTrue)
	this.So(images[intgrid.NewPoint(3, 0)], should.BeTrue)
	this.So(images[intgrid.NewPoint(3, 4)], should.BeTrue)
	this.So(images[intgrid.NewPoint(4, 2)], should.BeTrue)
	this.So(images[intgrid.NewPoint(4, 4)], should.BeTrue)
}
func (this *ImageEnhancementSuite) TestNiner() {
	image := parseImage(sampleImage)
	this.So(image.niner(intgrid.NewPoint(0, 4), false), should.Equal, 0)
	this.So(image.niner(intgrid.NewPoint(0, 0), false), should.Equal, 18)
	this.So(image.niner(intgrid.NewPoint(-1, -1), false), should.Equal, 1)
	this.So(image.niner(intgrid.NewPoint(-1, -1), true), should.Equal, 511)
}
func (this *ImageEnhancementSuite) TestCountLit() {
	this.So(parseImage(sampleImage).count(), should.Equal, 10)
}
func (this *ImageEnhancementSuite) TestEnhance() {
	algorithm := parseAlgorithm(sampleAlgorithm)
	image1 := parseImage(sampleImage)
	image2 := image1.enhance(algorithm, false)
	this.So(image2.count(), should.Equal, 24)
}

const sampleAlgorithm = "" +
	"..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#" +
	"....#..#..##..###..######.###...####..#..#####..##..#.####" +
	"#...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##" +
	"..#..#..#####.....#.#....###..#.##......#.....#..#..#..##." +
	".#...##.######.####.####.#.#...#.......#..#.#.#...####.##." +
	"#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#" +
	".#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####.." +
	"..##...##..#...#......#.#.......#.......##..####..#...#.#." +
	"#...##..#.#..###..#####........#..####......#..#"

var sampleImage = strings.Join([]string{
	"#..#.",
	"#....",
	"##..#",
	"..#..",
	"..###",
}, "\n")

func (this *ImageEnhancementSuite) TestPart1() {
	lines := util.InputLines()
	algorithm := parseAlgorithm(lines[0])
	image := parseImage(strings.Join(lines[2:], "\n"))
	image = image.enhance(algorithm, false)
	image = image.enhance(algorithm, true)
	this.So(image.count(), should.NOT.Equal, 5938)
	this.So(image.count(), should.Equal, 5259)
}
func (this *ImageEnhancementSuite) LongTestPart2() {
	lines := util.InputLines()
	algorithm := parseAlgorithm(lines[0])
	image := parseImage(strings.Join(lines[2:], "\n"))
	for x := 0; x < 25; x++ {
		image = image.enhance(algorithm, false)
		image = image.enhance(algorithm, true)
	}
	this.So(image.count(), should.Equal, 15287)
}
