package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestCombinationsFixture(t *testing.T) {
	suite.Run(&CombinationsFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type CombinationsFixture struct {
	*suite.T
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
