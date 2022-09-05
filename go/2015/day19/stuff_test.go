package main

import (
	"sort"
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestStuff(t *testing.T) {
	suite.Run(&ReplacementFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type ReplacementFixture struct {
	*suite.T
}

func (this *ReplacementFixture) TestReplacement_H_to_OH() {
	machine := NewMoleculeMachine()
	machine.RegisterReplacement("H => OH")
	results := machine.Calibrate("HOH")
	this.So(results, should.Equal, []string{"OHOH", "HOOH"})
}

func (this *ReplacementFixture) TestReplacement_H_to_HO() {
	machine := NewMoleculeMachine()
	machine.RegisterReplacement("H => HO")
	results := machine.Calibrate("HOH")
	this.So(results, should.Equal, []string{"HOOH", "HOHO"})
}

func (this *ReplacementFixture) TestReplacement_O_to_HO() {
	machine := NewMoleculeMachine()
	machine.RegisterReplacement("O => HH")
	results := machine.Calibrate("HOH")
	this.So(results, should.Equal, []string{"HHHH"})
}

func (this *ReplacementFixture) TestHOH() {
	machine := NewMoleculeMachine()
	machine.RegisterReplacement("H => HO")
	machine.RegisterReplacement("H => OH")
	machine.RegisterReplacement("O => HH")

	results := machine.Calibrate("HOH")
	sort.Strings(results)

	this.So(results, should.Equal, []string{
		"HHHH", // (via O => HH).
		"HOHO", // (via H => HO on the second H).
		"HOOH", // (via H => HO on the first H).
		"OHOH", // (via H => OH on the first H).
		//"HOOH", // (via H => OH on the second H). (NOT UNIQUE)
	})
}
