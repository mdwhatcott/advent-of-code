package main

func ComputeUltimateCookie(totalTeaspoons int, ingredients ...Ingredient) (max, max500Calories int) {
	for _, mixture := range mixtures(len(ingredients), totalTeaspoons) {
		var mixingBowl Ingredients
		for a, amount := range mixture {
			for x := 0; x < amount; x++ {
				mixingBowl = append(mixingBowl, ingredients[a])
			}
		}

		if attributes := mixingBowl.Sum(); attributes.Valid() {
			product := attributes.Product()
			if product > max {
				max = product
			}
			if attributes.Calories == 500 && product > max500Calories {
				max500Calories = product
			}
		}
	}

	return max, max500Calories
}

func mixtures(ingredients, totalUnits int) (all [][]int) {
	var start int
	if ingredients == 1 {
		start = totalUnits
	}

	for i := start; i < totalUnits+1; i++ {
		left := totalUnits - i
		if ingredients-1 != 0 {
			for _, mixture := range mixtures(ingredients-1, left) {
				all = append(all, append([]int{i}, mixture...))
			}
		} else {
			all = append(all, []int{i})
		}
	}

	return all
}

type Ingredient struct {
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

func (this Ingredient) Valid() bool {
	if this.Capacity < 1 {
		return false
	}
	if this.Durability < 1 {
		return false
	}
	if this.Flavor < 1 {
		return false
	}
	if this.Texture < 1 {
		return false
	}
	return true
}

func (this Ingredient) Product() int {
	return this.Capacity * this.Durability * this.Flavor * this.Texture
}

type Ingredients []Ingredient

func (this Ingredients) Sum() (sum Ingredient) {
	for _, i := range this {
		sum.Capacity += i.Capacity
	}
	for _, i := range this {
		sum.Durability += i.Durability
	}
	for _, i := range this {
		sum.Flavor += i.Flavor
	}
	for _, i := range this {
		sum.Texture += i.Texture
	}
	for _, i := range this {
		sum.Calories += i.Calories
	}
	return sum
}
