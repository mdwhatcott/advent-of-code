package advent

import (
	"bufio"
	"fmt"
	"log"
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

	departureIndex := map[int][]string{}

	for name, allowed := range input.DepartureFields {
		for x := 0; x < len(input.ValidTickets[0]); x++ {
			valid := true
			for _, ticket := range input.ValidTickets {
				value := ticket[x]
				if !allowed[value] {
					valid = false
					break
				}
			}
			if valid {
				departureIndex[x] = append(departureIndex[x], name)
				break
			}
		}
	}

	if len(departureIndex) > 6 {
		log.Panicln("too many fields:", len(departureIndex), departureIndex)
	}

	fmt.Println(departureIndex)
	departureProduct = 1
	for i := range departureIndex {
		departureProduct *= input.YourTicket[i]
	}
	return departureProduct
}

type Input struct {
	AllValidFieldValues map[int]bool

	DepartureFields map[string]map[int]bool

	YourTicket   []int
	AllTickets   [][]int
	ValidTickets [][]int
}

func ScanInput(scanner *bufio.Scanner) *Input {
	input := Input{
		AllValidFieldValues: make(map[int]bool),
		DepartureFields:     make(map[string]map[int]bool),
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
			line = halves[1]
			fields := strings.Fields(line)
			for _, field := range fields {
				bounds := strings.Split(field, "-")
				lower := util.ParseInt(bounds[0])
				upper := util.ParseInt(bounds[1])
				for x := lower; x <= upper; x++ {
					input.AllValidFieldValues[x] = true

					if strings.Contains(name, "departure") {
						allowed, ok := input.DepartureFields[name]
						if !ok {
							allowed = make(map[int]bool)
							input.DepartureFields[name] = allowed
						}
						allowed[x] = true
					}
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
