package advent

// https://runestone.academy/runestone/books/published/pythonds/BasicDS/InfixPrefixandPostfixExpressions.html

func ParseShuntingYard(precedence map[rune]int, input string) []rune {
	operators := new(RuneStack)
	output := new(RuneStack)
	for _, c := range input {
		switch c {
		case ' ':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			output.Push(c)
		case '(':
			operators.Push(c)
		case ')':
			for operators.Peek() != '(' {
				output.Push(operators.Pop())
			}
			operators.Pop() // left paren
		default:
			for operators.Len() > 0 && operators.Peek() != '(' && precedence[operators.Peek()] >= precedence[c] {
				output.Push(operators.Pop())
			}
			operators.Push(c)
		}
	}
	for operators.Len() > 0 {
		output.Push(operators.Pop())
	}
	return output.items
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
