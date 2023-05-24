package day11

import (
	"github.com/mdwhatcott/advent-of-code-go-lib/grid"
)

func MaxPowerXYSize(table SummedAreaTable, Length, MinSize, MaxSize int) (maxPoint grid.Point, maxSize int) {
	var max int
	for size := MaxSize; size >= MinSize; size-- {
		for y := 0; y < Length-size+1; y++ {
			for x := 0; x < Length-size+1; x++ {
				upperLeft := grid.NewPoint(float64(x), float64(y))
				power := table.SummedArea(upperLeft, size)
				if power > max {
					max = power
					maxSize = size
					maxPoint = grid.NewPoint(float64(x), float64(y))
				}
			}
		}
	}
	return maxPoint, maxSize
}

func InitializePowerGrid(length int, serialNumber int) map[grid.Point]int {
	field := make(map[grid.Point]int, length*length)
	for x := float64(0); x < float64(length); x++ {
		for y := float64(0); y < float64(length); y++ {
			field[grid.NewPoint(x, y)] = power(int(x), int(y), serialNumber)
		}
	}
	return field
}

func power(x, y int, serial int) (result int) {
	rack := x + 10
	result = rack * y
	result += serial
	result *= rack
	result = hundreds(result)
	result -= 5
	return result
}

func hundreds(value int) (end int) {
	for value >= 1000 {
		value -= 1000
	}
	for value >= 100 {
		end++
		value -= 100
	}
	return end
}
