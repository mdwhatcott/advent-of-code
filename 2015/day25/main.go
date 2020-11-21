package main

import "fmt"

func main() {
	fmt.Println(calculateLocation(3029, 2947)) // 19980801
}

func calculateLocation(targetCol, targetRow int) int {
	n := 20151125

	for x := 0; x < totalCalculations(targetCol, targetRow); x++ {
		n = next(n)
	}

	return n
}

func next(current int) int {
	return current * 252533 % 33554393
}

func sumTo(col int) (result int) {
	for x := 1; x <= col; x++ {
		result += x
	}
	return result
}

func totalCalculations(col, row int) int {
	targetColumn := col + row - 1
	return sumTo(targetColumn) - (targetColumn - col) - 1
}
