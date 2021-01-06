package day04

import "fmt"

type SleepyGuardProjection struct {
	guards map[int]map[int]int // map[guard]map[minute]sleepCount
	guard  int
	minute int
}

func NewSleepyGuardProjection() *SleepyGuardProjection {
	return &SleepyGuardProjection{guards: make(map[int]map[int]int)}
}

func (this *SleepyGuardProjection) Handle(message interface{}) {
	switch event := message.(type) {
	case ClockedInEvent:
		this.initializeGuard(event.GuardID)
		this.initializeMinute(event)
	case AsleepEvent:
		this.countMinutesAsleep(event)
	case AwakeEvent:
		this.countMinutesAwake(event)
	default:
		panic(fmt.Sprint("Unprocessable entity:", event))
	}

}

func (this *SleepyGuardProjection) initializeGuard(guard int) {
	this.guard = guard
	if this.guards[guard] == nil {
		this.guards[guard] = make(map[int]int)
	}
}

func (this *SleepyGuardProjection) initializeMinute(event ClockedInEvent) {
	this.minute = 0
	if event.Stamp.Hour() == 0 {
		this.minute = event.Stamp.Minute()
	}
}

func (this *SleepyGuardProjection) countMinutesAsleep(event AsleepEvent) {
	for this.minute < event.Stamp.Minute() {
		this.minute++
	}
}

func (this *SleepyGuardProjection) countMinutesAwake(event AwakeEvent) {
	for profile := this.guards[this.guard]; this.minute < event.Stamp.Minute(); this.minute++ {
		profile[this.minute]++ // as we catch up to the awakening, mark each minute as asleep
	}
}

func (this *SleepyGuardProjection) Sleepiest() (sleepiest, minute int) {
	maxSleep := -1
	sleepiest = -1
	for guard, sleep := range this.guards {
		if sum := sumSleep(sleep); sum > maxSleep {
			sleepiest = guard
			maxSleep = sum
		}
	}
	return sleepiest, sleepiestMinute(this.guards[sleepiest])
}

func (this *SleepyGuardProjection) SleepiestOnSingleMinute() (sleepiest, minute int) {
	maxSleepCount := -1
	sleepiest = -1
	for guard, sleep := range this.guards {
		for MINUTE, sleepCount := range sleep {
			if sleepCount > maxSleepCount {
				minute = MINUTE
				sleepiest = guard
				maxSleepCount = sleepCount
			}
		}
	}
	return sleepiest, minute
}

func sumSleep(sleep map[int]int) (sum int) {
	for _, count := range sleep {
		sum += count
	}
	return sum
}

func sleepiestMinute(profile map[int]int) int {
	max := 0
	maxMinute := -1
	for minute, count := range profile {
		if count > max {
			max = count
			maxMinute = minute
		}
	}
	return maxMinute
}
