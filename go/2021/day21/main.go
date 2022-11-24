package day21

import "advent/lib/util"

type Game struct {
	turn  bool
	die   int
	rolls int
	p1    *Player
	p2    *Player
}

func NewGame(p1, p2 int) *Game {
	return &Game{
		p1: &Player{at: p1},
		p2: &Player{at: p2},
	}
}

func (this *Game) Turn() bool {
	defer func() { this.turn = !this.turn }()
	player := this.p1
	if this.turn {
		player = this.p2
	}
	player.Turn(this.roll() + this.roll() + this.roll())
	return player.score < 1000
}

func (this *Game) roll() int {
	// In part 1, the die should reset after hitting 1000, but the game ends before that.
	this.die++
	this.rolls++
	return this.die
}

func (this *Game) Answer() int {
	return util.Min(this.p1.score, this.p2.score) * this.rolls
}

type Player struct {
	at    int
	score int
}

func (this *Player) Turn(rolls int) {
	this.at += rolls
	add := this.at % 10
	if add == 0 {
		add = 10
	}
	this.score += add
}
