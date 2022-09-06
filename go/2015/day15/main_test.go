package main

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

func Test(t *testing.T) {
	max, max500Calories := ComputeUltimateCookie(100, parseIngredients()...)
	t.Log("Part 1 (the ultimate cookie)")
	should.So(t, max, should.Equal, 21367368)

	t.Log("Part 2 (the meal replacement)")
	should.So(t, max500Calories, should.Equal, 1766400)
}

func parseIngredients() (ingredients Ingredients) {
	for scanner := util.InputScanner(); scanner.Scan(); {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		line = strings.Replace(line, ",", "", -1)
		words := strings.Fields(line)
		ingredients = append(ingredients, Ingredient{
			Capacity:   util.ParseInt(words[2]),
			Durability: util.ParseInt(words[4]),
			Flavor:     util.ParseInt(words[6]),
			Texture:    util.ParseInt(words[8]),
			Calories:   util.ParseInt(words[10]),
		})
	}
	return ingredients
}
