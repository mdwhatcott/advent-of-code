package advent

import (
	"sort"
	"strings"

	"github.com/smartystreets/logging"
)

// TODO: Really nice solution

type ChemicalConversion struct {
	RequiredInputQuantities map[string]int
	ProducedOutputQuantity  int
}

type Calculator struct {
	log *logging.Logger

	outputConversions map[string]ChemicalConversion
	reservoir         map[string]int
	consumed          int
}

func NewCalculator() *Calculator {
	return &Calculator{
		outputConversions: make(map[string]ChemicalConversion),
		reservoir:         make(map[string]int),
	}
}

func (this *Calculator) TotalOREConsumed() int {
	return this.consumed
}

func sortKeys(m map[string]int) (keys []string) {
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func (this *Calculator) Resolve(chemicals map[string]int, depth int) {
	for _, chemical := range sortKeys(chemicals) {
		quantity := chemicals[chemical]
		this.resolve(chemical, quantity, depth)
	}
}

func (this *Calculator) resolve(chemical string, quantity int, depth int) {
	conversion := this.outputConversions[chemical]
	inputs := copyMap(conversion.RequiredInputQuantities)
	if inputs["ORE"] > 0 {
		for this.reservoir[chemical] > 0 && quantity > 0 {
			quantity--
			this.reservoir[chemical]--

			this.logReservoirDrain(depth, chemical)
		}

		for quantity > 0 {
			oreToConsume := conversion.RequiredInputQuantities["ORE"]
			this.consumed += oreToConsume
			quantity -= conversion.ProducedOutputQuantity

			this.logOreConsumed(depth, chemical, oreToConsume, conversion)
		}

		this.reservoir[chemical] += -quantity

		delete(inputs, chemical)
	} else {
		sets := quantity / conversion.ProducedOutputQuantity
		if quantity%conversion.ProducedOutputQuantity != 0 {
			sets++
		}

		this.logUpcomingCost(depth, chemical, quantity, sets, inputs)
		this.resolveMultiple(sets, inputs, depth)
	}
}

func (this *Calculator) resolveMultiple(quantity int, inputs map[string]int, depth int) {
	for x := 0; x < quantity; x++ {
		this.logResolutionProgress(depth, quantity, inputs, x)
		this.Resolve(inputs, depth+1)
	}
}

func (this *Calculator) logUpcomingCost(
	depth int,
	outputChemical string,
	outputQuantity int,
	inputQuantity int,
	inputs map[string]int,
) {
	this.log.Printf(
		"%s %d %s will cost %d %v",
		indent(depth),
		outputQuantity,
		outputChemical,
		inputQuantity,
		inputs,
	)
}

func (this *Calculator) logReservoirDrain(depth int, chemical string) {
	this.log.Printf(
		"%s Reservoir drained of 1 %s\n",
		indent(depth),
		chemical,
	)
}

func (this *Calculator) logOreConsumed(
	depth int,
	chemical string,
	oreToConsume int,
	conversion ChemicalConversion,
) {
	this.log.Printf(
		"%s Consuming %d ORE to produce %d %s (Total ORE consumed: %d)\n",
		indent(depth),
		oreToConsume,
		conversion.ProducedOutputQuantity,
		chemical,
		this.consumed,
	)
}
func (this *Calculator) logResolutionProgress(
	depth int,
	quantity int,
	inputs map[string]int,
	x int,
) {
	this.log.Printf(
		"%s Resolving %d/%d instances of %v",
		indent(depth),
		x+1,
		quantity,
		inputs,
	)
}
func indent(depth int) string {
	return strings.Repeat(".", depth+1)
}

func copyMap(a map[string]int) map[string]int {
	b := make(map[string]int, len(a))
	for key, value := range a {
		b[key] = value
	}
	return b
}
