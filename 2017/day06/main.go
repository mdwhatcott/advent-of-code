package day06

import (
	"fmt"
	"strings"

	"advent/lib/util"
)

func Part1() int {
	d := NewDebugger(util.ParseInts(strings.Fields(util.InputString())))
	return d.Debug() + 1
}

func Part2() int {
	d := NewDebugger(util.ParseInts(strings.Fields(util.InputString())))
	d.Debug() // perform part 1 calculations...
	d.states = make(map[string]struct{})
	d.states[fmt.Sprint(d.registers)] = struct{}{}
	return d.Debug()
}

type Debugger struct {
	registers []int
	cursor    int
	states    map[string]struct{}
}

func NewDebugger(registers []int) *Debugger {
	return &Debugger{
		registers: registers,
		states:    make(map[string]struct{}),
	}
}

func (this *Debugger) Debug() int {
	for {
		this.redistributeLargestBlock()
		if this.alreadySeenCurrentRegisterState() {
			break
		}
	}

	return len(this.states)
}
func (this *Debugger) redistributeLargestBlock() {
	largest := this.findLargestRegister()
	value := this.registers[largest]
	this.registers[largest] = 0
	this.cursor = largest
	for ; value > 0; value-- {
		this.cursor++
		if this.cursor >= len(this.registers) {
			this.cursor = 0
		}
		this.registers[this.cursor]++
	}
}
func (this *Debugger) findLargestRegister() int {
	max := -1
	index := -1
	for i, value := range this.registers {
		if value > max {
			index = i
			max = value
		}
	}
	return index
}
func (this *Debugger) alreadySeenCurrentRegisterState() bool {
	state := fmt.Sprint(this.registers)
	_, found := this.states[state]
	if !found {
		this.states[state] = struct{}{}
	}
	return found
}
