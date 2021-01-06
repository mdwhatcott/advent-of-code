package day12

import (
	"strings"
)

type Row struct {
	state map[int]bool
	rules []Rule
}

func NewRowOfPots(initial string, rules ...Rule) *Row {
	state := make(map[int]bool, len(initial))
	for i, c := range initial {
		state[i] = c == '#'
	}
	return &Row{
		state: state,
		rules: rules,
	}
}

func (this *Row) Scan() (result map[int]bool) {
	result = make(map[int]bool)

	min, max := this.MinMax()
	for pot := min - 5; pot < max+5; pot++ {
		for _, rule := range this.rules {
			if rule.IsSatisfiedBy(this.neighbors(pot)) {
				result[pot] = rule.result == "#"
			}
		}
	}

	return result
}
func (this *Row) neighbors(pot int) string {
	return "" +
		this.at(pot-2) +
		this.at(pot-1) +
		this.at(pot+0) +
		this.at(pot+1) +
		this.at(pot+2)
}
func (this *Row) at(i int) string {
	if this.state[i] {
		return "#"
	} else {
		return "."
	}
}

func (this *Row) Update(scan map[int]bool) {
	this.state = scan
}

func (this *Row) Sum() (sum int) {
	for pot, plant := range this.state {
		if plant {
			sum += pot
		}
	}
	return sum
}

func (this *Row) Render() string {
	min, max := this.MinMax()
	builder := new(strings.Builder)
	for x := min; x <= max; x++ {
		builder.WriteString(this.at(x))
	}
	return builder.String()
}

func (this *Row) MinMax() (min, max int) {
	for pot := range this.state {
		if pot < min {
			min = pot
		}
		if pot > max {
			max = pot
		}
	}
	return min, max
}

func (row *Row) Grow(generations int) interface{} {
	for x := 0; x < generations; x++ {
		row.Update(row.Scan())
	}
	return row.Sum()
}
