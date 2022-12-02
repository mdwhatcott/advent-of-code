package advent

import (
	"testing"

	day01 "advent/2019/day01"
	day02 "advent/2019/day02"
	day03 "advent/2019/day03"
	day04 "advent/2019/day04"
	day05 "advent/2019/day05"
	day06 "advent/2019/day06"
	day07 "advent/2019/day07"
	day08 "advent/2019/day08"
	day09 "advent/2019/day09"
	day10 "advent/2019/day10"
	day11 "advent/2019/day11"
	day12 "advent/2019/day12"
	day13 "advent/2019/day13"
	day14 "advent/2019/day14"
	day15 "advent/2019/day15"
	day16 "advent/2019/day16"
	day17 "advent/2019/day17"
	day18 "advent/2019/day18"
	day19 "advent/2019/day19"
	day20 "advent/2019/day20"
	day21 "advent/2019/day21"
	day22 "advent/2019/day22"
	day23 "advent/2019/day23"
	day24 "advent/2019/day24"
	day25 "advent/2019/day25"

	"github.com/mdwhatcott/testing/should"
)

func TestFixture2019(t *testing.T) {
	should.Run(&Fixture2019{T: should.New(t)}, should.Options.UnitTests())
}

type Fixture2019 struct {
	*should.T
}

func (this *Fixture2019) LongTestDay01() {
	this.So(day01.Part1(), should.Equal, 3563458)
	this.So(day01.Part2(), should.Equal, 5342292)
}
func (this *Fixture2019) LongTestDay02() {
	this.So(day02.Part1(), should.Equal, 3101878)
	this.So(day02.Part2(), should.Equal, 8444)
}
func (this *Fixture2019) LongTestDay03() {
	this.So(day03.Part1(), should.Equal, 2050)
	this.So(day03.Part2(), should.Equal, 21666)
}
func (this *Fixture2019) LongTestDay04() {
	this.So(day04.Part1(), should.Equal, 1929)
	this.So(day04.Part2(), should.Equal, 1306)
}
func (this *Fixture2019) LongTestDay05() {
	this.So(day05.Part1(), should.Equal, 16348437)
	this.So(day05.Part2(), should.Equal, 6959377)
}
func (this *Fixture2019) LongTestDay06() {
	this.So(day06.Part1(), should.Equal, 295936)
	this.So(day06.Part2(), should.Equal, 457)
}
func (this *Fixture2019) LongTestDay07() {
	this.So(day07.Part1(), should.Equal, 75228)
	this.So(day07.Part2(), should.Equal, 79846026)
}
func (this *Fixture2019) LongTestDay08() {
	this.So(day08.Part1(), should.Equal, 2684)
	this.So(day08.Part2(), should.Equal, "YGRYZ")
}
func (this *Fixture2019) LongTestDay09() {
	this.So(day09.Part1(), should.Equal, []int{3742852857})
	this.So(day09.Part2(), should.Equal, []int{73439})
}
func (this *Fixture2019) LongTestDay10() {
	this.So(day10.Part1(), should.Equal, 253)
	this.So(day10.Part2(), should.Equal, 815)
}
func (this *Fixture2019) LongTestDay11() {
	this.So(day11.Part1(), should.Equal, 1732)
	this.So(day11.Part2(), should.Equal, "ABCLFUHJ")
}
func (this *Fixture2019) LongTestDay12() {
	this.So(day12.Part1(), should.Equal, 12644)
	this.So(day12.Part2(), should.Equal, 290314621566528)
}
func (this *Fixture2019) LongTestDay13() {
	this.So(day13.Part1(), should.Equal, 226)
	this.So(day13.Part2(), should.Equal, 10800)
}
func (this *Fixture2019) LongTestDay14() {
	this.So(day14.Part1(), should.Equal, 198984)
	this.So(day14.Part2(), should.Equal, 7659732)
}
func (this *Fixture2019) LongTestDay15() {
	this.So(day15.Part1(), should.Equal, 308)
	this.So(day15.Part2(), should.Equal, 328)
}
func (this *Fixture2019) LongTestDay16() {
	this.So(day16.Part1(), should.Equal, "30369587")
	this.So(day16.Part2(), should.Equal, nil)
}
func (this *Fixture2019) LongTestDay17() {
	this.So(day17.Part1(), should.Equal, 7720)
	this.So(day17.Part2(), should.Equal, nil)
}
func (this *Fixture2019) LongTestDay18() {
	this.So(day18.Part1(), should.Equal, nil)
	this.So(day18.Part2(), should.Equal, nil)
}
func (this *Fixture2019) LongTestDay19() {
	this.So(day19.Part1(), should.Equal, nil)
	this.So(day19.Part2(), should.Equal, nil)
}
func (this *Fixture2019) LongTestDay20() {
	this.So(day20.Part1(), should.Equal, nil)
	this.So(day20.Part2(), should.Equal, nil)
}
func (this *Fixture2019) LongTestDay21() {
	this.So(day21.Part1(), should.Equal, nil)
	this.So(day21.Part2(), should.Equal, nil)
}
func (this *Fixture2019) LongTestDay22() {
	this.So(day22.Part1(), should.Equal, nil)
	this.So(day22.Part2(), should.Equal, nil)
}
func (this *Fixture2019) LongTestDay23() {
	this.So(day23.Part1(), should.Equal, nil)
	this.So(day23.Part2(), should.Equal, nil)
}
func (this *Fixture2019) LongTestDay24() {
	this.So(day24.Part1(), should.Equal, nil)
	this.So(day24.Part2(), should.Equal, nil)
}
func (this *Fixture2019) LongTestDay25() {
	this.So(day25.Part1(), should.Equal, nil)
	this.So(day25.Part2(), should.Equal, nil)
}
