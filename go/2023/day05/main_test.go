package day05

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
	"github.com/mdwhatcott/funcy"
	_ "github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/go-set/v2/set"
	_ "github.com/mdwhatcott/go-set/v2/set"
	_ "github.com/mdwhatcott/must/must"
	"github.com/mdwhatcott/must/strconvmust"
	"github.com/mdwhatcott/testing/should"
)

const TODO = -1

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
	should.Run(&Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Suite struct {
	*should.T
}

func (this *Suite) TestPart1A() {
	this.So(this.Solve(sampleInput, this.Part1SeedParser), should.Equal, 35)
}
func (this *Suite) TestPart1Full() {
	this.So(this.Solve(inputLines, this.Part1SeedParser), should.Equal, 261668924)
}
func (this *Suite) SkipTestPart2A() {
	this.So(this.Solve(sampleInput, this.Part2SeedParser), should.Equal, 46)
}
func (this *Suite) SkipTestPart2Full() {
	this.So(this.Solve(inputLines, this.Part2SeedParser), should.Equal, TODO)
}
func (this *Suite) Part1SeedParser(raw string) []int {
	return funcy.Map(strconvmust.Atoi, strings.Fields(raw))
}
func (this *Suite) Part2SeedParser(raw string) (results []int) {
	numbers := funcy.Map(strconvmust.Atoi, strings.Fields(raw))
	for x := 0; x < len(numbers); x += 2 {
		start, count := numbers[x], numbers[x+1]
		for y := start; y < start+count; y++ {
			results = append(results, y)
		}
	}
	return results
}
func (this *Suite) Solve(input []string, seedParser func(string) []int) (result int) {
	maps := make(map[string]func(int) int)
	seeds := set.Of[int]()
	title := "seed"
	var steps []string
	var ranges []int
	for _, line := range append(input, "") {
		if strings.HasPrefix(line, "seeds: ") && seeds.Len() == 0 {
			seeds.Add(seedParser(strings.TrimPrefix(line, "seeds: "))...)
		} else if line == "" && len(ranges) > 0 {
			maps[title] = rangeConverter(ranges)
		} else if strings.HasSuffix(line, " map:") {
			line = strings.TrimSuffix(line, " map:")
			words := strings.Split(line, "-")
			title = words[0]
			steps = append(steps, title)
			ranges = nil
		} else {
			ranges = append(ranges, funcy.Map(strconvmust.Atoi, strings.Fields(line))...)
		}
	}

	var locations []int
	for seed := range seeds {
		for _, step := range steps {
			converter := maps[step]
			if converter == nil {
				break
			}
			seed = converter(seed)
		}
		locations = append(locations, seed)
	}
	return funcy.Min(locations)
}

func rangeConverter(numbers []int) func(int) int {
	if len(numbers)%3 != 0 {
		panic(fmt.Sprintf("invalid number count: %d", numbers))
	}
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
