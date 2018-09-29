package main

import (
	"fmt"
	"log"

	"advent/lib/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	fmt.Println(assert.So(part1(), should.Equal, 2696))
	fmt.Println(assert.So(part2(), should.Equal, 1084))
}

func part1() int { return max(runSimulation().distances) }
func part2() int { return max(runSimulation().points) }

func runSimulation() *Simulator {
	simulator := NewSimulator()
	scanner := util.InputScanner()
	for scanner.Scan() {
		simulator.Register(ParseReindeer(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	simulator.Tick(timeTrialInSeconds)
	return simulator
}

const timeTrialInSeconds = 2503
