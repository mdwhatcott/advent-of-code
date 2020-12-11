package advent

import (
	"fmt"
	"sort"
	"testing"

	"advent/lib/util"
)

var firstExample = []int{
	16,
	10,
	15,
	5,
	1,
	11,
	7,
	19,
	6,
	12,
	4,
}

func TestStuff1(t *testing.T) {
	sort.Ints(firstExample)
	for x, a := range firstExample {
		if x == 0 {
			fmt.Println(x, a)
			continue
		}
		fmt.Println(x, a, a-firstExample[x-1])
	}
}

var secondExample = []int{
	1,  // 1
	2,  // 1
	3,  // 1
	4,  // 1
	7,  // 3
	8,  // 1
	9,  // 1
	10, // 1
	11, // 1
	14, // 3
	17, // 3
	18, // 1
	19, // 1
	20, // 1
	23, // 3
	24, // 1
	25, // 1
	28, // 3
	31, // 3
	32, // 1
	33, // 1
	34, // 1
	35, // 1
	38, // 3
	39, // 1
	42, // 3
	45, // 3
	46, // 1
	47, // 1
	48, // 1
	49, // 1

	//52, // 3 // final device jolts
}

func TestStuff2(t *testing.T) {
	sort.Ints(secondExample)
	for x, a := range secondExample {
		if x == 0 {
			t.Logf("%2d, // %d\n", a, a-0)
			continue
		}
		t.Logf("%2d, // %d\n", a, a-secondExample[x-1])
	}

	actual := calculateTotalAdapterArrangements(secondExample)
	if actual != 19208 {
		t.Error("Expected 19208, got:", actual)
	}
}

func TestActualInput(t *testing.T) {
	actual := util.InputInts("\n")
	sort.Ints(actual)

	for x, a := range actual {
		if x == 0 {
			fmt.Printf("%2d, // %d\n", a, a-0)
			continue
		}
		fmt.Printf("%2d, // %d\n", a, a-actual[x-1])
	}
}
