package advent

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestRecipeFixture(t *testing.T) {
	gunit.Run(new(RecipeFixture), t)
}

type RecipeFixture struct {
	*gunit.Fixture
}

func (this *RecipeFixture) Setup() {
}

func (this *RecipeFixture) Test() {
}
