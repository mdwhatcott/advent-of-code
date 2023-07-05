package main

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
)

func main() {
	fmt.Println("digraph G {")
	lines := inputs.Read(2022, 16).Lines()
	//lines := exampleInput
	names := map[string]string{}
	for _, line := range lines {
		words := split(line)
		origin := words[1]
		names[origin] = fmt.Sprintf("%s_%s", origin, words[5])
	}
	for _, line := range lines {
		words := split(line)
		origin := names[words[1]]
		targets := words[10:]
		for _, target := range targets {
			target = names[target]
			fmt.Printf("\t%s -> %s\n", origin, target)
		}
	}
	fmt.Println("}")
}
func split(line string) []string {
	line = strings.ReplaceAll(line, ",", " ")
	line = strings.ReplaceAll(line, ";", " ")
	line = strings.ReplaceAll(line, "=", " ")
	return strings.Fields(line)
}

var exampleInput = strings.Split(strings.TrimSpace(`
Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II
`), "\n")
