package part1

import (
	"strconv"
	"strings"
)

func Calculate(input string) (answer int) {
	for strings.Contains(input, "(") {
		input = simplify(input)
	}
	return calculate(input)
}

func simplify(original string) string {
	var leftParen int
	for w, char := range original {
		if char == '(' {
			leftParen = w
		} else if char == ')' {
			answer := strconv.Itoa(calculate(original[leftParen+1 : w]))
			beforeParens := original[:leftParen]
			afterParens := original[w+1:]
			return beforeParens + answer + afterParens
		}
	}
	panic("not possible")
}

func calculate(input string) (answer int) {
	operator := ops["+"]
	words := strings.Fields(input)
	for _, word := range words {
		number, err := strconv.Atoi(word)
		if err == nil {
			answer = operator(answer, number)
		} else {
			operator = ops[word]
		}
	}
	return answer
}

var ops = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"*": func(a, b int) int { return a * b },
}
