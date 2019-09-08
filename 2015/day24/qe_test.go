package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestQuantumEntanglementFixture(t *testing.T) {
	gunit.Run(new(QuantumEntanglementFixture), t)
}

type QuantumEntanglementFixture struct {
	*gunit.Fixture
}

func (this *QuantumEntanglementFixture) assertQE(expected int, weights []int) {
	this.So(QuantumEntanglement(weights...), should.Equal, expected)
}

func (this *QuantumEntanglementFixture) TestQE() {
	this.assertQE(0, nil)
	this.assertQE(0, []int{0})
	this.assertQE(1, []int{1})
	this.assertQE(1, []int{1, 1})
	this.assertQE(2, []int{1, 2})

	this.assertQE(99, []int{11, 9,})
	this.assertQE(90, []int{10, 9, 1})
	this.assertQE(160, []int{10, 8, 2})
	this.assertQE(210, []int{10, 7, 3})
	this.assertQE(200, []int{10, 5, 4, 1})
	this.assertQE(300, []int{10, 5, 3, 2})
	this.assertQE(240, []int{10, 4, 3, 2, 1})
	this.assertQE(216, []int{9, 8, 3})
	this.assertQE(252, []int{9, 7, 4})
	this.assertQE(360, []int{9, 5, 4, 2})
	this.assertQE(280, []int{8, 7, 5})
	this.assertQE(480, []int{8, 5, 4, 3})
	this.assertQE(420, []int{7, 5, 4, 3, 1})
}
