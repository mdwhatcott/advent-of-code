package advent

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func Test5_Examples(t *testing.T) {
	should.So(t, nice("aaa"), should.Equal, 1)
	should.So(t, nice("ugknbfddgicrmopn"), should.Equal, 1)

	should.So(t, nice("jchzalrnumimnmhp"), should.Equal, 0)
	should.So(t, nice("haegwjzuvuyypxyu"), should.Equal, 0)
	should.So(t, nice("dvszwmarrgswjxmb"), should.Equal, 0)

	n := 0
	for _, line := range strings.Split(util.InputString(), "\n") {
		n += nice(line)
	}
	should.So(t, n, should.Equal, 258)
	//t.Log("How many are nice?", n)
}

func Test5_Examples2(t *testing.T) {
	should.So(t, nice2("qjhvhtzxzqqjkmpb"), should.BeTrue)
	should.So(t, nice2("xxyxx"), should.BeTrue)

	should.So(t, nice2("xxxx"), should.BeFalse)
	should.So(t, nice2("uurcxstgmygtbstg"), should.BeFalse)
	should.So(t, nice2("ieodomkazucvgmuy"), should.BeFalse)

	n := 0
	for _, line := range strings.Split(util.InputString(), "\n") {
		if nice2(line) {
			n++
		}
	}
	should.So(t, n, should.Equal, 53)
	//t.Log("How many are nice2?", n)
}
