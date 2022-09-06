package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestCombinationsFixture(t *testing.T) {
	should.Run(&CombinationsFixture{T: should.New(t)}, should.Options.UnitTests())
}

type CombinationsFixture struct {
	*should.T
}

func (this *CombinationsFixture) TestCombinations() {
	out := combinations([]int{0, 1, 2, 3}, 3)
	this.So(this.gather(out), should.Equal, [][]int{
		{0, 1, 2},
		{0, 1, 3},
		{0, 2, 3},
		{1, 2, 3},
	})
}

func (this *CombinationsFixture) gather(stream chan []int) (all [][]int) {
	for each := range stream {
		all = append(all, each)
	}
	return all
}
