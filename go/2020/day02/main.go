package advent

import (
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/parse"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func Part1() interface{} {
	lines := util.InputLines()
	valid := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		range_ := fields[0]
		ranges := strings.Split(range_, "-")
		lower := parse.Int(ranges[0])
		upper := parse.Int(ranges[1])
		letter := fields[1][0:1]
		value := fields[2]
		count := strings.Count(value, letter)
		if count < lower || count > upper {
			continue
		}
		valid++
	}
	return valid
}

func Part2() interface{} {
	lines := util.InputLines()
	valid := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		options := fields[0]
		options2 := strings.Split(options, "-")
		optionA := parse.Int(options2[0])
		optionB := parse.Int(options2[1])
		letter := fields[1][0:1]
		value := fields[2]
		if string(value[optionA-1]) == letter && string(value[optionB-1]) == letter {
			continue
		}
		if string(value[optionA-1]) == letter || string(value[optionB-1]) == letter {
			valid++
		}
	}
	return valid
}
