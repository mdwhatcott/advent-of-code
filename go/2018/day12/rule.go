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
