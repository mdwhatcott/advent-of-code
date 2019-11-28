package day11

import (
	"fmt"

	"advent/lib/util"
)

func Part1() interface{} {
	field := InitializePowerGrid(300, util.InputInt())
	table := NewSummedAreaTable(field)
	point, _ := MaxPowerXYSize(table, 300, 3, 3)
	return fmt.Sprintf("%v,%v", point.X(), point.Y())
}

func Part2() interface{} {
	field := InitializePowerGrid(300, util.InputInt())
	table := NewSummedAreaTable(field)
	point, size := MaxPowerXYSize(table, 300, 1, 300)
	return fmt.Sprintf("%v,%v,%v", point.X(), point.Y(), size)
}
