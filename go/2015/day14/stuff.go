package day14

import (
	"strings"

	"advent/lib/parse"
)

func ParseReindeer(line string) Reindeer {
	fields := strings.Fields(line)
	return Reindeer{
		Name:     fields[0],
		Velocity: parse.Int(fields[3]),
		Sustain:  parse.Int(fields[6]),
		Rest:     parse.Int(fields[13]),
	}
}

type Reindeer struct {
	Name     string
	Velocity int
	Sustain  int
	Rest     int
}

type Simulator struct {
	distances map[Reindeer]int
	points    map[Reindeer]int
	ticks     int
}

func NewSimulator() *Simulator {
	return &Simulator{
		distances: make(map[Reindeer]int),
		points:    make(map[Reindeer]int),
	}
}

func (this *Simulator) Register(reindeer Reindeer) {
	this.distances[reindeer] = 0
}

func (this *Simulator) Tick(seconds int) {
	for x := 0; x < seconds; x++ {
		this.advanceReindeer()
		this.awardPoints()
		this.ticks++
	}
}
func (this *Simulator) advanceReindeer() {
	for reindeer := range this.distances {
		cycle := reindeer.Sustain + reindeer.Rest
		step := this.ticks % cycle
		if step < reindeer.Sustain {
			this.distances[reindeer] += reindeer.Velocity
		}
	}
}

func (this *Simulator) awardPoints() {
	for _, far := range this.farthestReindeer() {
		this.points[far]++
	}
}

func (this *Simulator) farthestReindeer() (farthest []Reindeer) {
	maxDistance := max(this.distances)

	for reindeer, distance := range this.distances {
		if distance == maxDistance {
			farthest = append(farthest, reindeer)
		}
	}

	return farthest
}

func max(all map[Reindeer]int) (max int) {
	for _, x := range all {
		if x > max {
			max = x
		}
	}
	return max
}
