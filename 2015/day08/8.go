package advent

func CodeLength(value string) (length int) {
	return len(value)
}

func MemoryLength(value string) (length int) {
	runes := []rune(value)

	for x := 1; x < len(runes)-1; x++ {
		current := runes[x]
		if current == '\\' {
			if runes[x+1] == 'x' {
				x += 2
			}
			x++
			length++

			continue
		}
		length++
	}
	return
}

func EscapedLength(value string) (length int) {
	length = 2 // opening and closing quotes
	runes := []rune(value)
	for x := 0; x < len(runes); x++ {
		if runes[x] == '"' {
			length++
		} else if runes[x] == '\\' {
			length++
		}
		length++
	}
	return
}
