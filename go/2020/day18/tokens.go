package advent

type TokenType string

const (
	TokenUnrecognized     TokenType = "?"
	TokenNumber           TokenType = "N"
	TokenAddition         TokenType = "+"
	TokenMultiplication   TokenType = "*"
	TokenLeftParenthesis  TokenType = "("
	TokenRightParenthesis TokenType = ")"
)

type Token struct {
	Type  TokenType
	Value int
}
