package advent

// https://runestone.academy/runestone/books/published/pythonds/BasicDS/InfixPrefixandPostfixExpressions.html

func ParseShuntingYard(precedence map[rune]int, input string) []rune {
	ops := new(RuneStack)
	result := new(RuneStack)
	for _, c := range input {
		switch c {
		case ' ':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			result.Push(c)
		case '(':
			ops.Push(c)
		case ')':
			for ops.Peek() != '(' {
				result.Push(ops.Pop())
			}
			ops.Pop() // left paren
		default:
			for ops.Len() > 0 && ops.Peek() != '(' && precedence[ops.Peek()] >= precedence[c] {
				result.Push(ops.Pop())
			}
			ops.Push(c)
		}
	}
	for ops.Len() > 0 {
		result.Push(ops.Pop())
	}
	return result.items
}

func EvalPostfix(postfix string) (result int) {
	operands := new(IntStack)
	for _, c := range postfix {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			operands.Push(number(c))
		case '+':
			operands.Push(operands.Pop() + operands.Pop())
		case '*':
			operands.Push(operands.Pop() * operands.Pop())
		}
	}
	return operands.Pop()
}

func number(r rune) int {
	return int(r - '0')
}

type RuneStack struct {
	items []rune
}

func (this *RuneStack) Len() int {
	return len(this.items)
}
func (this *RuneStack) Push(s rune) {
	this.items = append(this.items, s)
}
func (this *RuneStack) Pop() rune {
	item := this.items[len(this.items)-1]
	this.items = this.items[:len(this.items)-1]
	return item
}
func (this *RuneStack) Peek() rune {
	return this.items[len(this.items)-1]
}

type IntStack struct {
	items []int
}

func (this *IntStack) Len() int {
	return len(this.items)
}
func (this *IntStack) Push(s int) {
	this.items = append(this.items, s)
}
func (this *IntStack) Pop() int {
	item := this.items[len(this.items)-1]
	this.items = this.items[:len(this.items)-1]
	return item
}
func (this *IntStack) Peek() int {
	return this.items[len(this.items)-1]
}
