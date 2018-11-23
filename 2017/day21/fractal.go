package day21

import "strings"

type FractalArt struct {
	rules EnhancementRules
	grid  string
}

func CountFractalPixels(rules EnhancementRules, rounds int) int {
	art := &FractalArt{rules: rules, grid: start}
	for ; rounds > 0; rounds-- {
		art.Enhance()
	}
	return art.LightedPixels()
}

func (this *FractalArt) Enhance() {
	var updated []string
	for _, pattern := range SplitPatterns(this.grid) {
		updated = append(updated, this.rules.Enhance(pattern))
	}
	this.grid = ReassembleGrid(updated...)
}

func (this *FractalArt) LightedPixels() (count int) {
	return strings.Count(this.grid, "#")
}

const start = ".#./..#/###"
