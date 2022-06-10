package day14

import (
	"sort"
	"strings"
)

func Solve(start string, conversions []string, rounds int) int {
	ticker := NewTicker(start, conversions)
	for x := rounds; x > 0; x-- {
		ticker.Tick()
	}
	return maxMinusMin(ticker.pairs)
}

type Ticker struct {
	pairs      frequencies
	converters []*Converter
}

func NewTicker(start string, conversions []string) *Ticker {
	return &Ticker{
		pairs:      pairFrequencies(start),
		converters: parseConverters(conversions),
	}
}

func (this *Ticker) Tick() {
	for _, converter := range this.filterConverters() {
		converter.apply(this.pairs)
	}
}
func (this *Ticker) filterConverters() (filtered []*Converter) {
	for _, converter := range this.converters {
		if converter.countHits(this.pairs) > 0 {
			filtered = append(filtered, converter)
		}
	}
	return filtered
}
func maxMinusMin(pairs frequencies) int {
	frequencies := make(frequencies)
	for pair, count := range pairs {
		for _, char := range pair {
			frequencies[string(char)] += count
		}
	}
	var counts []int
	for _, count := range frequencies {
		counts = append(counts, (count+1)/2)
	}
	sort.Ints(counts)
	return counts[len(counts)-1] - counts[0]
}

type Converter struct {
	search string
	right  string
	left   string
	hits   int
}

func NewConverter(line string) *Converter {
	words := strings.Fields(line)
	search := words[0]
	left := string(search[0])
	middle := words[2]
	right := string(search[1])
	return &Converter{
		search: search,
		left:   left + middle,
		right:  middle + right,
	}
}
func (this *Converter) countHits(pairs frequencies) int {
	this.hits = pairs[this.search]
	return this.hits
}
func (this *Converter) apply(pairs frequencies) {
	pairs[this.search] -= this.hits
	pairs[this.left] += this.hits
	pairs[this.right] += this.hits
}

func pairFrequencies(start string) frequencies {
	pairs := make(frequencies)
	for x := 0; x < len(start)-1; x++ {
		pairs[start[x:x+2]]++
	}
	return pairs
}
func parseConverters(conversions []string) []*Converter {
	var allConverters []*Converter
	for _, conversion := range conversions {
		allConverters = append(allConverters, NewConverter(conversion))
	}
	return allConverters
}

type frequencies map[string]int

func (this frequencies) MaxValue() (max int) {
	for _, value := range this {
		if value > max {
			max = value
		}
	}
	return max
}
func (this frequencies) MinValue() (min int) {
	min = 0xFFFFFFFF
	for _, value := range this {
		if value < min {
			min = value
		}
	}
	return min
}
