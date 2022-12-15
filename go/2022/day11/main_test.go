package day10

import (
	"sort"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestDay11(t *testing.T) {
	should.So(t, Business(sampleMonkeys(), 20, 3), should.Equal, 10605)
	should.So(t, Business(inputMonkeys(), 20, 3), should.Equal, 100345)

	// Ok, part 2 gets all math-y and I had to check the subreddit for clues.
	// the concept of 'relief' becomes the product of the modulus operands of each monkey.
	// since all the math is based on these prime modulus operands, it is safe to just multiply them together for use as a general modulus.
	// The thing I don't really understand is why dividing by 3 was ok for part 1--why can't we just use the product-modulus value for both parts?
	// I guess there's still something about part 2, mathematically, that eludes me.
	const sampleMonkeyModulusProduct = 23 * 19 * 13 * 17
	const inputMonkeyModulusProduct = 13 * 5 * 19 * 11 * 17 * 3 * 7 * 2
	should.So(t, Business(sampleMonkeys(), 10000, sampleMonkeyModulusProduct), should.Equal, 2713310158)
	should.So(t, Business(inputMonkeys(), 10000, inputMonkeyModulusProduct), should.Equal, 28537348205)
}
func Business(monkeys []*Monkey, rounds int, relief int) int {
	Rounds(monkeys, rounds, relief)
	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].Inspected > monkeys[j].Inspected })
	return monkeys[0].Inspected * monkeys[1].Inspected
}
func Rounds(monkeys []*Monkey, rounds int, relief int) {
	for x := 0; x < rounds; x++ {
		Round(monkeys, relief)
	}
}
func Round(monkeys []*Monkey, relief int) {
	for _, monkey := range monkeys {
		for len(monkey.Items) > 0 {
			monkey.Inspected++
			item := monkey.Items[0]
			item = monkey.Operation(item)
			if relief == 3 {
				item /= relief
			} else {
				item %= relief
			}
			receiver := monkeys[monkey.Test(item)]
			receiver.Items = append(receiver.Items, item)
			monkey.Items = monkey.Items[1:]
		}
	}
}
func inputMonkeys() []*Monkey {
	return []*Monkey{
		{ID: 0, Items: []int{80}, Operation: func(old int) (new int) { return old * 5 },
			Test: func(n int) int { return map[bool]int{true: 4, false: 3}[n%2 == 0] }},
		{ID: 1, Items: []int{75, 83, 74}, Operation: func(old int) (new int) { return old + 7 },
			Test: func(n int) int { return map[bool]int{true: 5, false: 6}[n%7 == 0] }},
		{ID: 2, Items: []int{86, 67, 61, 96, 52, 63, 73}, Operation: func(old int) (new int) { return old + 5 },
			Test: func(n int) int { return map[bool]int{true: 7, false: 0}[n%3 == 0] }},
		{ID: 3, Items: []int{85, 83, 55, 85, 57, 70, 85, 52}, Operation: func(old int) (new int) { return old + 8 },
			Test: func(n int) int { return map[bool]int{true: 1, false: 5}[n%17 == 0] }},
		{ID: 4, Items: []int{67, 75, 91, 72, 89}, Operation: func(old int) (new int) { return old + 4 },
			Test: func(n int) int { return map[bool]int{true: 3, false: 1}[n%11 == 0] }},
		{ID: 5, Items: []int{66, 64, 68, 92, 68, 77}, Operation: func(old int) (new int) { return old * 2 },
			Test: func(n int) int { return map[bool]int{true: 6, false: 2}[n%19 == 0] }},
		{ID: 6, Items: []int{97, 94, 79, 88}, Operation: func(old int) (new int) { return old * old },
			Test: func(n int) int { return map[bool]int{true: 2, false: 7}[n%5 == 0] }},
		{ID: 7, Items: []int{77, 85}, Operation: func(old int) (new int) { return old + 6 },
			Test: func(n int) int { return map[bool]int{true: 4, false: 0}[n%13 == 0] }},
	}
}
func sampleMonkeys() []*Monkey {
	return []*Monkey{
		{ID: 0, Items: []int{79, 98}, Operation: func(old int) (new int) { return old * 19 },
			Test: func(n int) int { return map[bool]int{true: 2, false: 3}[n%23 == 0] }},
		{ID: 1, Items: []int{54, 65, 75, 74}, Operation: func(old int) (new int) { return old + 6 },
			Test: func(n int) int { return map[bool]int{true: 2, false: 0}[n%19 == 0] }},
		{ID: 2, Items: []int{79, 60, 97}, Operation: func(old int) (new int) { return old * old },
			Test: func(n int) int { return map[bool]int{true: 1, false: 3}[n%13 == 0] }},
		{ID: 3, Items: []int{74}, Operation: func(old int) (new int) { return old + 3 },
			Test: func(n int) int { return map[bool]int{true: 0, false: 1}[n%17 == 0] }},
	}
}

type Monkey struct {
	ID        int
	Items     []int
	Operation func(old int) (new int)
	Test      func(n int) (monkeyID int)
	Inspected int
}
