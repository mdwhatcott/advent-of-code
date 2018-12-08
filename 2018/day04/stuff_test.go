package day04

import (
	"strings"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuffFixture(t *testing.T) {
    gunit.Run(new(StuffFixture), t)
}

const toy = `[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up`

type StuffFixture struct {
    *gunit.Fixture
}

func (this *StuffFixture) TestToy() {
	projection := processGuardDutyLogs(strings.Split(toy, "\n"))
	sleepiest, minute := projection.Sleepiest()
	this.So(sleepiest, should.Equal, 10)
	this.So(minute, should.Equal, 24)
}

func (this *StuffFixture) TestToyPart2() {
	projection := processGuardDutyLogs(strings.Split(toy, "\n"))
	sleepiest, minute := projection.SleepiestOnSingleMinute()
	this.So(sleepiest, should.Equal, 99)
	this.So(minute, should.Equal, 45)
}
