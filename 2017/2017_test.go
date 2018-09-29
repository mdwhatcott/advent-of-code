package advent

import (
	"testing"

	"advent/2017/day01"
	"advent/2017/day02"
	"advent/2017/day03"
	"advent/2017/day04"
	"advent/2017/day05"
	"advent/2017/day06"
	"advent/2017/day07"
	"advent/2017/day08"
	"advent/2017/day09"
	"advent/2017/day10"
	"advent/2017/day11"
	"advent/2017/day12"
	"advent/2017/day13"
	"advent/2017/day14"
	"advent/2017/day15"
	"advent/2017/day16"
	"advent/2017/day17"
	"advent/2017/day18"
	"advent/2017/day19"
	"advent/2017/day20"
	"advent/2017/day21"
	"advent/2017/day22"
	"advent/2017/day23"
	"advent/2017/day24"
	"advent/2017/day25"
)

func TestDay01(t *testing.T) {
	t.Parallel()
	AssertEqual(t, day01.Part1(), 1251)
	AssertEqual(t, day01.Part2(), 1244)
}
func TestDay02(t *testing.T) {
	t.Parallel()
	AssertEqual(t, day02.Part1(), 34925)
	AssertEqual(t, day02.Part2(), 221)
}
func TestDay03(t *testing.T) {
	t.Parallel()

	AssertEqual(t, day03.Part1(), 438)
	AssertEqual(t, day03.Part2(), 266330)
}
func TestDay04(t *testing.T) {
	t.Parallel()

	AssertEqual(t, day04.Part1(), 451)
	AssertEqual(t, day04.Part2(), 223)
}
func TestDay05(t *testing.T) {
	t.Parallel()

	AssertEqual(t, day05.Part1(), 351282)
	AssertEqual(t, day05.Part2(), 24568703)
}
func TestDay06(t *testing.T) {
	t.Parallel()

	AssertEqual(t, day06.Part1(), 5042)
	AssertEqual(t, day06.Part2(), 1086)
}
func TestDay07(t *testing.T) {
	t.Parallel()

	part1, part2 := day07.Answers()
	AssertEqual(t, part1, "dgoocsw")
	AssertEqual(t, part2, 1275)
}
func TestDay08(t *testing.T) {
	t.Parallel()

	part1, part2 := day08.Answers()
	AssertEqual(t, part1, 4902)
	AssertEqual(t, part2, 7037)
}
func TestDay09(t *testing.T) {
	t.Parallel()

	AssertEqual(t, day09.Part1(), 14212)
	AssertEqual(t, day09.Part2(), 6569)
}
func TestDay10(t *testing.T) {
	t.Parallel()

	AssertEqual(t, day10.Part1(), 6952)
	AssertEqual(t, day10.Part2(), "28e7c4360520718a5dc811d3942cf1fd")
}
func TestDay11(t *testing.T) {
	t.Parallel()

	part1, part2 := day11.Answers()
	AssertEqual(t, part1, 707)
	AssertEqual(t, part2, 1490)
}
func TestDay12(t *testing.T) {
	t.Parallel()

	part1, part2 := day12.Answers()
	AssertEqual(t, part1, 115)
	AssertEqual(t, part2, 221)
}
func TestDay13(t *testing.T) {
	t.Parallel()

	part1, part2 := day13.Answers()
	AssertEqual(t, part1, 788)
	AssertEqual(t, part2, 3905748)
}
func TestDay14(t *testing.T) {
	t.Parallel()

	AssertEqual(t, day14.Part1(), 8222)
	AssertEqual(t, day14.Part2(), 1086)
}
func TestDay15(t *testing.T) {
	t.Parallel()

	AssertEqual(t, day15.Part1(), 592)
	AssertEqual(t, day15.Part2(), 320)
}
func TestDay16(t *testing.T) {
	t.Parallel()

	AssertEqual(t, day16.Part1(), "giadhmkpcnbfjelo")
	AssertEqual(t, day16.Part2(), "njfgilbkcoemhpad")
}
func TestDay17(t *testing.T) {
	t.Skip("Not implemented")
	t.Parallel()

	AssertEqual(t, day17.Part1(), 0)
	AssertEqual(t, day17.Part2(), 0)
}
func TestDay18(t *testing.T) {
	t.Skip("Not implemented")
	t.Parallel()

	AssertEqual(t, day18.Part1(), 0)
	AssertEqual(t, day18.Part2(), 0)
}
func TestDay19(t *testing.T) {
	t.Skip("Not implemented")
	t.Parallel()

	AssertEqual(t, day19.Part1(), 0)
	AssertEqual(t, day19.Part2(), 0)
}
func TestDay20(t *testing.T) {
	t.Skip("Not implemented")
	t.Parallel()

	AssertEqual(t, day20.Part1(), 0)
	AssertEqual(t, day20.Part2(), 0)
}
func TestDay21(t *testing.T) {
	t.Skip("Not implemented")
	t.Parallel()

	AssertEqual(t, day21.Part1(), 0)
	AssertEqual(t, day21.Part2(), 0)
}
func TestDay22(t *testing.T) {
	t.Skip("Not implemented")
	t.Parallel()

	AssertEqual(t, day22.Part1(), 0)
	AssertEqual(t, day22.Part2(), 0)
}
func TestDay23(t *testing.T) {
	t.Skip("Not implemented")
	t.Parallel()

	AssertEqual(t, day23.Part1(), 0)
	AssertEqual(t, day23.Part2(), 0)
}
func TestDay24(t *testing.T) {
	t.Skip("Not implemented")
	t.Parallel()

	AssertEqual(t, day24.Part1(), 0)
	AssertEqual(t, day24.Part2(), 0)
}
func TestDay25(t *testing.T) {
	t.Skip("Not implemented")
	t.Parallel()

	AssertEqual(t, day25.Part1(), 0)
	AssertEqual(t, day25.Part2(), 0)
}
