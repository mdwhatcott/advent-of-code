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

func (this Rule) Transform(pots string) string {
	if pots == this.pots {
		return this.result
	}
	return pots[2:3]
}
