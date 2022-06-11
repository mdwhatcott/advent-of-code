package util

func Min(all ...int) int {
	min := 0xffffffff
	for _, a := range all {
		if a < min {
			min = a
		}
	}
	return min
}

func Max(all ...int) (max int) {
	max = -0xffffffff
	for _, a := range all {
		if a > max {
			max = a
		}
	}
	return max
}

func MinMax(all ...int) (int, int) {
	return Min(all...), Max(all...)
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
