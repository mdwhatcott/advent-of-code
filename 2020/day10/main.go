package advent

import (
	"sort"

	"advent/lib/util"
)

func Part1() interface{} {
	adapters := util.InputInts("\n")
	sort.Ints(adapters)

	current := 0
	diffs := map[int]int{}

	for _, adapter := range adapters {
		diffs[adapter-current]++
		current = adapter
	}
	diffs[3]++

	return diffs[1] * diffs[3]
}

func Part2() interface{} {
	adapters := util.InputInts("\n")
	return calculateTotalAdapterArrangements(adapters)
}

func calculateTotalAdapterArrangements(adapters []int) interface{} {
	sort.Ints(adapters)

	sum := 1
	consecutive := 0
	diff := 0
	for x, a := range adapters {
		if x == 0 {
			diff = a - 0
		} else {
			diff = a - adapters[x-1]
		}

		if diff == 1 {
			consecutive++
		} else {
			if consecutive > 1 {
				sum *= consecutiveOnesMultipliers[consecutive]
			}
			consecutive = 0
		}
	}
	sum *= consecutiveOnesMultipliers[consecutive]
	return sum
}

var consecutiveOnesMultipliers = map[int]int{
	2: 2, // if we have 2 consecutive diffs of 1, we can toggle one of them, doubling the possibilities
	3: 4, // if we have 3 consecutive diffs of 1, we can toggle two of them, quadrupling the possibilities
	4: 7, // if we have 4 consecutive diffs of 1, we multiply the possibilities by 7
	// There's probably some cool pattern (fibonacci?) at play here...
	// ...oh, it's called tribonacci https://encyclopediaofmath.org/wiki/Tribonacci_sequence
}

/*

The Python REPL session wherein I realized an off-by-one error in my paper-pencil combinatorial experiments:

Python 2.7.16 (default, Jun  5 2020, 22:59:21)
[GCC 4.2.1 Compatible Apple LLVM 11.0.3 (clang-1103.0.29.20) (-macos10.15-objc- on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>> 6 * 2 * 4 * 6 * 6 * 6 * 6 * 6 * 6 * 4 * 4 * 6 * 4 * 6 * 6 * 6 * 4 * 6 * 4
17832200896512
>>> 6 * 6 * 4 * 2 * 6 * 6
10368
>>> 6 * 6 * 4 * 2 * 6 * 4
6912
>>> 6 * 6 * 4 * 2 * 6 * 6
10368
>>> 19208 / 6 / 6 / 4 / 2 / 6 / 6
1
>>> 19208.0 / 6 / 6 / 4 / 2 / 6 / 6
1.8526234567901234
>>> 19208.0 / 6 / 6 / 4 / 2 / 6 / 6
KeyboardInterrupt
>>> 6 * 6 * 4 * 2 * 6 * 7
12096
>>> 6 * 6 * 4 * 2 * 7 * 7
14112
>>> 6 * 7 * 4 * 2 * 7 * 7
16464
>>> 7 * 7 * 4 * 2 * 7 * 7
19208  # This is where I realize that I was off by one on my assumption that 4 diffs of 1 in a row would multiply the solution space by 6, but it was actually 7!

 */