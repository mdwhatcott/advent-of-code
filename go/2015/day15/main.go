package main

import (
	"fmt"
	"strings"

	"advent/lib/util"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func main() {
	max, max500Calories := ComputeUltimateCookie(100, parseIngredients()...)
	fmt.Println("Part 1 (the ultimate cookie):", assert.So(max, should.Equal, 21367368))
	fmt.Println("Part 2 (the meal replacement):", assert.So(max500Calories, should.Equal, 1766400))
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
