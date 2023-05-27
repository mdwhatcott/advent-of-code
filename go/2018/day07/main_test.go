package day07_test

import (
	"testing"

	"github.com/mdwhatcott/advent-of-code/go/2018/day07"
	"github.com/mdwhatcott/testing/should"
)

func TestDay07(t *testing.T) {
	t.Parallel()

	should.So(t, day07.Part1(), should.Equal, "JDEKPFABTUHOQSXVYMLZCNIGRW")
	should.So(t, day07.Part2(), should.Equal, 1048)
}
