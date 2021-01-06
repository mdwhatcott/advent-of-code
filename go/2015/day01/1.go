package advent

func basement(input string) int {
	counter := 0
	for i, c := range input {
		if c == '(' {
			counter++
		} else {
			counter--
		}
		if counter == -1 {
			return i + 1
		}
	}
	panic("BOINK!")
}

func endingFloor(input string) int {
	counter := 0
	for _, c := range input {
		if c == '(' {
			counter++
		} else {
			counter--
		}
	}
	return counter
}
