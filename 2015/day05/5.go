package advent

func nice(input string) int {
	vowels := 0
	twice := false
	var last rune
	for i, this := range input {
		switch this {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
		if i > 0 {
			if this == last {
				twice = true
			}
			if this == 'b' && last == 'a' {
				return 0
			}
			if this == 'd' && last == 'c' {
				return 0
			}
			if this == 'q' && last == 'p' {
				return 0
			}
			if this == 'y' && last == 'x' {
				return 0
			}
		}
		last = this
	}
	if vowels >= 3 && twice {
		return 1
	}
	return 0
}

func nice2(input string) bool {
	pairs := []string{}
	triplets := []string{}
	for i := range input {
		if i > 0 {
			pairs = append(pairs, input[i-1:i+1])
		}
		if i > 1 {
			triplets = append(triplets, input[i-2:i+1])
		}
	}
	pair := false
	triplet := false

	for i, a := range pairs[:len(pairs)-1] {
		for _, b := range pairs[i+2:] {
			if a == b {
				pair = true
			}
		}
	}

	for _, t := range triplets {
		if t[0] == t[2] && t[1] != t[0] {
			triplet = true
		}
	}

	return pair && triplet
}
