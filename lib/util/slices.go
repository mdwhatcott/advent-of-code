package util

type Strings []string

func (haystack Strings) Index(needle string) int {
	for i, straw := range haystack {
		if straw == needle {
			return i
		}
	}
	return -1
}

func (haystack Strings) Contains(needle string) bool {
	return haystack.Index(needle) > -1
}

func (haystack Strings) Min() (min string) {
	if len(haystack) == 0 {
		return ""
	}
	min = haystack[0]
	for _, value := range haystack {
		if value < min {
			min = value
		}
	}
	return min
}

//////////////////////////////////////////////////////////

type Ints []int

func (haystack Ints) Index(needle int) int {
	for i, straw := range haystack {
		if straw == needle {
			return i
		}
	}
	return -1
}

func (haystack Ints) Contains(needle int) bool {
	return haystack.Index(needle) > -1
}


func (this Ints) Min() (min int) {
	return Min(this...)
}

func (haystack Ints) Sum() (sum int) {
	for _, value := range haystack {
		sum += value
	}
	return sum
}