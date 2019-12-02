package advent

func RunProgram(program []int) {
	cursor := 0
	position, value := Compute(cursor, program)
	for position >= 0 {
		program[position] = value
		cursor += 4
		position, value = Compute(cursor, program)
	}
}

func Compute(cursor int, program []int) (position, value int) {
	if program[cursor] == 99 {
		position = -1
		return position, value
	}
	if program[cursor] == 1 {
		value = program[program[cursor+1]] + program[program[cursor+2]]
	} else if program[cursor] == 2 {
		value = program[program[cursor+1]] * program[program[cursor+2]]
	}
	position = program[cursor+3]
	return position, value
}
