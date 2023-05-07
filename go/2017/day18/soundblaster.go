package day18

import (
	"strings"

	"github.com/mdwhatcott/advent-of-code/go/lib/parse"
)

type SoundBlaster struct {
	registers map[string]int
	program   []string
	cursor    int
	sound     int
	recovered int
}

func NewSoundBlaster(program []string) *SoundBlaster {
	return &SoundBlaster{
		program:   program,
		registers: make(map[string]int),
	}
}

func (this *SoundBlaster) Sound() int {
	return this.sound
}

func (this *SoundBlaster) Run() {
	for this.cursorInProgram() && this.recovered == 0 {
		this.cursor = this.execute(this.program[this.cursor])
	}
}

func (this *SoundBlaster) cursorInProgram() bool {
	return this.cursor >= 0 && this.cursor < len(this.program)
}

func (this *SoundBlaster) resolve(label string) int {
	value, found := this.registers[label]
	if found {
		return value
	}
	return parse.Int(label)
}
func (this *SoundBlaster) execute(instruction string) int {
	fields := strings.Fields(instruction)
	switch fields[0] {
	case "snd":
		this.sound = this.resolve(fields[1])
	case "set":
		this.registers[fields[1]] = this.resolve(fields[2])
	case "add":
		this.registers[fields[1]] += this.resolve(fields[2])
	case "mul":
		this.registers[fields[1]] *= this.resolve(fields[2])
	case "mod":
		this.registers[fields[1]] %= this.resolve(fields[2])
	case "rcv":
		if this.resolve(fields[1]) > 0 {
			this.recovered = this.sound
		}
	case "jgz":
		if this.resolve(fields[1]) > 0 {
			return this.cursor + this.resolve(fields[2])
		}
	}
	return this.cursor + 1
}
