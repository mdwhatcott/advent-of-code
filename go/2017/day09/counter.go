package day09

func groupScore(input string) (total int) {
	depth := 0
	for item := range LexItems(input) {
		switch item {
		case OpenGroup:
			depth++
			total += depth
		case CloseGroup:
			depth--
		}
	}
	return total
}

func garbageScore(input string) (total int) {
	for item := range LexItems(input) {
		if item == Garbage {
			total++
		}
	}
	return total
}
