package day12

import "strings"

type Graph map[string][]string

func parseInput(lines []string) (result Graph) {
	for _, line := range lines {
		parts := strings.Split(line, "-")
		result[parts[0]] = append(result[parts[0]], parts[1])
	}
	return result
}

func (this Graph) Paths(path []string, twice bool) {
}
