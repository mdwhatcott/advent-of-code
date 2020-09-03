package advent

import (
	"strings"

	"advent/lib/util"
)

func (this *Calculator) ParseAllConversions(lines []string) {
	for _, line := range lines {
		this.parseConversion(line)
	}
}
func (this *Calculator) parseConversion(line string) {
	var (
		inputOutput                    = strings.Split(line, " => ")
		input                          = inputOutput[0]
		output                         = inputOutput[1]
		inputs                         = parseRequiredInputQuantities(input)
		outputChemical, outputQuantity = parseChemicalAndQuantity(output)
	)
	this.registerConversion(outputChemical, outputQuantity, inputs)
}
func parseRequiredInputQuantities(rawInputs string) map[string]int {
	requiredInputQuantities := make(map[string]int)
	for _, rawInput := range strings.Split(rawInputs, ", ") {
		chemical, quantity := parseChemicalAndQuantity(rawInput)
		requiredInputQuantities[chemical] = quantity
	}
	return requiredInputQuantities
}
func parseChemicalAndQuantity(rawInput string) (name string, quantity int) {
	outputParts := strings.Fields(rawInput)
	return outputParts[1], util.ParseInt(outputParts[0])
}
func (this *Calculator) registerConversion(chemical string, quantity int, inputs map[string]int) {
	this.outputConversions[chemical] = ChemicalConversion{
		RequiredInputQuantities: inputs,
		ProducedOutputQuantity:  quantity,
	}
}
