package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestCombinationsFixture(t *testing.T) {
	gunit.Run(new(CombinationsFixture), t)
}

type CombinationsFixture struct {
	*gunit.Fixture
}

func (this *CombinationsFixture) TestCombinations() {
	out := combinations([]int{0, 1, 2, 3}, 3)
	this.So(this.gather(out), should.Resemble, [][]int{
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
