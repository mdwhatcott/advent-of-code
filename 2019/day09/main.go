package advent

import (
	"advent/2019/intcode"
	"advent/lib/util"
)

func Part1() interface{} {
	harness := intcode.NewHarness(util.InputInts(","), 1)
	harness.Run()
	return harness.Outputs()
}

func Part2() interface{} {
	return nil
}
