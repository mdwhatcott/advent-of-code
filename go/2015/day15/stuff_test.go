package day15

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestStuff(t *testing.T) {
	should.Run(&Stuff{T: should.New(t)}, should.Options.UnitTests())
}

type Stuff struct {
	*should.T

	butterscotch Ingredient
	cinnamon     Ingredient
}

func (this *Stuff) Setup() {
	this.butterscotch = Ingredient{
		Capacity:   -1,
		Durability: -2,
		Flavor:     6,
		Texture:    3,
		Calories:   8,
	}
	this.cinnamon = Ingredient{
		Capacity:   2,
		Durability: 3,
		Flavor:     -2,
		Texture:    -1,
		Calories:   3,
	}
}

func (this *Stuff) TestIngredientProportions() {
	max, max500Calories := ComputeUltimateCookie(100, this.butterscotch, this.cinnamon)
	this.So(max, should.Equal, 62842880)
	this.So(max500Calories, should.Equal, 57600000)
}

func (this *Stuff) TestMixtures() {
	all := mixtures(3, 6)
	if !this.So(len(all), should.Equal, 28) {
		return
	}
	this.So(all[0], should.Equal, []int{0, 0, 6})
	this.So(all[1], should.Equal, []int{0, 1, 5})
	this.So(all[2], should.Equal, []int{0, 2, 4})
	this.So(all[3], should.Equal, []int{0, 3, 3})
	this.So(all[4], should.Equal, []int{0, 4, 2})
	this.So(all[5], should.Equal, []int{0, 5, 1})
	this.So(all[6], should.Equal, []int{0, 6, 0})
	this.So(all[7], should.Equal, []int{1, 0, 5})
	this.So(all[8], should.Equal, []int{1, 1, 4})
	this.So(all[9], should.Equal, []int{1, 2, 3})
	this.So(all[10], should.Equal, []int{1, 3, 2})
	this.So(all[11], should.Equal, []int{1, 4, 1})
	this.So(all[12], should.Equal, []int{1, 5, 0})
	this.So(all[13], should.Equal, []int{2, 0, 4})
	this.So(all[14], should.Equal, []int{2, 1, 3})
	this.So(all[15], should.Equal, []int{2, 2, 2})
	this.So(all[16], should.Equal, []int{2, 3, 1})
	this.So(all[17], should.Equal, []int{2, 4, 0})
	this.So(all[18], should.Equal, []int{3, 0, 3})
	this.So(all[19], should.Equal, []int{3, 1, 2})
	this.So(all[20], should.Equal, []int{3, 2, 1})
	this.So(all[21], should.Equal, []int{3, 3, 0})
	this.So(all[22], should.Equal, []int{4, 0, 2})
	this.So(all[23], should.Equal, []int{4, 1, 1})
	this.So(all[24], should.Equal, []int{4, 2, 0})
	this.So(all[25], should.Equal, []int{5, 0, 1})
	this.So(all[26], should.Equal, []int{5, 1, 0})
	this.So(all[27], should.Equal, []int{6, 0, 0})
}

func (this *Stuff) TestIngredientMath() {
	var ingredients Ingredients
	for x := 0; x < 44; x++ {
		ingredients = append(ingredients, this.butterscotch)
	}
	for x := 0; x < 56; x++ {
		ingredients = append(ingredients, this.cinnamon)
	}

	sum := ingredients.Sum()
	this.So(sum, should.Equal, Ingredient{
		Capacity:   68,
		Durability: 80,
		Flavor:     152,
		Texture:    76,
		Calories:   520,
	})
	this.So(sum.Product(), should.Equal, 62842880)
}

func (this *Stuff) TestValidation() {
	this.So(Ingredient{1, 1, 1, 1, 1}.Valid(), should.BeTrue)
	this.So(Ingredient{0, 1, 1, 1, 1}.Valid(), should.BeFalse)
	this.So(Ingredient{1, 0, 1, 1, 1}.Valid(), should.BeFalse)
	this.So(Ingredient{1, 1, 0, 1, 1}.Valid(), should.BeFalse)
	this.So(Ingredient{1, 1, 1, 0, 1}.Valid(), should.BeFalse)
	//this.So(Ingredient{1, 1, 1, 1, 0}.Valid(), should.BeFalse)
}
