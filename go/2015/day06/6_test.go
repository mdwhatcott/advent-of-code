package advent

import (
	"fmt"
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func Test6(t *testing.T) {
	if testing.Short() {
		//t.Skip("Long-running")
		return
	}
	grid := NewLightGrid()
	for _, instruction := range util.InputLines() {
		grid.Execute(instruction)
	}

	fmt.Println("How many lights on?", grid.HowManyLightsOn())
	fmt.Println("How bright are the lights?", grid.HowBrightAreAllLights())
}

var testInput = []string{
	"turn on 0,0 through 2,2",
}
