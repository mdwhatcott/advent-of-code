package part2

import (
	"log"
	"strconv"
	"strings"
)

func Calculate(input string) (answer int) {
	for strings.Contains(input, "(") {
		input = simplify(input)
	}
	if strings.Contains(input, "*") && strings.Contains(input, "+") {
		input = furtherSimplify(input)
		return Calculate(input)
	}
	return calculate(input)
}

func simplify(expression string) string {
	var leftParen int
	for w, char := range expression {
		if char == '(' {
			leftParen = w
		} else if char == ')' {
			beforeParens := expression[:leftParen]
			betweenParens := expression[leftParen+1 : w]
			afterParens := expression[w+1:]
			if strings.Contains(betweenParens, "*") && strings.Contains(betweenParens, "+") {
				return beforeParens + furtherSimplify(betweenParens) + afterParens
			}
			answer := strconv.Itoa(calculate(betweenParens))
			return beforeParens + answer + afterParens
		}
	}
	panic("not possible")
}

func furtherSimplify(expression string) string {
	words := strings.Fields(expression)
	for w, word := range words {
		if word == "+" {
			words[w-1] = "(" + words[w-1]
			words[w+1] = words[w+1] + ")"
		}
	}
	simplified := strings.Join(words, " ")
	log.Println(expression, "was simplified to", simplified)
	return simplified
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
