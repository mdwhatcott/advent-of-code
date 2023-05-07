package advent

import (
	"github.com/mdwhatcott/advent-of-code/go/lib/maths"
	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func Part1() int {
	numbers := util.InputInts("\n")
	for x := 26; x < len(numbers); x++ {
		if !canSum(numbers[x-26 : x+1]...) {
			return numbers[x]
		}
	}
	panic("boink")
}

func Part2() interface{} {
	numbers := util.InputInts("\n")
	search := Part1()
	for x := 0; x < len(numbers); x++ {
		sum := 0
		for y := x + 1; sum < search; y++ {
			sum += numbers[y]
			if sum == search {
				return maths.Min(numbers[x:y+1]...) + maths.Max(numbers[x:y+1]...)
			}
		}
	}
	panic("boink")
}

func canSum(numbers ...int) bool {
	search := numbers[len(numbers)-1]
	for x := 0; x < len(numbers); x++ {
		for y := x + 1; y < len(numbers)-1; y++ {
			if numbers[x] == numbers[y] {
				continue
			}
			if numbers[x]+numbers[y] == search {
				return true
			}
		}
	}
	return false
}
