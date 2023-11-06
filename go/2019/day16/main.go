package advent

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/maths"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func Part1() string {
	return JoinDigits(PhaseN(100, SplitDigits(util.InputString())))[:8]
}

func Part2() interface{} {
	return JoinDigits(PhaseN(100, SplitDigits(strings.Repeat(util.InputString(), 10_000))))[:8]
}

func SplitDigits(s string) (output []int) {
	for _, c := range s {
		output = append(output, int(c-'0'))
	}
	return output
}
func JoinDigits(ints []int) string {
	var b strings.Builder
	for _, n := range ints {
		_, _ = fmt.Fprintf(&b, "%d", n)
	}
	return b.String()
}
func PhaseN(n int, inputs []int) []int {
	repeater := NewRepeater(inputs)
	for ; n > 0; n-- {
		fmt.Println("Phase:", n)
		repeater.Phase()
	}
	return repeater.inputs
}

var base = []int{0, 1, 0, -1}

type RepeatingPattern struct {
	inputs  []int
	outputs []int
	period  int
	repeats int
	base    int
}

func NewRepeater(inputs []int) *RepeatingPattern {
	this := &RepeatingPattern{
		inputs:  inputs,
		outputs: make([]int, len(inputs)),
	}
	this.Reset()
	return this
}

func (this *RepeatingPattern) Reset() {
	this.period = 1
	this.repeats = 1
	this.base = 0
}

func (this *RepeatingPattern) Phase() {
	this.Reset()
	for x := range this.inputs {
		sum := 0
		for _, n := range this.inputs {
			sum += n * this.Next()
		}
		this.IncrementPeriod()
		this.outputs[x] = maths.Abs(sum % 10)
	}
	this.inputs, this.outputs = this.outputs, this.inputs
}

func (this *RepeatingPattern) IncrementPeriod() {
	this.period++
	this.repeats = 1
	this.base = 0
}

func (this *RepeatingPattern) Next() int {
	if this.repeats%this.period == 0 {
		this.base++
	}
	this.repeats++
	return base[this.base%4]
}
