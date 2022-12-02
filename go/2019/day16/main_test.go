package advent

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func (this *RepeatingPattern) Take(i int) (result []int) {
	for x := 0; x < i; x++ {
		result = append(result, this.Next())
	}
	return result
}
func TestRepeatingPattern(t *testing.T) {
	pattern := NewRepeater()
	should.So(t, pattern.Take(12), should.Equal, []int{1, 0, -1, 0, 1, 0, -1, 0, 1, 0, -1, 0})

	pattern.IncrementPeriod()
	should.So(t, pattern.Take(15), should.Equal, []int{0, 1, 1, 0, 0, -1, -1, 0, 0, 1, 1, 0, 0, -1, -1})

	pattern.IncrementPeriod()
	should.So(t, pattern.Take(11), should.Equal, []int{0, 0, 1, 1, 1, 0, 0, 0, -1, -1, -1})
}
func TestSplitAndJoinDigits(t *testing.T) {
	should.So(t, SplitDigits("1234567890"), should.Equal, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	should.So(t, JoinDigits([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}), should.Equal, "1234567890")
}
func TestPhase(t *testing.T) {
	should.So(t, Phase([]int{1, 2, 3, 4, 5, 6, 7, 8}), should.Equal, []int{4, 8, 2, 2, 6, 1, 5, 8})
	should.So(t, PhaseN(4, []int{1, 2, 3, 4, 5, 6, 7, 8}), should.Equal, []int{0, 1, 0, 2, 9, 4, 9, 8})
	should.So(t, PhaseN(100, SplitDigits("80871224585914546619083218645595"))[:8], should.Equal, SplitDigits("24176176"))
	should.So(t, PhaseN(100, SplitDigits("19617804207202209144916044189917"))[:8], should.Equal, SplitDigits("73745418"))
	should.So(t, PhaseN(100, SplitDigits("69317163492948606335995924319873"))[:8], should.Equal, SplitDigits("52432133"))
}
func TestSolutions(t *testing.T) {
	should.So(t, Part1(), should.Equal, "30369587")
	//should.So(t, Part2(), should.Equal, "")
}
