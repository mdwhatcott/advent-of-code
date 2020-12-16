package advent

import (
	"fmt"
	"log"
	"strings"

	"advent/lib/util"
)

func Part1() interface{} {
	lines := util.InputLines()
	timestamp := util.ParseInt(lines[0])
	busses := util.ParseInts(strings.Split(lines[1], ","))
	min := 1_000_000
	minBus := -1
	for _, bus := range busses {
		if bus == 0 {
			continue
		}
		for x := 2; ; x++ {
			b := bus * x
			if b > timestamp {
				if b-timestamp < min {
					min = b - timestamp
					minBus = bus
				}
				break
			}
		}
	}
	return minBus * min
}

func Part2() interface{} {
	busses := loadBusses()
	timestamp := 100_000_000_000_000
	maxBus := util.Max(busses...)
	maxBusIndex := MaxIndex(busses...)
	for ; timestamp%maxBus != 0; timestamp++ {
		// Find first timestamp after 100_000_000_000_000 (the clue from instructions) that is divisible by the largest bus number (823 for my input).
	}
	for x := 0 ; ; x++ {
		if x % 1_000_000_000 == 0 {
			log.Println("PROGRESS:", x)
		}
		// Increment by max bus number (823) until we find a timestamp that matches.
		// Hopefully, starting above the provided clue and going 823 times faster than a totally naive brute force method will find the answer?
		count := checkCount(busses, maxBusIndex, timestamp)
		if count == len(busses) {
			return timestamp - maxBusIndex
		}
		if count >= len(busses)-5 {
			fmt.Println("Close, but not quite:", timestamp-maxBus)
		}
		timestamp += maxBus
	}
}

func loadBusses() []int {
	return util.ParseInts(strings.Split(util.InputLines()[1], ","))
}

func MaxIndex(busses ...int) int {
	max := 0
	maxIndex := 0
	for b, bus := range busses {
		if bus > max {
			max = bus
			maxIndex = b
		}
	}
	return maxIndex
}

func check(busses []int, bus int, timestamp int) bool {
	return checkCount(busses, bus, timestamp) == len(busses)
}

func checkCount(busses []int, bus int, timestamp int) (count int) {
	timestamp -= bus
	for _, b := range busses {
		count++
		if b == 0 {
			timestamp++
			continue
		}
		if timestamp%b != 0 {
			return count
		}
		timestamp++
	}
	return count
}
