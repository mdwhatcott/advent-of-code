package part2

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/suite"
)

func TestCalculatorSuite(t *testing.T) {
	suite.Run(&CalculatorSuite{T: t})
}

type CalculatorSuite struct {
	*testing.T
}

func (this *CalculatorSuite) Test() {
	assert.With(this).That(calculate("1")).Equals(1)
	assert.With(this).That(calculate("1 + 2")).Equals(3)
	assert.With(this).That(calculate("1 * 2")).Equals(2)
	assert.With(this).That(calculate("1 + 2 * 3")).Equals(9)
	assert.With(this).That(simplify("10 * ( 2 + 3 )")).Equals("10 * 5")
	assert.With(this).That(Calculate("1 + 2 * 3")).Equals(9)
	assert.With(this).That(Calculate("10 * (2 + 3)")).Equals(50)
	assert.With(this).That(Calculate("1 + 2 * 3 + 4")).Equals(21)
	assert.With(this).That(furtherSimplify("1 + 2 * 3 + 4")).Equals("(1 + 2) * (3 + 4)")
	assert.With(this).That(Calculate("1 + (2 * 3) + (4 * (5 + 6))")).Equals(51)
	//assert.With(this).That(Calculate("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2")).Equals(23340)
}
