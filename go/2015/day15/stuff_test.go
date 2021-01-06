package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestStuff(t *testing.T) {
	gunit.Run(new(Stuff), t)
}

type Stuff struct {
	*gunit.Fixture

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
	this.So(all[0], should.Resemble, []int{0, 0, 6})
	this.So(all[1], should.Resemble, []int{0, 1, 5})
	this.So(all[2], should.Resemble, []int{0, 2, 4})
	this.So(all[3], should.Resemble, []int{0, 3, 3})
	this.So(all[4], should.Resemble, []int{0, 4, 2})
	this.So(all[5], should.Resemble, []int{0, 5, 1})
	this.So(all[6], should.Resemble, []int{0, 6, 0})
	this.So(all[7], should.Resemble, []int{1, 0, 5})
	this.So(all[8], should.Resemble, []int{1, 1, 4})
	this.So(all[9], should.Resemble, []int{1, 2, 3})
	this.So(all[10], should.Resemble, []int{1, 3, 2})
	this.So(all[11], should.Resemble, []int{1, 4, 1})
	this.So(all[12], should.Resemble, []int{1, 5, 0})
	this.So(all[13], should.Resemble, []int{2, 0, 4})
	this.So(all[14], should.Resemble, []int{2, 1, 3})
	this.So(all[15], should.Resemble, []int{2, 2, 2})
	this.So(all[16], should.Resemble, []int{2, 3, 1})
	this.So(all[17], should.Resemble, []int{2, 4, 0})
	this.So(all[18], should.Resemble, []int{3, 0, 3})
	this.So(all[19], should.Resemble, []int{3, 1, 2})
	this.So(all[20], should.Resemble, []int{3, 2, 1})
	this.So(all[21], should.Resemble, []int{3, 3, 0})
	this.So(all[22], should.Resemble, []int{4, 0, 2})
	this.So(all[23], should.Resemble, []int{4, 1, 1})
	this.So(all[24], should.Resemble, []int{4, 2, 0})
	this.So(all[25], should.Resemble, []int{5, 0, 1})
	this.So(all[26], should.Resemble, []int{5, 1, 0})
	this.So(all[27], should.Resemble, []int{6, 0, 0})
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
	this.So(sum, should.Resemble, Ingredient{
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
