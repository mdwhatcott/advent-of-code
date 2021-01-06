package lcd

import (
	"fmt"
	"strings"
	"time"

	"github.com/buger/goterm"
)

/**************************************************************************/

type Presenter interface {
	Present(lcd *LCD)
	Finalize()
}

/**************************************************************************/

type AnimatedPresenter struct{}

func NewAnimatedPresenter() *AnimatedPresenter {
	goterm.Clear()
	return &AnimatedPresenter{}
}

func (this *AnimatedPresenter) Present(lcd *LCD) {
	state := lcd.String()
	state = strings.Replace(state, ".", " ", -1)
	goterm.MoveCursor(1, 1)
	goterm.Println(state)
	goterm.Flush()
	time.Sleep(time.Millisecond * 75)
}

func (this *AnimatedPresenter) Finalize() {}

/**************************************************************************/

type StaticPresenter struct {
	lcd *LCD
}

func NewStaticPresenter() *StaticPresenter {
	return &StaticPresenter{}
}

func (this *StaticPresenter) Present(lcd *LCD) {
	this.lcd = lcd
}

func (this *StaticPresenter) Finalize() {
	fmt.Println(this.lcd)
	fmt.Println()

	count := 0
	for _, value := range this.lcd.grid {
		if value {
			count++
		}
	}
	fmt.Println("Pixels on:", count)
}
