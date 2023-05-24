package advent

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/maths"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

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
	for ; n > 0; n-- {
		inputs = Phase(inputs)
	}
	return inputs
}
func Phase(inputs []int) (outputs []int) {
	repeater := NewRepeater()
	for range inputs {
		sum := 0
		for _, n := range inputs {
			sum += n * repeater.Next()
		}
		repeater.IncrementPeriod()
		outputs = append(outputs, maths.Abs(sum%10))
	}
	return outputs
}

var base = []int{0, 1, 0, -1}

type RepeatingPattern struct {
	period  int
	repeats int
	base    int
}

func NewRepeater() *RepeatingPattern {
	return &RepeatingPattern{
		period:  1,
		repeats: 1,
		base:    0,
	}
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

func Part1() string {
	return JoinDigits(PhaseN(100, SplitDigits(util.InputString())))[:8]
}

func Part2() interface{} {

	return nil
}
