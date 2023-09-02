package day21

import (
	"strings"

	"github.com/mdwhatcott/must/strconvmust"
)

type Monkey struct {
	Name     string
	Value    int
	OperandA string
	Operator string
	OperandB string
}

func parse(line string) (result Monkey) {
	fields := strings.Split(line, ":")
	result.Name = fields[0]
	fields = strings.Fields(fields[1])
	if len(fields) == 1 {
		result.Value = strconvmust.Atoi(fields[0])
		return result
	}
	result.OperandA = fields[0]
	result.Operator = fields[1]
	result.OperandB = fields[2]
	return result
}
func parse2(line string) (result Monkey) {
	result = parse(line)
	if result.Name == "root" {
		result.Operator = "-"
	}
	return result
}

func abs(result int) int {
	if result < 0 {
		return -result
	}
	return result
}
