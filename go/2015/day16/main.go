package main

import (
	"strings"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

func main() {
	candidates := gatherCandidates()
	var search = AuntSue{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	assert.So(nil, part1(candidates, search), should.Equal, 213)
	assert.So(nil, part2(candidates, search), should.Equal, 323)
}

func gatherCandidates() (candidates []AuntSue) {
	scanner := util.InputScanner()
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		line = strings.Replace(line, ":", "", -1)
		line = strings.Replace(line, ",", "", -1)
		words := strings.Fields(line)
		candidates = append(candidates, AuntSue{
			words[2]: util.ParseInt(words[3]),
			words[4]: util.ParseInt(words[5]),
			words[6]: util.ParseInt(words[7]),
		})
	}
	return candidates
}

func part1(candidates []AuntSue, search AuntSue) int {
	for i, sue := range candidates {
		if sue.Matches(search) {
			return i + 1
		}
	}
	panic("Matching Aunt Sue not found!")
}

func part2(candidates []AuntSue, search AuntSue) int {
	for i, sue := range candidates {
		if sue.MatchesRange(search) {
			return i + 1
		}
	}
	panic("Matching Aunt Sue not found!")
}
