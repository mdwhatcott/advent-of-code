package main

import (
	"fmt"

	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

func main() {
	input := util.InputLines()
	var containers []int
	for _, line := range input {
		containers = append(containers, util.ParseInt(line))
	}

	containersUsed := make(map[int]int)
	for mask := 1; mask < 1<<uint(len(containers)); mask++ {
		var combination []int
		for i, quantity := range containers {
			if (mask & (1 << uint(i))) > 0 {
				combination = append(combination, quantity)
			}
		}
		if sum(combination) == 150 {
			containersUsed[len(combination)]++
		}
	}

	fmt.Println("Part 1 - number of combinations adding up to 150:")
	should.So(nil, sumValues(containersUsed), should.Equal, 4372)

	fmt.Println("Part 2 - number of minimum container combinations:")
	should.So(nil, minKey(containersUsed), should.Equal, 4)
}

func sum(values []int) (total int) {
	for _, v := range values {
		total += v
	}
	return total
}

func sumValues(values map[int]int) (sum int) {
	for _, value := range values {
		sum += value
	}
	return sum
}
func minKey(values map[int]int) (min int) {
	min = 0xffffffff
	for key := range values {
		if key < min {
			min = key
		}
	}
	return min
}

/*

I achieved my answer by translating this incredible solution (which, from a mathematical standpoint, eludes me):
https://www.reddit.com/r/adventofcode/comments/3x6cyr/day_17_solutions/cy211o2/

from collections import defaultdict
dimensions = [50, 44, 11, 49, 42, 46, 18, 32, 26, 40, 21, 7, 18, 43, 10, 47, 36, 24, 22, 40]

dist = defaultdict(int)
for mask in xrange(1, 1<<len(dimensions)):
    p = [d for i,d in enumerate(dimensions) if (mask & (1 << i)) > 0]
    if sum(p) == 150: dist[len(p)] += 1

print "total:", sum(dist.values())
print "min:", dist[min(dist.keys())]

*/
