/*
This implementation, which solves both parts is interesting because
part 1 is solved much more slowly than part 2, which is not how things
usually work on the advent of code. This is a good example of how
sometimes making a really fast thing a bit slower can help make a
really slow thing much faster.

- Part 1 is now much slower (from 0s to 34s)...but
- Part 2 is now much faster (from 15m to 3s)!
- That's a net speed improvement of 14m. (I'll take it!)

I had wondered doing a backwards lookup was the way to speed up part 2
but I couldn't quite visualize it. This implementation was inspired by
my former mentor's lovely Clojure solution:

https://github.com/slagyr/advent-of-code/blob/master/src/aoc/2023/day5.clj
*/
package day05

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/must/strconvmust"
	"github.com/mdwhatcott/testing/should"
)

var (
	inputLines  = inputs.Read(2023, 5).Lines()
	sampleInput = strings.Split(`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`, "\n")
)

func TestSuite(t *testing.T) {
	should.Run(&Suite{T: should.New(t)}, should.Options.IntegrationTests())
}

type Suite struct {
	*should.T
}

func (this *Suite) TestPart1A() {
	this.So(this.SolveForwards(sampleInput), should.Equal, 35)
}
func (this *Suite) TestPart1Full() {
	this.So(this.SolveForwards(inputLines), should.Equal, 261668924)
}
func (this *Suite) TestPart2A() {
	this.So(this.SolveBackwards(sampleInput), should.Equal, 46)
}
func (this *Suite) TestPart2Full() {
	this.So(this.SolveBackwards(inputLines), should.Equal, 24261545)
}
func (this *Suite) SolveForwards(input []string) (result int) {
	seeds := this.parseSeeds(input[0])
	converters := this.parseConverters(input[2:], false)
	for s, seed := range seeds {
		location := this.convertAll(seed, converters...)
		if s == 0 || location < result {
			result = location
		}
	}
	return result
}
func (this *Suite) SolveBackwards(input []string) (result int) {
	seedRangePairs := this.parseSeeds(input[0])
	isSeed := this.seedChecker(seedRangePairs...)
	converters := this.parseConverters(input[2:], true)
	for location := 0; ; location++ {
		candidate := this.convertAll(location, converters...)
		if isSeed(candidate) {
			return location
		}
	}
}
func (this *Suite) convertAll(location int, converters ...func(int) int) int {
	for _, converter := range converters {
		location = converter(location)
	}
	return location
}
func (this *Suite) parseConverters(lines []string, backwards bool) (results []func(int) int) {
	if backwards {
		lines = funcy.Reverse(lines)
	}
	var numbers []int
	for _, line := range append(lines, "") {
		if strings.Contains(line, "map:") {
			continue
		}
		if line == "" && len(numbers) > 0 {
			results = append(results, this.rangeConverter(numbers))
			numbers = nil
			continue
		}
		if line == "" {
			continue
		}
		parsed := funcy.Map(strconvmust.Atoi, strings.Fields(line))
		if backwards {
			numbers = append(numbers, parsed[1], parsed[0], parsed[2])
		} else {
			numbers = append(numbers, parsed...)
		}
	}
	return results
}
func (this *Suite) seedChecker(rangePairs ...int) func(int) bool {
	return func(seed int) bool {
		for x := 0; x < len(rangePairs); x += 2 {
			lower, count := rangePairs[x], rangePairs[x+1]
			if lower <= seed && seed <= lower+count {
				return true
			}
		}
		return false
	}
}
func (this *Suite) rangeConverter(numbers []int) func(int) int {
	return func(i int) int {
		for x := 0; x < len(numbers); x += 3 {
			dest, source, length := numbers[x], numbers[x+1], numbers[x+2]
			if source <= i && i < source+length {
				return dest + (i - source)
			}
		}
		return i
	}
}
func (this *Suite) parseSeeds(line string) []int {
	return funcy.Map(strconvmust.Atoi, strings.Fields(line[len("seeds:"):]))
}
