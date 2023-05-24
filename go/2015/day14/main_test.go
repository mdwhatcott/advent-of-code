package day14

import (
	"log"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code-go-lib/util"
)

func Test(t *testing.T) {
	should.So(t, part1(), should.Equal, 2696)
	should.So(t, part2(), should.Equal, 1084)
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
