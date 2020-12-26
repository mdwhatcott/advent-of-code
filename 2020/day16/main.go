package advent

import (
	"bufio"
	"strings"

	"advent/lib/util"
)

func Part1() (errorRate int) {
	input := ScanInput(util.InputScanner().Scanner)
	for _, ticket := range input.AllTickets {
		for _, value := range ticket {
			if !input.AllValidFieldValues[value] {
				errorRate += value
			}
		}
	}
	return errorRate
}

func Part2() (departureProduct int) {
	input := ScanInput(util.InputScanner().Scanner)
	for _, ticket := range input.AllTickets {
		valid := true
		for _, value := range ticket {
			if !input.AllValidFieldValues[value] {
				valid = false
				break
			}
		}
		if valid {
			input.ValidTickets = append(input.ValidTickets, ticket)
		}
	}

	return departureProduct
}

type Input struct {
	AllValidFieldValues map[int]bool

	YourTicket   []int
	AllTickets   [][]int
	ValidTickets [][]int
}

func ScanInput(scanner *bufio.Scanner) *Input {
	input := Input{
		AllValidFieldValues: make(map[int]bool),
	}
	section := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else if line == "your ticket:" {
			section++
			continue
		} else if line == "nearby tickets:" {
			section++
			continue
		}

		switch section {

		case 0:
			line := strings.Replace(line, " or ", " ", 1)
			halves := strings.Split(line, ": ")
			name := halves[0]
			_ = name
			line = halves[1]
			fields := strings.Fields(line)
			for _, field := range fields {
				bounds := strings.Split(field, "-")
				lower := util.ParseInt(bounds[0])
				upper := util.ParseInt(bounds[1])
				for x := lower; x <= upper; x++ {
					input.AllValidFieldValues[x] = true
				}
			}

		case 1:
			input.YourTicket = util.ParseInts(strings.Split(line, ","))
		case 2:
			input.AllTickets = append(input.AllTickets, util.ParseInts(strings.Split(line, ",")))
		}
	}
	return &input
}
