package main

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestWinningElfFixture(t *testing.T) {
	should.Run(&WinningElfFixture{T: should.New(t)}, should.Options.UnitTests())
}

type WinningElfFixture struct {
	*should.T
}

func (this *WinningElfFixture) Setup() {
}

func (this *WinningElfFixture) Test1() {
	this.So(WinningElf(5), should.Equal, 3)
}

func (this *WinningElfFixture) Test2() {
	this.So(WinningElf2(5), should.Equal, 2)
}
