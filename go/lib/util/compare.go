package util

import "advent/lib/constraints"

func Sum[T constraints.Number](all ...T) (result T) {
	for _, a := range all {
		result += a
	}
	return result
}

func Min[T constraints.Number](all ...T) T {
	min := all[0]
	for _, a := range all[1:] {
		if a < min {
			min = a
		}
	}
	return min
}

func Max[T constraints.Number](all ...T) (max T) {
	max = all[0]
	for _, a := range all[1:] {
		if a > max {
			max = a
		}
	}
	return max
}

func MinMax[T constraints.Number](all ...T) (min T, max T) {
	return Min(all...), Max(all...)
}

func Abs[T constraints.Number](a T) T {
	if a < 0 {
		return -a
	}
	return a
}
