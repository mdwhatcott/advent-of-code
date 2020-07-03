package advent

import (
	"advent/2019/intcode"
	"advent/lib/util"
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
