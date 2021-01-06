package day21

import "strings"

func SplitPatterns(grid string) (patterns []string) {
	if size := sizeOf(grid); size <= 3 {
		return append(patterns, grid)
	} else if size%2 == 0 {
		return splitIntoSize2Patterns(grid, size)
	} else if size%3 == 0 {
		return splitIntoSize3Patterns(grid, size)
	}
	return patterns
}

func splitIntoSize2Patterns(grid string, size int) (patterns []string) {
	for row := 0; row < size; row += 2 {
		for col := 0; col < size; col += 2 {
			index := col + (size+1)*row
			patterns = append(patterns, string([]byte{
				grid[index], grid[index+1], '/',
				grid[index+size+1], grid[index+size+1+1],
			}))
		}
	}
	return patterns
}

func splitIntoSize3Patterns(grid string, size int) (patterns []string) {
	for row := 0; row < size; row += 3 {
		for col := 0; col < size; col += 3 {
			index := col + (size+1)*row
			patterns = append(patterns, string([]byte{
				grid[index], grid[index+1], grid[index+2], '/',
				grid[index+size+1], grid[index+size+1+1], grid[index+size+2+1], '/',
				grid[index+size+size+2], grid[index+size+size+1+2], grid[index+size+size+2+2],
			}))
		}
	}
	return patterns
}

func ReassembleGrid(patterns ...string) string {
	grid := new(strings.Builder)
	patternSize := sizeOf(patterns[0])
	rowColCount := sqrt(len(patterns))

	for row := 0; row < rowColCount; row++ {
		for line := 0; line < patternSize; line++ {
			for col := 0; col < rowColCount; col++ {
				p := strings.Replace(patterns[col+rowColCount*row], "/", "", -1)
				for c := 0; c < patternSize; c++ {
					grid.WriteByte(p[c+patternSize*line])
				}
			}
			grid.WriteByte('/')
		}
	}
	return strings.Replace(grid.String(), "/", "\n", -1)
}
