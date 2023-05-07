package day13

import (
	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func Part1() string {
	m := NewMap(string(util.InputBytes()))
	go func() {
		for m.Tick() {
		}
	}()
	return (<-m.Signals).String()
}

func Part2() interface{} {
	m := NewMap(string(util.InputBytes()))
	go func() {
		for m.Tick() {
		}
	}()

	var last Point
	for p := range m.Signals {
		last = p
	}
	return last.String()
}
