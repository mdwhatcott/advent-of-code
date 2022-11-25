package day21

import (
	"fmt"

	"advent/lib/util"
)

type DeterministicGame struct {
	die     *DeterministicDie
	players [2]Player
	turn    int
}

func NewDeterministicGame(p1, p2 int) *DeterministicGame {
	return &DeterministicGame{
		die: new(DeterministicDie),
		players: [2]Player{
			{at: p1},
			{at: p2},
		},
	}
}
func (this *DeterministicGame) Play() int64 {
	for this.Turn() {
	}
	return this.Answer()
}
func (this *DeterministicGame) Turn() bool {
	defer this.alternate()
	player := this.players[this.turn].Turn(this.die.RollN(3))
	this.players[this.turn] = player
	return player.score < 1000
}
func (this *DeterministicGame) alternate() {
	this.turn ^= 1
}
func (this *DeterministicGame) Answer() int64 {
	return this.die.rolls * util.Min(
		this.players[0].score,
		this.players[1].score,
	)
}

type DeterministicDie struct {
	die   int
	rolls int64
}

func (this *DeterministicDie) RollN(n int) (result int) {
	for ; n > 0; n-- {
		result += this.Roll()
	}
	return result
}

func (this *DeterministicDie) Roll() int {
	this.die++
	this.rolls++
	return this.die % 1000
}

//////////////////////////////////////////////////////////////

type Player struct {
	at    int
	score int64
}

func (this Player) Turn(rolls int) Player {
	at := this.at + rolls
	return Player{
		at:    at,
		score: this.score + int64((at-1)%10+1),
	}
}

//////////////////////////////////////////////////////////////

func PlayDirac(p1, p2 int) int64 {
	game := &DiracGame{cache: make(map[string]int64)}
	start := DiracStep{players: [2]Player{{at: p1}, {at: p2}}}
	return game.Play(start)
}

type DiracGame struct {
	cache map[string]int64
}

// Play is inspired by: https://github.com/1e9y/adventofcode/blob/main/2021/day21/day21.go
func (this *DiracGame) Play(step DiracStep) (sum int64) {
	key := step.String()
	if cached, ok := this.cache[key]; ok {
		return cached
	}
	defer func() { this.cache[key] = sum }()

	for dice, universes := range diracDiceRollDistribution {
		next := DiracStep{turn: step.turn, players: step.players}
		next.players[next.turn] = next.players[next.turn].Turn(dice)
		if next.players[next.turn].score >= 21 {
			if next.turn == 0 {
				sum += universes
			}
		} else {
			next.turn ^= 1
			sum += universes * this.Play(next)
		}
	}
	return sum
}

type DiracStep struct {
	turn    int
	players [2]Player
}

func (this DiracStep) String() string {
	return fmt.Sprintf("%d %d %d %d %d",
		this.turn,
		this.players[0].at,
		this.players[1].at,
		this.players[0].score,
		this.players[1].score,
	)
}

var diracDiceRollDistribution = map[int]int64{
	3: int64(len([][]int{
		{1, 1, 1},
	})),
	4: int64(len([][]int{
		{1, 1, 2},
		{1, 2, 1},
		{2, 1, 1},
	})),
	5: int64(len([][]int{
		{1, 1, 3},
		{1, 3, 1},
		{3, 1, 1},
		{2, 2, 1},
		{2, 1, 2},
		{1, 2, 2},
	})),
	6: int64(len([][]int{
		{1, 2, 3},
		{1, 3, 2},
		{3, 1, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 2, 1},
		{2, 2, 2},
	})),
	7: int64(len([][]int{
		{2, 2, 3},
		{2, 3, 2},
		{3, 2, 2},
		{1, 3, 3},
		{3, 1, 3},
		{3, 3, 1},
	})),
	8: int64(len([][]int{
		{3, 3, 2},
		{3, 2, 3},
		{2, 3, 3},
	})),
	9: int64(len([][]int{
		{3, 3, 3},
	})),
}
