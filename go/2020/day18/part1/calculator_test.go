package part1

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/suite"
)

func TestInterpreterSuite(t *testing.T) {
	suite.Run(&InterpreterSuite{T: t})
}

type InterpreterSuite struct {
	*testing.T
}

func (this *InterpreterSuite) Test() {
	assert.With(this).That(calculate("1")).Equals(1)
	assert.With(this).That(calculate("1 + 2")).Equals(3)
	assert.With(this).That(calculate("1 * 2")).Equals(2)
	assert.With(this).That(calculate("1 + 2 * 3")).Equals(9)
	assert.With(this).That(simplify("10 * ( 2 + 3 )")).Equals("10 * 5")
	assert.With(this).That(Calculate("1 + 2 * 3")).Equals(9)
	assert.With(this).That(Calculate("10 * (2 + 3)")).Equals(50)
}
