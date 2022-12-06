package day06

import (
	"testing"

	"github.com/mdwhatcott/go-collections/set"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

func TestDay06Part1(t *testing.T) {
	should.So(t, detectStartMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), should.Equal, 7)
	should.So(t, detectStartMarker("bvwbjplbgvbhsrlpgdmjqwftvncz"), should.Equal, 5)
	should.So(t, detectStartMarker("nppdvjthqldpwncqszvftbrmjlhg"), should.Equal, 6)
	should.So(t, detectStartMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), should.Equal, 10)
	should.So(t, detectStartMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), should.Equal, 11)
	should.So(t, detectStartMarker(util.InputString()), should.Equal, 1343)
}
func TestDay06Part2(t *testing.T) {
	should.So(t, detectMessageMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), should.Equal, 19)
	should.So(t, detectMessageMarker("bvwbjplbgvbhsrlpgdmjqwftvncz"), should.Equal, 23)
	should.So(t, detectMessageMarker("nppdvjthqldpwncqszvftbrmjlhg"), should.Equal, 23)
	should.So(t, detectMessageMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), should.Equal, 29)
	should.So(t, detectMessageMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), should.Equal, 26)
	should.So(t, detectMessageMarker(util.InputString()), should.Equal, 2193)
}
func detectMessageMarker(input string) int { return detectMarker(input, 14) }
func detectStartMarker(input string) int   { return detectMarker(input, 4) }
func detectMarker(input string, length int) int {
	for x := length; x < len(input); x++ {
		if set.From([]rune(input[x-length:x])...).Len() == length {
			return x
		}
	}
	panic("nope")
}
