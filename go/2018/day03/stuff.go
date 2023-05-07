package day03

import (
	"strings"

	"github.com/mdwhatcott/advent-of-code/go/lib/grid"
	"github.com/mdwhatcott/advent-of-code/go/lib/parse"
)

type PlotPoint struct {
	claims []int
}

func (this *PlotPoint) Claim(id int) {
	this.claims = append(this.claims, id)
}

type Fabric map[grid.Point]*PlotPoint

func (this Fabric) IsUndisputed(claim Claim) bool {
	for _, point := range claim {
		if len(this[point].claims) > 1 {
			return false
		}
	}
	return true
}

func plotClaimsOnFabric(claims []Claim) Fabric {
	fabric := make(Fabric)
	for _, claim := range claims {
		for id, point := range claim {
			plotPoint, found := fabric[point]
			if !found {
				plotPoint = new(PlotPoint)
				fabric[point] = plotPoint
			}
			plotPoint.Claim(id + 1)
		}
	}
	return fabric
}

func parseFields(line string) []int {
	line = strings.Replace(line, "#", " ", -1)
	line = strings.Replace(line, "@", " ", -1)
	line = strings.Replace(line, ",", " ", -1)
	line = strings.Replace(line, ":", " ", -1)
	line = strings.Replace(line, "x", " ", -1)
	return parse.Ints(strings.Fields(line))
}

type Claim []grid.Point

func parseClaims(cloth string) (claims []Claim) {
	for _, line := range strings.Split(cloth, "\n") {
		var claim Claim
		ints := parseFields(line)
		_, x, y, xmax, ymax := ints[0], ints[1], ints[2], ints[3], ints[4]
		xmax = x + xmax
		ymax = y + ymax
		for xx := x; xx < xmax; xx++ {
			for yy := y; yy < ymax; yy++ {
				point := grid.NewPoint(float64(xx), float64(yy))
				claim = append(claim, point)
			}
		}
		claims = append(claims, claim)
	}
	return claims
}
