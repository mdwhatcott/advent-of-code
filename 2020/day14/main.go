package advent

import (
	"bufio"
	"strconv"
	"strings"

	"advent/lib/util"
)

func Part1() interface{} {
	return part1(util.InputScanner().Scanner)
}

func part1(scanner *bufio.Scanner) interface{} {
	memory := make(map[uint64]uint64)
	var mask string

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if strings.HasPrefix(line, "mask") {
			mask = fields[2]
		} else {
			rawAddress := fields[0]
			rawAddress = strings.TrimPrefix(rawAddress, "mem[")
			rawAddress = strings.TrimSuffix(rawAddress, "]")
			address, _ := strconv.ParseUint(rawAddress, 10, 36)
			value := ApplyMask(fields[2], mask)
			memory[address] = value
		}
	}

	sum := uint64(0)
	for _, value := range memory {
		sum += value
	}

	return sum
}

func Part2() interface{} {
	return nil
}

func ApplyMask(rawInput, mask string) (output uint64) {
	OR, _ := strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 36)
	AND, _ := strconv.ParseUint(strings.ReplaceAll(mask, "X", "1"), 2, 36)
	input, _ := strconv.ParseUint(rawInput, 10, 36)
	output = input
	output = output | OR
	output = output & AND
	return output
}
