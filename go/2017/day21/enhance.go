package day21

import (
	"fmt"
	"strings"
)

type EnhancementRules map[string]string

func RegisterEnhancementRules(lines ...string) EnhancementRules {
	rules := make(EnhancementRules)
	for _, line := range lines {
		fields := strings.Fields(line)
		for _, transformation := range Transformations(fields[0]) {
			rules[transformation] = fields[2]
		}
	}
	return rules
}

func (this EnhancementRules) Enhance(pattern string) string {
	if enhancement, found := this[pattern]; found {
		return enhancement
	} else {
		panic(fmt.Sprintln("Missing transformation for input pattern:", pattern))
	}
}
