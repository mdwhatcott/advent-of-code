package lcd

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestLCDFixture(t *testing.T) {
	suite.Run(&LCDFixture{T: suite.New(t)}, suite.Options.UnitTests())
}

type LCDFixture struct {
	*suite.T

	lcd *LCD
}

func (this *LCDFixture) Setup() {
	this.lcd = NewLCD(7, 3)
}

func (this *LCDFixture) assertScreen(expected string) {
	result := this.lcd.String()
	result = strings.Replace(result, "\n", "|", -1)
	result = strings.Replace(result, " ", ".", -1)
	this.So(result, should.Equal, expected)
}

func (this *LCDFixture) TestAllOff() {
	this.assertScreen(".......|.......|.......")
}

func (this *LCDFixture) TestOneOn() {
	this.lcd.RectangleOn(1, 1)
	this.assertScreen("#......|.......|.......")
}

func (this *LCDFixture) TestAllOn() {
	this.lcd.RectangleOn(7, 3)
	this.assertScreen("#######|#######|#######")
}

func (this *LCDFixture) TestSequence() {
	this.lcd.RectangleOn(3, 2)
	this.assertScreen("###....|###....|.......")

	this.lcd.RotateColumn(1, 1)
	this.assertScreen("#.#....|###....|.#.....")

	this.lcd.RotateRow(0, 4)
	this.assertScreen("....#.#|###....|.#.....")

	this.lcd.RotateColumn(1, 1)
	this.assertScreen(".#..#.#|#.#....|.#.....")
}
