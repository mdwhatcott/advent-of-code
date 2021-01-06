package advent

import (
	"fmt"
	"strconv"
	"strings"
)

type LightGrid struct {
	grid   map[string]bool
	bright map[string]int
}

func NewLightGrid() *LightGrid {
	return &LightGrid{
		grid:   make(map[string]bool),
		bright: make(map[string]int),
	}
}

func (this *LightGrid) Execute(instruction string) {
	for _, i := range parse(instruction) {
		if light := i.String(); i.on {
			this.grid[light] = true
			this.bright[light]++
		} else if i.off {
			this.grid[light] = false
			this.bright[light]--
			if this.bright[light] < 0 {
				this.bright[light] = 0
			}
		} else if i.toggle {
			this.grid[light] = !this.grid[light]
			this.bright[light] += 2
		}
	}
}

func parse(instruction string) (instructions []GridInstruction) {
	instruction = strings.Replace(instruction, "turn", "", 1)
	instruction = strings.Replace(instruction, "through", "", 1)
	words := strings.Fields(instruction)
	action := words[0]
	fromXY := strings.Split(words[1], ",")
	fromX, _ := strconv.Atoi(fromXY[0])
	fromY, _ := strconv.Atoi(fromXY[1])
	toXY := strings.Split(words[2], ",")
	toX, _ := strconv.Atoi(toXY[0])
	toY, _ := strconv.Atoi(toXY[1])
	for x := fromX; x <= toX; x++ {
		for y := fromY; y <= toY; y++ {
			details := GridInstruction{
				x:      x,
				y:      y,
				on:     action == "on",
				off:    action == "off",
				toggle: action == "toggle",
			}
			instructions = append(instructions, details)
		}
	}
	return instructions
}

func (this *LightGrid) HowManyLightsOn() int {
	on := 0
	for _, state := range this.grid {
		if state {
			on++
		}
	}
	return on
}

func (this *LightGrid) HowBrightAreAllLights() int {
	brightness := 0
	for _, level := range this.bright {
		brightness += level
	}
	return brightness
}

type GridInstruction struct {
	x int
	y int

	on     bool
	off    bool
	toggle bool
}

func (this GridInstruction) String() string {
	return fmt.Sprintf("%d|%d", this.x, this.y)
}
