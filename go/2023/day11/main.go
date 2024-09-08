package day11

import (
	"iter"
	"maps"
	"strings"

	"github.com/mdwhatcott/advent-of-code-go-lib/intgrid"
	. "github.com/mdwhatcott/funcy/ranger"
	"github.com/mdwhatcott/go-set/v2/set"
)

func LocateGalaxies(grid string) set.Set[intgrid.Point] {
	result := set.Of[intgrid.Point]()
	for y, line := range strings.Split(grid, "\n") {
		for x, char := range line {
			if char == '#' {
				result.Add(intgrid.NewPoint(x, y))
			}
		}
	}
	return result
}
func ExpandUniverse(u set.Set[intgrid.Point]) set.Set[intgrid.Point] {
	emptyRows, emptyColumns := scanEmptyRowsAndColumns(u)
	for row := range emptyRows {
		for g := range u {
			if g.Y() > row {
				u.Remove(g)
				u.Add(g.Move(intgrid.Up))
			}
		}
	}
}

func scanEmptyRowsAndColumns(universe set.Set[intgrid.Point]) (emptyRows, emptyCols iter.Seq[int]) {
	rows := set.FromSeq(Map(intgrid.Point.Y, maps.Keys(universe)))
	cols := set.FromSeq(Map(intgrid.Point.X, maps.Keys(universe)))
	emptyRows = Remove(rows.Contains, Range(Min(maps.Keys(rows)), Max(maps.Keys(rows))))
	emptyCols = Remove(cols.Contains, Range(Min(maps.Keys(cols)), Max(maps.Keys(cols))))
	return emptyRows, emptyCols
}
