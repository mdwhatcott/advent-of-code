package main

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/must/strconvmust"
)

func Test(t *testing.T) {
	graph := make(map[string][]string)
	rates := make(map[string]int)
	for _, line := range exampleInput { // inputs.Read(2022, 16).Lines() {
		words := split(line)
		valve := words[1]
		flowRate := strconvmust.Atoi(words[5])
		rates[valve] = flowRate
		neighbors := words[10:]
		for _, neighbor := range neighbors {
			graph[valve] = append(graph[valve], neighbor)
		}
	}

	t.Log("rates:", rates)
	t.Log("graph:", graph)

	cursor := State{
		Elapsed: 1,
		At:      "AA",
		Trail:   "AA",
		Opened:  "AA",
	}

	max := 0
	queue := make([]State, 0)
	queue = append(queue, cursor)
	for len(queue) > 0 {
		current, queue := queue[0], queue[1:]
		if current.Elapsed >= 2 {
			if tally := current.Tally(rates); tally > max {
				max = tally
				t.Log("setting max:", max)
			}
			t.Log("finished branch", current.Trail)
			continue
		}
		if !strings.Contains(current.Opened, current.At) && rates[current.At] > 0 {
			queue = append(queue, current.Open())
		}
		for _, neighbor := range graph[current.At] {
			queue = append(queue, current.Move(neighbor))
		}
	}
	t.Log("max:", max)
}

type State struct {
	Elapsed int    // how many minutes elapsed
	At      string // current valve
	Trail   string // concatenated valves
	Opened  string // concatenation of valves arranged according to the tick they were opened
}

func (this State) Tally(rates map[string]int) (result int) {
	for tick, valve := range strings.Split(this.Opened, "|") {
		result += rates[valve] * (30 - tick)
	}
	return result
}
func (this State) Move(to string) State {
	return State{
		Elapsed: this.Elapsed + 1,
		At:      to,
		Trail:   this.Trail + " " + to,
		Opened:  this.Opened,
	}
}
func (this State) Open() State {
	return State{
		Elapsed: this.Elapsed + 1,
		At:      this.At,
		Trail:   this.Trail,
		Opened:  this.Opened + " " + this.At,
	}
}
