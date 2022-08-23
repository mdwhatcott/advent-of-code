package main

import (
	"log"
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

func Test(t *testing.T) {
	assert.Error(t).So(part1(), should.Equal, 2696)
	assert.Error(t).So(part2(), should.Equal, 1084)
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
