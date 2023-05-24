package day15

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code-go-lib/parse"
	"github.com/mdwhatcott/advent-of-code-go-lib/util"
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
			Capacity:   parse.Int(words[2]),
			Durability: parse.Int(words[4]),
			Flavor:     parse.Int(words[6]),
			Texture:    parse.Int(words[8]),
			Calories:   parse.Int(words[10]),
		})
	}
	return ingredients
}
