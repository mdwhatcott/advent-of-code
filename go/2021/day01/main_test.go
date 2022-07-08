package day01

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

func TestDay01Suite(t *testing.T) {
	in := util.InputInts("\n")
	assert.Error(t).So(Part1(in), should.Equal, 1688)
	assert.Error(t).So(Part2(in), should.Equal, 1728)
}

func Part1(depths []int) (result int) {
	for x := 1; x < len(depths); x++ {
		if depths[x] > depths[x-1] {
			result++
		}
	}
	return result
}

func Part2(depths []int) (result int) {
	for x := 3; x < len(depths); x++ {
		a, b, c := depths[x-3], depths[x-2], depths[x-1]
		A, B, C := depths[x-2], depths[x-1], depths[x-0]
		if A+B+C > a+b+c {
			result++
		}
	}
	return result
}
