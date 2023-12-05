package day05

import (
	"fmt"
	"strings"
	"testing"
	"time"

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
	now time.Time
}

func (this *Suite) Setup() {
	this.now = time.Now()
}
func (this *Suite) TestPart1A() {
	this.So(this.Solve(sampleInput, this.Part1SeedStreamer), should.Equal, 35)
}
func (this *Suite) TestPart1Full() {
	this.So(this.Solve(inputLines, this.Part1SeedStreamer), should.Equal, 261668924)
}
func (this *Suite) TestPart2A() {
	this.So(this.Solve(sampleInput, this.Part2SeedStreamer), should.Equal, 46)
}
func (this *Suite) LongTestPart2Full() {
	this.So(this.Solve(inputLines, this.Part2SeedStreamer), should.Equal, 24261545)
}
func (this *Suite) Part1SeedStreamer(raw string) chan int {
	result := make(chan int)
	go funcy.Load(result, funcy.Map(strconvmust.Atoi, strings.Fields(raw)))
	return result
}
func (this *Suite) Part2SeedStreamer(raw string) chan int {
	result := make(chan int)
	numbers := funcy.Map(strconvmust.Atoi, strings.Fields(raw))
	go func() {
		sent := 0
		for x := 0; x < len(numbers); x += 2 {
			start, count := numbers[x], numbers[x+1]
			this.Printf("seed range %d/%d starting at %d w/ count %d", x, len(numbers), start, count)
			for y := start; y < start+count; y++ {
				result <- y
				sent++
			}
		}
		close(result)
		this.Println("seeds:", sent, time.Since(this.now))
	}()
	return result
}
func (this *Suite) Solve(input []string, seedStreamer func(string) chan int) (result int) {
	maps := make(map[string]func(int) int)
	var seeds string
	title := "seed"
	var steps []string
	var ranges []int
	for _, line := range append(input, "") {
		if strings.HasPrefix(line, "seeds: ") {
			seeds = strings.TrimPrefix(line, "seeds: ")
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

	minLocation := 0xFFFFFFFF
	progress := 0
	for seed := range seedStreamer(seeds) {
		for _, step := range steps {
			progress++
			if progress%100_000_000 == 0 {
				this.Println("progress:", progress, time.Since(this.now))
			}
			converter := maps[step]
			if converter == nil {
				break
			}
			seed = converter(seed)
		}
		if seed < minLocation {
			minLocation = seed
			this.Println("new min:", seed, time.Since(this.now))
		}
	}
	this.Println("absolute min:", minLocation)
	return minLocation
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
