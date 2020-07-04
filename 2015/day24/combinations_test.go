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

func (this *CombinationsFixture) Setup() {
}

func (this *CombinationsFixture) SkipTestProduct() {
	this.So(this.gather(product([]interface{}{0, 1}, 3)), should.Resemble, [][]interface{}{
		{0, 0, 0},
		{0, 0, 1},
		{0, 1, 0},
		{0, 1, 1},
		{1, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 1},
	})
}

func (this *CombinationsFixture) SkipTestCombinations() {
	this.So(this.gather(combinations([]interface{}{0, 1, 2, 3}, 3)), should.Resemble, [][]int{
		{0, 1, 2},
		{0, 1, 3},
		{0, 2, 3},
		{1, 2, 3},
	})
}

func (this *CombinationsFixture) gather(stream chan []interface{}) (all [][]interface{}) {
	for each := range stream {
		all = append(all, each)
	}
	return all
}
