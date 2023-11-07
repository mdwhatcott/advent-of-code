package starter

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/grid"
)

/*
for rounds := 0; ; rounds++
	sort characters in reading order
	for each character in reading order:
		if no targets:
			return rounds * sum(map(hit-points, living characters))

		if not adjacent to a target and there is any pathway to a target:
			move()

		if there is no pathway to any target:
			continue

		if adjacent to a target:
			attack()
*/

type World struct {
	height     int
	width      int
	cave       set.Set[grid.Point[int]]
	characters map[grid.Point[int]]*Character
}

func (this *World) String() string {
	var result strings.Builder
	result.WriteString("\n")
	for y := 0; y < this.height; y++ {
		var characters []*Character
		for x := 0; x < this.width; x++ {
			at := grid.NewPoint(x, y)
			if character, ok := this.characters[at]; ok {
				characters = append(characters, character)
				result.WriteRune(character.spec)
			} else if this.cave.Contains(at) {
				result.WriteRune('.')
			} else {
				result.WriteRune('#')
			}
		}
		if len(characters) == 0 {
			result.WriteString("\n")
			continue
		}
		result.WriteString("   ")
		for c, character := range characters {
			_, _ = fmt.Fprintf(&result, "%c(%d)", character.spec, 200) // TODO: HP
			if c < len(characters)-1 {
				result.WriteString(", ")
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}

func ParseWorld(lines []string) fmt.Stringer {
	characters := ParseCharacters(lines)
	index := make(map[grid.Point[int]]*Character)
	for _, character := range characters {
		index[character.loc] = character
	}
	world := &World{
		height:     len(lines),
		width:      len(lines[0]),
		cave:       ParseCaveMap(lines),
		characters: index, // TODO: associate
	}
	return world
}

func ParseCaveMap(lines []string) (result set.Set[grid.Point[int]]) {
	result = set.Make[grid.Point[int]](0)
	for y, line := range lines {
		for x, char := range line {
			if x >= len(lines[0]) {
				break
			}
			if char != '#' {
				result.Add(grid.NewPoint(x, y))
			}
		}
	}
	return result
}
func ParseCharacters(lines []string) (result []*Character) {
	for y, line := range lines {
		for x, char := range line {
			if char == 'G' || char == 'E' {
				result = append(result, NewCharacter(char, x, y))
			}
		}
	}
	return result
}

type Character struct {
	spec    rune
	loc     grid.Point[int]
	targets []*Character
}

func NewCharacter(spec rune, x, y int) *Character {
	return &Character{spec: spec, loc: grid.NewPoint(x, y)}
}

func AssociateCharacters(all ...*Character) {
	for c, c1 := range all {
		for _, c2 := range all[c+1:] {
			if c1.spec == c2.spec {
				continue
			}
			c1.targets = append(c1.targets, c2)
			c2.targets = append(c2.targets, c1)
		}
	}
}
