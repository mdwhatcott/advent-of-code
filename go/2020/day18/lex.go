package advent

import (
	"strconv"
	"strings"
)

type Lexer struct {
	words []string
}

func NewLexer(expression string) *Lexer {
	expression = strings.ReplaceAll(expression, "(", "( ")
	expression = strings.ReplaceAll(expression, ")", " )")
	return &Lexer{words: strings.Fields(expression)}
}

func (this Lexer) Lex() (tokens []Token) {
	for _, word := range this.words {
		tokens = append(tokens, makeToken(word))
	}
	return tokens
}
func makeToken(word string) Token {
	switch word {
	case "+":
		return Token{Type: TokenAddition}
	case "*":
		return Token{Type: TokenMultiplication}
	case "(":
		return Token{Type: TokenLeftParenthesis}
	case ")":
		return Token{Type: TokenRightParenthesis}
	}
	number, err := strconv.Atoi(word)
	if err != nil {
		return Token{Type: TokenUnrecognized}
	}
	return Token{Type: TokenNumber, Value: number}
}
