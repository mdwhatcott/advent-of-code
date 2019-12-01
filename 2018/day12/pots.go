package day12

import (
	"bytes"
	"strings"
)

type Rule struct {
	pots   string
	result string
}

func ParseRule(raw string) Rule {
	parts := strings.Split(raw, " => ")
	return Rule{
		pots:   parts[0],
		result: parts[1],
	}
}

func (this Rule) String() string {
	return this.result
}

func (this Rule) IsSatisfiedBy(pots string) bool {
	return pots == this.pots
}

type Row struct {
	min    int
	max    int
	state  string
	rules  []Rule
	buffer *bytes.Buffer
}

func NewRowOfPots(initial string, rules ...Rule) *Row {
	return &Row{
		min:    0,
		max:    len(initial),
		state:  initial,
		rules:  rules,
		buffer: new(bytes.Buffer),
	}
}

func (this *Row) Scan() {
	this.buffer.WriteString(this.state)
	for p, pot := range this.state {
		var neighbors string

		if p == 0 {
			neighbors += ".."
			this.min -= 2
		} else if p == 1 {
			neighbors += "."
			this.min -= 1
		}

		neighbors += string(pot)

		if p == len(this.state)-1 {
			neighbors += "."
			this.max += 1
		} else if p == len(this.state)-2 {
			neighbors += ".."
			this.max += 2
		}

		before := this.buffer.Len()

		for _, rule := range this.rules {
			if rule.IsSatisfiedBy(neighbors) {
				this.buffer.WriteString(rule.String())
				break
			}
		}

		if this.buffer.Len() == before {
			this.buffer.WriteRune(pot)
		}
	}
	/*
		for pot in state:
			if pot is near the edge:
				fill in adjacent (empty) neighbor pots
					remember to adjust min/max accordingly...
			for rule in rules: // optimization: match the neighbors to rules by traversing a tree...
				if rule.IsSatisfiedBy(neighbors):
					write rule.String() to buffer
					break
			else:
				write current pot to buffer
	*/
}

func (this *Row) Update() {
	this.state = this.buffer.String()
	this.buffer.Reset()
}

func (this *Row) Sum() (sum int) {
	for x, value := 0, this.min; value < this.max; x, value = x+1, value+1 {
		if this.state[x:x+1] == "#" {
			sum += value
		}
	}
	return sum
}

func (this *Row) Render() string {
	return ""
}
