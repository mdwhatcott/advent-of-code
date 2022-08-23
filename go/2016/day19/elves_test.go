package main

import (
	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"

	"testing"
)

func TestWinningElfFixture(t *testing.T) {
	suite.Run(&WinningElfFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type WinningElfFixture struct {
	*suite.T
}

func (this *WinningElfFixture) Setup() {
}

func (this *WinningElfFixture) Test1() {
	this.So(WinningElf(5), should.Equal, 3)
}

func (this *WinningElfFixture) Test2() {
	this.So(WinningElf2(5), should.Equal, 2)
}
