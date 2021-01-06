package advent

import (
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

func (this *StuffFixture) Setup() {
}

func (this *StuffFixture) TestScan() {
	field := scanField(exampleMap1)
	this.So(field, should.Resemble, AsteroidField{
		{X: 1, Y: 0},
		{X: 4, Y: 0},
		{X: 0, Y: 2},
		{X: 1, Y: 2},
		{X: 2, Y: 2},
		{X: 3, Y: 2},
		{X: 4, Y: 2},
		{X: 4, Y: 3},
		{X: 3, Y: 4},
		{X: 4, Y: 4},
	})
}

func (this *StuffFixture) TestSlope() {
	a := NewAsteroid(0, 0)
	b := NewAsteroid(1, 1)
	c := NewAsteroid(2, 2)
	d := NewAsteroid(6, 7)

	this.So(Slope(a, b), should.Equal, 1)
	this.So(Slope(a, c), should.Equal, 1)
	this.So(Slope(c, b), should.Equal, 1)
	this.So(Slope(a, d), should.Equal, 6.0/7.0)
}

func (this *StuffFixture) TestCountVisible() {
	field := scanField(exampleMap1)

	this.So(CountVisible(field, field[0]), should.Equal, 7)
	this.So(CountVisible(field, field[1]), should.Equal, 7)
	this.So(CountVisible(field, field[2]), should.Equal, 6)
	this.So(CountVisible(field, field[3]), should.Equal, 7)
	this.So(CountVisible(field, field[4]), should.Equal, 7)
	this.So(CountVisible(field, field[5]), should.Equal, 7)
	this.So(CountVisible(field, field[6]), should.Equal, 5)
	this.So(CountVisible(field, field[7]), should.Equal, 7)
	this.So(CountVisible(field, field[8]), should.Equal, 8)
	this.So(CountVisible(field, field[9]), should.Equal, 7)
}

func (this *StuffFixture) TestBestPlace() {
	this.So(BestPlaceWithCount(scanField(exampleMap1)), should.Resemble, PlaceCount{Place: NewAsteroid(3, 4), Count: 8})
	this.So(BestPlaceWithCount(scanField(exampleMap2)), should.Resemble, PlaceCount{Place: NewAsteroid(5, 8), Count: 33})
	this.So(BestPlaceWithCount(scanField(exampleMap3)), should.Resemble, PlaceCount{Place: NewAsteroid(1, 2), Count: 35})
	this.So(BestPlaceWithCount(scanField(exampleMap4)), should.Resemble, PlaceCount{Place: NewAsteroid(6, 3), Count: 41})
	this.So(BestPlaceWithCount(scanField(exampleMap5)), should.Resemble, PlaceCount{Place: NewAsteroid(11, 13), Count: 210})
}

var exampleMap1 = strings.Split(strings.TrimSpace(""+
	".#..#\n"+
	".....\n"+
	"#####\n"+
	"....#\n"+
	"...##\n",
), "\n")

var exampleMap2 = strings.Split(strings.TrimSpace(""+
	"......#.#.\n"+
	"#..#.#....\n"+
	"..#######.\n"+
	".#.#.###..\n"+
	".#..#.....\n"+
	"..#....#.#\n"+
	"#..#....#.\n"+
	".##.#..###\n"+
	"##...#..#.\n"+
	".#....####\n",
), "\n")

var exampleMap3 = strings.Split(strings.TrimSpace(""+
	"#.#...#.#.\n"+
	".###....#.\n"+
	".#....#...\n"+
	"##.#.#.#.#\n"+
	"....#.#.#.\n"+
	".##..###.#\n"+
	"..#...##..\n"+
	"..##....##\n"+
	"......#...\n"+
	".####.###.\n",
), "\n")

var exampleMap4 = strings.Split(strings.TrimSpace(""+
	".#..#..###\n"+
	"####.###.#\n"+
	"....###.#.\n"+
	"..###.##.#\n"+
	"##.##.#.#.\n"+
	"....###..#\n"+
	"..#.#..#.#\n"+
	"#..#.#.###\n"+
	".##...##.#\n"+
	".....#.#..\n",
), "\n")

var exampleMap5 = strings.Split(strings.TrimSpace(""+
	".#..##.###...#######\n"+
	"##.############..##.\n"+
	".#.######.########.#\n"+
	".###.#######.####.#.\n"+
	"#####.##.#.##.###.##\n"+
	"..#####..#.#########\n"+
	"####################\n"+
	"#.####....###.#.#.##\n"+
	"##.#################\n"+
	"#####.##.###..####..\n"+
	"..######..##.#######\n"+
	"####.##.####...##..#\n"+
	".#####..#.######.###\n"+
	"##...#.##########...\n"+
	"#.##########.#######\n"+
	".####.#.###.###.#.##\n"+
	"....##.##.###..#####\n"+
	".#.#.###########.###\n"+
	"#.#.#.#####.####.###\n"+
	"###.##.####.##.#..##\n",
), "\n")
