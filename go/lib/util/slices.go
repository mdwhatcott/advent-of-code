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

func (this Strings) Unpack1() string {
	return this.Get(0)
}
func (this Strings) Unpack2() (string, string) {
	return this.Get(0), this.Get(1)
}
func (this Strings) Unpack3() (string, string, string) {
	return this.Get(0), this.Get(1), this.Get(2)
}
func (this Strings) Unpack4() (string, string, string, string) {
	return this.Get(0), this.Get(1), this.Get(2), this.Get(3)
}
func (this Strings) Unpack5() (string, string, string, string, string) {
	return this.Get(0), this.Get(1), this.Get(2), this.Get(3), this.Get(4)
}
func (this Strings) Get(index int) string {
	if len(this) > 0 && index < len(this) {
		return this[index]
	} else {
		return ""
	}
}

//////////////////////////////////////////////////////////

type Ints []int

func (this Ints) Index(needle int) int {
	for i, straw := range this {
		if straw == needle {
			return i
		}
	}
	return -1
}

func (this Ints) Contains(needle int) bool {
	return this.Index(needle) > -1
}

func (this Ints) Min() (min int) {
	return Min(this...)
}

func (this Ints) Sum() (sum int) {
	for _, value := range this {
		sum += value
	}
	return sum
}

func (this Ints) Unpack1() int {
	return this.Get(0)
}
func (this Ints) Unpack2() (int, int) {
	return this.Get(0), this.Get(1)
}
func (this Ints) Unpack3() (int, int, int) {
	return this.Get(0), this.Get(1), this.Get(2)
}
func (this Ints) Unpack4() (int, int, int, int) {
	return this.Get(0), this.Get(1), this.Get(2), this.Get(3)
}
func (this Ints) Unpack5() (int, int, int, int, int) {
	return this.Get(0), this.Get(1), this.Get(2), this.Get(3), this.Get(4)
}
func (this Ints) Get(index int) int {
	if len(this) > 0 && index < len(this) {
		return this[index]
	} else {
		return 0
	}
}
