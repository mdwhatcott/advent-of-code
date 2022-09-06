package main

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestBotFixture(t *testing.T) {
	should.Run(&BotFixture{T: should.New(t)}, should.Options.UnitTests())
}

type BotFixture struct {
	*should.T

	bot  *Bot
	low  *Bot
	high *Bot
}

func (this *BotFixture) Setup() {
	this.low = &Bot{}
	this.high = &Bot{}
	this.bot = &Bot{
		low:  this.low,
		high: this.high,
	}

}

func (this *BotFixture) TestFirstValueStored() {
	this.bot.Receive(2)
	this.So(this.bot.value, should.Equal, 2)
	this.So(this.low.value, should.Equal, 0)
	this.So(this.high.value, should.Equal, 0)
}

func (this *BotFixture) TestSecondValueGivenAwayWithFirst() {
	this.bot.Receive(2)
	this.bot.Receive(1)
	this.So(this.low.value, should.Equal, 1)
	this.So(this.high.value, should.Equal, 2)
}

func (this *BotFixture) TestInstructionsFollowed() {
	start := Parse(input)
	start.Receive(start.start)
	this.So(start.outs[0].value, should.Equal, 5)
	this.So(start.outs[1].value, should.Equal, 2)
	this.So(start.outs[2].value, should.Equal, 3)
}

var input = strings.Split(`value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2`, "\n")
