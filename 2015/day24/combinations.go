package main

func product(items ...int) (result int) {
	result = items[0]
	for _, item := range items[1:] {
		result *= item
	}
	return result
}

func combinations(iterable []int, count int) chan []int {
	stream := make(chan []int)
	go combinationGenerator(iterable, count, stream)
	return stream
}

// Credit: http://rosettacode.org/wiki/Combinations#Go
func combinationGenerator(source []int, count int, out chan []int) {
	working := make([]int, count)
	last := count - 1
	var generate func(int, int)
	generate = func(i, next int) {
		for j := next; j < len(source); j++ {
			working[i] = source[j]
			if i == last {
				combination := make([]int, len(working))
				copy(combination, working)
				out <- combination
			} else {
				generate(i+1, j+1)
			}
		}
	}
	generate(0, 0)
	close(out)
}
