package advent

import (
	"bytes"
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func Part1() interface{} {
	buffer := new(bytes.Buffer)
	console := NewGameConsole(util.InputInts(","))
	console.EnableRendering(buffer)
	console.Play()
	final := strings.Split(buffer.String(), "######################################")
	return strings.Count(final[len(final)-1], "-")
}

func Part2() interface{} {
	console := NewGameConsole(util.InputInts(","))
	console.InsertQuarters(2)
	return console.Play()
}
