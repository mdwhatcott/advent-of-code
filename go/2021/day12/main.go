package day12

import (
	"strings"

	"advent/lib/util"
)

func Part1(lines util.Slice[string]) (result int) { return CountPaths(lines, true) }
func Part2(lines util.Slice[string]) (result int) { return CountPaths(lines, false) }

func CountPaths(lines util.Slice[string], twice bool) (result int) {
	paths := make(chan util.Slice[string])
	go func() {
		graph := parseInput(lines)
		graph.YieldPaths(paths, util.Slice[string]{"start"}, twice)
		close(paths)
	}()
	for range paths {
		result++
	}
	return result
}

type Graph map[string]util.Slice[string]

func parseInput(lines util.Slice[string]) (result Graph) {
	result = make(Graph)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		result[parts[0]] = append(result[parts[0]], parts[1])
		result[parts[1]] = append(result[parts[1]], parts[0])
	}
	return result
}

func (this Graph) YieldPaths(results chan util.Slice[string], path util.Slice[string], twice bool) {
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
