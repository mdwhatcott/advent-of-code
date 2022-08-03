package day12

import (
	"strings"

	"advent/lib/util"
)

func Part1(lines util.Strings) (result int) { return CountPaths(lines, true) }
func Part2(lines util.Strings) (result int) { return CountPaths(lines, false) }

func CountPaths(lines util.Strings, twice bool) (result int) {
	paths := make(chan util.Strings)
	go func() {
		graph := parseInput(lines)
		graph.YieldPaths(paths, util.Strings{"start"}, twice)
		close(paths)
	}()
	for range paths {
		result++
	}
	return result
}

type Graph map[string]util.Strings

func parseInput(lines util.Strings) (result Graph) {
	result = make(Graph)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		result[parts[0]] = append(result[parts[0]], parts[1])
		result[parts[1]] = append(result[parts[1]], parts[0])
	}
	return result
}

func (this Graph) YieldPaths(results chan util.Strings, path util.Strings, twice bool) {
	// Inspiration: https://github.com/fogleman/AdventOfCode2021/blob/main/12.py
	last := path[len(path)-1]
	if last == "end" {
		results <- path
		return
	}

	for _, p := range this[last] {
		small := p == strings.ToLower(p)
		if p == "start" || (twice && small && path.Contains(p)) {
			continue
		}
		longer := append(path, p)
		this.YieldPaths(results, longer, twice || (small && longer.Count(p) > 1))
	}
}
