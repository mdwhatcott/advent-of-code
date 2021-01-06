package main

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
	"testing"
)

func TestWinningElfFixture(t *testing.T) {
	gunit.Run(new(WinningElfFixture), t)
}

type WinningElfFixture struct {
	*gunit.Fixture
}

func (this *WinningElfFixture) Setup() {
}

func (this *WinningElfFixture) Test1() {
	this.So(WinningElf(5), should.Equal, 3)
}

func (this *WinningElfFixture) Test2() {
	this.So(WinningElf2(5), should.Equal, 2)
}
