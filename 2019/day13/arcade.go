package advent

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"advent/2019/intcode"
	"advent/lib/grid"
)

type GameConsole struct {
	program  []int
	render   bool
	clear    string
	x, y     int
	screen   map[grid.Point]int
	ball     grid.Point
	paddle   grid.Point
	joystick int
	score    int
}

func NewGameConsole(program []int) *GameConsole {
	this := &GameConsole{
		program: program,
		screen:  make(map[grid.Point]int),
		x:       -2,
		y:       -2,
	}
	return this
}

func (this *GameConsole) InsertQuarters(quarters int) {
	this.program[0] = quarters
}

func (this *GameConsole) EnableRendering() {
	this.clear = "clear"
	if runtime.GOOS == "windows" {
		this.clear = "cls"
	}
	this.render = true
}

func (this *GameConsole) in() int {
	return this.joystick
}

func (this *GameConsole) out(i int) {
	if this.x == -2 {
		this.x = i
	} else if this.y == -2 {
		this.y = i
	} else if this.x == -1 && this.y == 0 {
		this.trackScore(i)
	} else {
		this.updatePixel(i)
	}

	if !this.render {
		return
	}
	this.clearScreen()
	fmt.Println(Render(this.screen))
	fmt.Println("Score:", this.score)
	fmt.Print(strings.Repeat("\n", 15))
}

func (this *GameConsole) clearScreen() {
	_ = exec.Command(this.clear).Run()
}

func (this *GameConsole) updatePixel(i int) {
	point := grid.NewPoint(float64(this.x), float64(this.y))
	this.screen[point] = i
	this.trackPaddle(i, point)
	this.trackBall(i, point)
	this.adjustJoystick()
	this.resetPixelBuffer()
}

func (this *GameConsole) trackBall(id int, location grid.Point) {
	if id == Ball {
		this.ball = location
	}
}
func (this *GameConsole) trackPaddle(id int, location grid.Point) {
	if id == Paddle {
		this.paddle = location
	}
}
func (this *GameConsole) adjustJoystick() {
	if this.ball.X() < this.paddle.X() {
		this.joystick = -1
	} else if this.ball.X() > this.paddle.X() {
		this.joystick = 1
	} else {
		this.joystick = 0
	}
}
func (this *GameConsole) trackScore(score int) {
	this.score = score
	this.resetPixelBuffer()
}

func (this *GameConsole) resetPixelBuffer() {
	this.x = -2
	this.y = -2
}

func (this *GameConsole) Play() (score int) {
	intcode.NewIntCodeInterpreter(this.program, this.in, this.out).RunProgram()
	return this.score
}
