package advent

func Part1() interface{} {
	return CalculateCombinedEnergy(1000, inputA, inputB, inputC, inputD)
}

func Part2() interface{} {
	return CalculatePeriodIntersection(CalculatePeriods(inputA, inputB, inputC, inputD))
}
