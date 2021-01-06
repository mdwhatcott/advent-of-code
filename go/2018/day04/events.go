package day04

import (
	"fmt"
	"strings"
	"time"

	"advent/lib/util"
)

type ClockedInEvent struct {
	Stamp   time.Time
	GuardID int
}

type AsleepEvent struct {
	Stamp time.Time
}

type AwakeEvent struct {
	Stamp time.Time
}

func parseEvent(line string) interface{} {
	line = strings.Replace(line, "#", "", -1)
	line = strings.TrimLeft(line, "[")
	fields := strings.Split(line, "]")
	stamp, err := time.Parse("2006-01-02 15:04", fields[0])
	if err != nil {
		panic(fmt.Sprint("time.Parse failure:", err))
	}
	switch fields = strings.Fields(fields[1]); fields[0] {
	case "Guard":
		return ClockedInEvent{Stamp: stamp, GuardID: util.ParseInt(fields[1])}
	case "falls":
		return AsleepEvent{Stamp: stamp}
	case "wakes":
		return AwakeEvent{Stamp: stamp}
	default:
		panic(fmt.Sprint("Parse failure:", line))
	}
}
