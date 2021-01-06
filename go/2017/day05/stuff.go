package day05

type Program struct {
	part2        bool
	cursor       int
	steps        int
	instructions []int
}

func NewProgram(instructions []int) *Program {
	return &Program{instructions: instructions}
}

func (this *Program) Part2() *Program {
	this.part2 = true
	return this
}

func (this *Program) Execute() int {
	for !this.finished() {
		this.Jump()
	}
	return this.steps
}

func (this *Program) Jump() {
	defer this.increment(this.cursor)
	this.steps++
	current := this.instructions[this.cursor]
	this.cursor += current
}

func (this *Program) finished() bool {
	return this.cursor < 0 || this.cursor >= len(this.instructions)
}

// NOTE: deferring the call to this method really slows the program down with more than 24 million iterations.
func (this *Program) increment(c int) {
	if this.part2 && this.instructions[c] >= 3 {
		this.instructions[c]--
	} else {
		this.instructions[c]++
	}
}
