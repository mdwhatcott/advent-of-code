package advent

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/suite"
)

func TestLexerFixture(t *testing.T) {
	suite.Run(&LexerFixture{T: t})
}

type LexerFixture struct{ *testing.T }

func (this *LexerFixture) assert(input string, expected ...Token) {
	actual := NewLexer(input).Lex()
	assert.With(this).That(actual).Equals(expected)
}
func (this *LexerFixture) TestLex() {
	this.assert("")
	this.assert("gibberish 1",
		Token{Type: TokenUnrecognized},
		Token{Type: TokenNumber, Value: 1},
	)
	this.assert("+",
		Token{Type: TokenAddition},
	)
	this.assert("1",
		Token{Type: TokenNumber, Value: 1},
	)
	this.assert("1234",
		Token{Type: TokenNumber, Value: 1234},
	)
	this.assert("12 + 34",
		Token{Type: TokenNumber, Value: 12},
		Token{Type: TokenAddition},
		Token{Type: TokenNumber, Value: 34},
	)
	this.assert("1 + (2 * 3)",
		Token{Type: TokenNumber, Value: 1},
		Token{Type: TokenAddition},
		Token{Type: TokenLeftParenthesis},
		Token{Type: TokenNumber, Value: 2},
		Token{Type: TokenMultiplication},
		Token{Type: TokenNumber, Value: 3},
		Token{Type: TokenRightParenthesis},
	)
}
