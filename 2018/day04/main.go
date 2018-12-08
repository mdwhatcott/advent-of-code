package day04

import "advent/lib/util"

func Part1() interface{} {
	projection := processGuardDutyLogs(util.InputLines())
	sleepiest, minute := projection.Sleepiest()
	return sleepiest * minute
}

func Part2() interface{} {
	projection := processGuardDutyLogs(util.InputLines())
	sleepiest, minute := projection.SleepiestOnSingleMinute()
	return sleepiest * minute
}

func processGuardDutyLogs(log []string) *SleepyGuardProjection {
	projection := NewSleepyGuardProjection()
	for _, line := range log {
		projection.Handle(parseEvent(line))
	}
	return projection
}

