package advent

import "github.com/mdwhatcott/go-collections/stack"

// https://runestone.academy/runestone/books/published/pythonds/BasicDS/InfixPrefixandPostfixExpressions.html

func ParseShuntingYard(precedence map[rune]int, input string) []rune {
	ops := stack.New[rune](0)
	result := stack.New[rune](0)
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
	return result.Slice()
}

func EvalPostfix(postfix string) (result int) {
	operands := stack.New[int](0)
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
