package day12

import "strings"

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
	min   int
	max   int
	state string
	rules []Rule
}

func NewRowOfPots(initial string, rules ...Rule) *Row {
	return &Row{
		min:   0,
		max:   len(initial),
		state: initial,
		rules: rules,
	}
}

func (this *Row) Scan() {
	/*
		reset buffer
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

}

func (this *Row) Sum() (sum int) {
	for x, value := 0, this.min; value < this.max; x, value = x+1, value+1 {
		if this.state[x:x+1] == "#" {
			sum += value
		}
	}
	return sum
}

func (this *Row) String() string {
	return ""
}
