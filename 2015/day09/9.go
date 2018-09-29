package advent

import (
	"crypto/rand"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func CalculateShortestAndLongestDistances(lines []string) (shortest, longest int) {
	distances := make(map[string]map[string]int)
	uniqueCities := make(map[string]struct{})

	for _, line := range lines {
		fields := strings.Fields(line)
		source := fields[0]
		destination := fields[2]
		distance, _ := strconv.Atoi(fields[4])
		uniqueCities[source] = struct{}{}
		uniqueCities[destination] = struct{}{}
		if distances[source] == nil {
			distances[source] = make(map[string]int)
		}
		if distances[destination] == nil {
			distances[destination] = make(map[string]int)
		}
		distances[source][destination] = distance
		distances[destination][source] = distance
	}

	shortest = math.MaxInt32
	longest = 0

	for x := 0; x < 100000; x++ {
		journey := NewMap(citiesSlice(uniqueCities), distances)
		ant := NewAnt(journey)
		distance := ant.MakeRandomJourney()
		if distance < shortest {
			shortest = distance
		}
		if distance > longest {
			longest = distance
		}
	}

	return shortest, longest
}

func citiesSlice(unique map[string]struct{}) []string {
	cities := []string{}
	for city := range unique {
		cities = append(cities, city)
	}
	return cities
}

/**************************************************************************/

type Ant struct {
	journey  *Journey
	distance int
}

func NewAnt(journey *Journey) *Ant {
	return &Ant{journey: journey}
}

func (this *Ant) MakeRandomJourney() int {
	for !this.journey.Finished() {
		this.distance += this.journey.TravelNextLeg()
	}
	return this.distance
}

/**************************************************************************/

type Journey struct {
	location  string
	cities    []string
	visited   []string
	distances map[string]map[string]int
}

func NewMap(cities []string, distances map[string]map[string]int) *Journey {
	journey := &Journey{cities: cities, distances: distances}
	start, i := journey.chooseCity()
	journey.location = start
	journey.updateHistory(start, i)
	return journey
}

func (this *Journey) Finished() bool {
	return len(this.cities) == 0
}

func (this *Journey) TravelNextLeg() int {
	destination, index := this.chooseCity()
	distance := this.distances[this.location][destination]
	this.updateHistory(destination, index)
	return distance
}

func (this *Journey) chooseCity() (string, int) {
	random, _ := rand.Int(rand.Reader, big.NewInt(int64(len(this.cities))))
	index := int(random.Int64())
	return this.cities[index], index
}

func (this *Journey) updateHistory(city string, index int) {
	this.cities = append(this.cities[:index], this.cities[index+1:]...)
	this.visited = append(this.visited, city)
	this.location = city
}
