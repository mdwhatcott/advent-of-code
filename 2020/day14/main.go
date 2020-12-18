package advent

import (
	"bufio"
	"fmt"
	"math"
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

func ApplyMask(rawInput, mask string) (output uint64) {
	OR, _ := strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 36)
	AND, _ := strconv.ParseUint(strings.ReplaceAll(mask, "X", "1"), 2, 36)
	input, _ := strconv.ParseUint(rawInput, 10, 36)
	output = input
	output = output | OR
	output = output & AND
	return output
}

func Part2() interface{} {
	return part2(util.InputScanner().Scanner)
}

func part2(scanner *bufio.Scanner) interface{} {
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
			binaryAddress := fmt.Sprintf("%036s", strconv.FormatUint(address, 2))
			maskedAddress := ApplyMask2(mask, binaryAddress)
			for _, floater := range PermuteFloatingBits(maskedAddress) {
				parsedAddress, _ := strconv.ParseUint(floater, 2, 36)
				parsedValue, _ := strconv.ParseUint(fields[2], 10, 36)
				memory[parsedAddress] = parsedValue
			}
		}
	}

	sum := uint64(0)
	for _, value := range memory {
		sum += value
	}

	return sum
}

func ApplyMask2(mask, value string) string {
	working := []byte(value)
	for x := 0; x < len(mask); x++ {
		if mask[x] == '1' {
			working[x] = '1'
		} else if mask[x] == 'X' {
			working[x] = 'X'
		}
	}
	return string(working)
}

func PermuteFloatingBits(value string) (result []string) {
	var floating []int
	for x := range value {
		if value[x] == 'X' {
			floating = append(floating, x)
		}
	}
	scope := float64(strings.Count(value, "X"))
	for x := 0.0; x < math.Pow(2, scope); x++ {
		binary := []byte(value)
		bits := strconv.FormatUint(uint64(x), 2)
		for len(bits) < int(scope) {
			bits = "0" + bits
		}
		for x, b := range bits {
			binary[floating[x]] = byte(b)
		}
		result = append(result, string(binary))
	}
	return result
}
