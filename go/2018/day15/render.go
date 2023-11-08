package starter

import (
	"fmt"
	"strings"

	"github.com/mdwhatcott/grid"
)

func RenderWorld(world *World) string {
	var result strings.Builder
	result.WriteString("\n")
	for y := 0; y < world.height; y++ {
		var units []*Unit
		for x := 0; x < world.width; x++ {
			at := grid.NewPoint(x, y)
			if unit, ok := world.units[at]; ok {
				units = append(units, unit)
				result.WriteRune(unit.species)
			} else if world.cave.Contains(at) {
				result.WriteRune('.')
			} else {
				result.WriteRune('#')
			}
		}
		if len(units) == 0 {
			result.WriteString("\n")
			continue
		}
		result.WriteString("   ")
		for u, unit := range units {
			_, _ = fmt.Fprintf(&result, "%c(%d)", unit.species, unit.HP())
			if u < len(units)-1 {
				result.WriteString(", ")
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}
