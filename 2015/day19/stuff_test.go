package main

import (
	"sort"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuff(t *testing.T) {
	gunit.Run(new(ReplacementFixture), t)
}

type ReplacementFixture struct {
	*gunit.Fixture
}

func (this *ReplacementFixture) TestReplacement_H_to_OH() {
	machine := NewMoleculeMachine()
	machine.RegisterReplacement("H => OH")
	results := machine.Calibrate("HOH")
	this.So(results, should.Resemble, []string{"OHOH", "HOOH"})
}

func (this *ReplacementFixture) TestReplacement_H_to_HO() {
	machine := NewMoleculeMachine()
	machine.RegisterReplacement("H => HO")
	results := machine.Calibrate("HOH")
	this.So(results, should.Resemble, []string{"HOOH", "HOHO"})
}

func (this *ReplacementFixture) TestReplacement_O_to_HO() {
	machine := NewMoleculeMachine()
	machine.RegisterReplacement("O => HH")
	results := machine.Calibrate("HOH")
	this.So(results, should.Resemble, []string{"HHHH"})
}

func (this *ReplacementFixture) TestHOH() {
	machine := NewMoleculeMachine()
	machine.RegisterReplacement("H => HO")
	machine.RegisterReplacement("H => OH")
	machine.RegisterReplacement("O => HH")

	results := machine.Calibrate("HOH")
	sort.Strings(results)

	this.So(results, should.Resemble, []string{
		"HHHH", // (via O => HH).
		"HOHO", // (via H => HO on the second H).
		"HOOH", // (via H => HO on the first H).
		"OHOH", // (via H => OH on the first H).
		//"HOOH", // (via H => OH on the second H). (NOT UNIQUE)
	})
}
