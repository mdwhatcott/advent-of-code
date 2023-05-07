package advent

import (
	"github.com/mdwhatcott/advent-of-code/go/2019/intcode"
	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

const (
	TestMode  = 1
	BoostMode = 2
)

func Part1() interface{} {
	harness := intcode.NewHarness(util.InputInts(","), TestMode)
	harness.Run()
	return harness.Outputs()
}

func Part2() interface{} {
	harness := intcode.NewHarness(util.InputInts(","), BoostMode)
	harness.Run()
	return harness.Outputs()
}
