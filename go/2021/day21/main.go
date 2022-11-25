package day21

import "advent/lib/util"

type DeterministicGame struct {
	die     *DeterministicDie
	players [2]*Player
	turn    int
}

func NewDeterministicGame(p1, p2 int) *DeterministicGame {
	return &DeterministicGame{
		die: new(DeterministicDie),
		players: [2]*Player{
			{at: p1},
			{at: p2},
		},
	}
}
func (this *DeterministicGame) Turn() bool {
	defer func() { this.turn++ }()
	player := this.players[this.turn%2]
	player.Turn(0 +
		this.die.Roll() +
		this.die.Roll() +
		this.die.Roll(),
	)
	return player.score < 1000
}
func (this *DeterministicGame) Answer() int {
	return this.die.rolls * util.Min(
		this.players[0].score,
		this.players[1].score,
	)
}

type DeterministicDie struct {
	die   int
	rolls int
}

func (this *DeterministicDie) Roll() int {
	this.die++
	this.rolls++
	return this.die % 1000
}

type Player struct {
	at    int
	score int
}

func (this *Player) Turn(rolls int) {
	this.at += rolls
	this.score += (this.at-1)%10 + 1
}
