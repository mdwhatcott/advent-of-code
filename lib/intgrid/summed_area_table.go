package intgrid

type SummedAreaTable [][]int

func NewSummedAreaTable(table [][]int) (summed SummedAreaTable) {
	for r := range table {
		summed = append(summed, nil)
		for c := range table[r] {
			summed[r] = append(summed[r], sumToUpperLeft(table, r, c))
		}
	}
	return summed
}

func sumToUpperLeft(table [][]int, row, column int) int {
	if row < 0 || column < 0 {
		return 0
	}
	return table[row][column] +
		sumToUpperLeft(table, row, column-1) +
		sumToUpperLeft(table, row-1, column) -
		sumToUpperLeft(table, row-1, column-1)
}

func (this SummedAreaTable) SumQuadrant(upperLeft, lowerRight Point) int {
	return this[upperLeft.Y()-1][upperLeft.X()-1] + this[lowerRight.X()][lowerRight.Y()] -
		this[upperLeft.Y()-1][lowerRight.X()] - this[lowerRight.Y()][upperLeft.X()-1]
}

