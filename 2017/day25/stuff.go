package day25

import "log"

func Process() (checksum int) {
	machine := NewTuringMachine()

	for x := 0; x < 12302209; x++ {
		if x%1_000_000 == 0 {
			log.Println("Progress:", x)
		}
		machine.state = machine.state(machine)
	}

	for _, slot := range machine.tape {
		if slot == One {
			checksum++
		}
	}

	return checksum
}

const (
	Zero = false
	One  = true

	Right = 1
	Left  = -1
)

type StateFunc func(*TuringMachine) StateFunc

type TuringMachine struct {
	cursor int
	state  StateFunc
	tape   map[int]bool
}

func NewTuringMachine() *TuringMachine {
	return &TuringMachine{
		state: A,
		tape:  make(map[int]bool, 15_000_000),
	}
}

func A(this *TuringMachine) StateFunc {
	if this.tape[this.cursor] == Zero {
		this.tape[this.cursor] = One
		this.cursor += Right
		return B
	} else {
		this.tape[this.cursor] = Zero
		this.cursor += Left
		return D
	}
}
func B(this *TuringMachine) StateFunc {
	if this.tape[this.cursor] == Zero {
		this.tape[this.cursor] = One
		this.cursor += Right
		return C
	} else {
		this.tape[this.cursor] = Zero
		this.cursor += Right
		return F
	}
}
func C(this *TuringMachine) StateFunc {
	if this.tape[this.cursor] == Zero {
		this.tape[this.cursor] = One
		this.cursor += Left
		return C
	} else {
		this.tape[this.cursor] = One
		this.cursor += Left
		return A
	}
}
func D(this *TuringMachine) StateFunc {
	if this.tape[this.cursor] == Zero {
		this.tape[this.cursor] = Zero
		this.cursor += Left
		return E
	} else {
		this.tape[this.cursor] = One
		this.cursor += Right
		return A
	}
}
func E(this *TuringMachine) StateFunc {
	if this.tape[this.cursor] == Zero {
		this.tape[this.cursor] = One
		this.cursor += Left
		return A
	} else {
		this.tape[this.cursor] = Zero
		this.cursor += Right
		return B
	}
}
func F(this *TuringMachine) StateFunc {
	if this.tape[this.cursor] == Zero {
		this.tape[this.cursor] = Zero
		this.cursor += Right
		return C
	} else {
		this.tape[this.cursor] = Zero
		this.cursor += Right
		return E
	}
}
