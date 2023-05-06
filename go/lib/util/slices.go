package util

import (
	"advent/lib/maths"
)

type Slice[T maths.Ordered] []T

func (haystack Slice[T]) Index(needle T) int {
	for i, straw := range haystack {
		if straw == needle {
			return i
		}
	}
	return -1
}

func (haystack Slice[T]) Count(needle T) (result int) {
	for _, straw := range haystack {
		if straw == needle {
			result++
		}
	}
	return result
}

func (haystack Slice[T]) Contains(needle T) bool {
	return haystack.Index(needle) > -1
}

func (haystack Slice[T]) Min() (min T) {
	var zero T
	if len(haystack) == 0 {
		return zero
	}
	min = haystack[0]
	for _, value := range haystack {
		if value < min {
			min = value
		}
	}
	return min
}

func (this Slice[T]) Unpack1() T {
	return this.Get(0)
}
func (this Slice[T]) Unpack2() (T, T) {
	return this.Get(0), this.Get(1)
}
func (this Slice[T]) Unpack3() (T, T, T) {
	return this.Get(0), this.Get(1), this.Get(2)
}
func (this Slice[T]) Unpack4() (T, T, T, T) {
	return this.Get(0), this.Get(1), this.Get(2), this.Get(3)
}
func (this Slice[T]) Unpack5() (T, T, T, T, T) {
	return this.Get(0), this.Get(1), this.Get(2), this.Get(3), this.Get(4)
}
func (this Slice[T]) Get(index int) T {
	if len(this) > 0 && index < len(this) {
		return this[index]
	} else {
		var zero T
		return zero
	}
}
